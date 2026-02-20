// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {AccessControl} from "openzeppelin-contracts/contracts/access/AccessControl.sol";
import {ReentrancyGuard} from "openzeppelin-contracts/contracts/utils/ReentrancyGuard.sol";
import {IBNPLManager} from "./interfaces/IBNPLManager.sol";
import {IDAOManager} from "./interfaces/IDAOManager.sol";

/// @title BNPLManager
/// @notice Full BNPL manager: validates DAO policy, computes installments, accepts payments, auto‑activates on first payment and applies late fees.
contract BNPLManager is AccessControl, ReentrancyGuard, IBNPLManager {
    bytes32 public constant BNPL_ADMIN_ROLE = keccak256("BNPL_ADMIN_ROLE");

    struct Arrangement {
        uint256 id;
        uint256 daoId;
        address payer;
        address recipient;
        uint256 totalAmount;
        uint256 numInstallments;
        uint256[] installmentAmounts;
        uint256 startTimestamp;
        uint256 intervalSeconds;
        uint256 lateFeeBps;
        uint8 status;
        mapping(uint256 => bool) installmentPaid;
    }

    uint256 private _nextId = 1;
    mapping(uint256 => Arrangement) private _arrangements;
    mapping(uint256 => mapping(uint8 => uint256)) private _paidAmounts;

    address public daoManager;

    /// @dev Grant DEFAULT_ADMIN_ROLE to deployer.
    constructor() {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
    }

    /// @notice Set the DAOManager address used for BNPL policy lookups.
    function setDaoManager(
        address _daoManager
    ) external onlyRole(DEFAULT_ADMIN_ROLE) {
        daoManager = _daoManager;
    }

    /// @inheritdoc IBNPLManager
    function createBNPL(
        uint256 daoId,
        address recipient,
        uint256 totalAmount,
        uint256 startTimestamp,
        uint256 intervalSeconds,
        bytes calldata
    ) external returns (uint256 arrangementId) {
        require(daoManager != address(0), "DAO_MANAGER_NOT_SET");

        uint256 numInstallments;
        uint256 allowedIntervalMinDays;
        uint256 allowedIntervalMaxDays;
        uint256 lateFeeBps;
        (
            numInstallments,
            allowedIntervalMinDays,
            allowedIntervalMaxDays,
            lateFeeBps,
            ,
            ,

        ) = IDAOManager(daoManager).getBnplTerms(daoId);
        require(numInstallments > 0, "DAO_BNPL_NOT_CONFIGURED");

        require(
            intervalSeconds >= allowedIntervalMinDays * 1 days,
            "INTERVAL_TOO_SHORT"
        );
        require(
            intervalSeconds <= allowedIntervalMaxDays * 1 days,
            "INTERVAL_TOO_LONG"
        );

        arrangementId = _nextId++;
        Arrangement storage a = _arrangements[arrangementId];
        a.id = arrangementId;
        a.daoId = daoId;
        a.payer = msg.sender;
        a.recipient = recipient;
        a.totalAmount = totalAmount;
        a.numInstallments = numInstallments;
        a.startTimestamp = startTimestamp;
        a.intervalSeconds = intervalSeconds;
        a.lateFeeBps = lateFeeBps;
        a.status = 0;

        uint256 base = totalAmount / numInstallments;
        uint256 remainder = totalAmount - (base * numInstallments);
        delete a.installmentAmounts;
        for (uint256 i = 0; i < numInstallments; i++) {
            uint256 amt = base + (i == numInstallments - 1 ? remainder : 0);
            a.installmentAmounts.push(amt);
        }

        emit BNPLCreated(
            arrangementId,
            daoId,
            msg.sender,
            recipient,
            totalAmount,
            a.numInstallments,
            block.timestamp
        );
        emit BNPLScheduleChosen(arrangementId, startTimestamp, intervalSeconds);
        return arrangementId;
    }

    /// @inheritdoc IBNPLManager
    function makePayment(
        uint256 arrangementId,
        uint8 installmentNumber
    ) external payable nonReentrant {
        Arrangement storage a = _arrangements[arrangementId];
        require(a.id != 0, "ARR_NOT_FOUND");
        require(installmentNumber < a.numInstallments, "INVALID_INSTALLMENT");
        uint256 expected = a.installmentAmounts[installmentNumber];
        require(msg.value >= expected, "INSUFFICIENT_PAYMENT");

        _paidAmounts[arrangementId][installmentNumber] += expected;
        a.installmentPaid[installmentNumber] = true;
        uint256 refund = msg.value - expected;

        emit BNPLPaymentMade(
            arrangementId,
            installmentNumber,
            msg.sender,
            expected,
            block.timestamp
        );

        if (refund > 0) {
            (bool sent, ) = msg.sender.call{value: refund}("");
            require(sent, "REFUND_FAILED");
        }

        if (installmentNumber == 0 && a.status == 0) {
            a.status = 1;
            emit BNPLActivated(arrangementId, block.timestamp);
        }

        bool allPaid = true;
        for (uint256 i = 0; i < a.numInstallments; i++) {
            if (!a.installmentPaid[i]) {
                allPaid = false;
                break;
            }
        }
        if (allPaid) {
            a.status = 2;
            emit BNPLCompleted(arrangementId, block.timestamp);
        }
    }

    /// @inheritdoc IBNPLManager
    function activateBNPL(uint256 arrangementId) external {
        Arrangement storage a = _arrangements[arrangementId];
        require(a.id != 0, "ARR_NOT_FOUND");
        require(a.status == 0, "NOT_PENDING");
        require(a.installmentPaid[0], "FIRST_PAYMENT_REQUIRED");
        a.status = 1;
        emit BNPLActivated(arrangementId, block.timestamp);
    }

    /// @inheritdoc IBNPLManager
    function applyLateFee(
        uint256 arrangementId,
        uint8 installmentNumber
    ) external {
        Arrangement storage a = _arrangements[arrangementId];
        require(a.id != 0, "ARR_NOT_FOUND");
        uint256 installment = a.installmentAmounts[installmentNumber];
        uint256 fee = (installment * a.lateFeeBps) / 10000;
        a.totalAmount += fee;

        if (daoManager != address(0) && fee > 0) {
            IDAOManager(daoManager).creditTreasury(a.daoId, fee);
        }

        emit BNPLLateFeeApplied(
            arrangementId,
            installmentNumber,
            fee,
            block.timestamp
        );
    }

    /// @inheritdoc IBNPLManager
    function reschedule(
        uint256 arrangementId,
        uint256 newStartTimestamp,
        uint256 newIntervalSeconds
    ) external {
        Arrangement storage a = _arrangements[arrangementId];
        require(a.id != 0, "ARR_NOT_FOUND");

        require(daoManager != address(0), "DAO_MANAGER_NOT_SET");
        (
            ,
            uint256 allowedIntervalMinDays,
            uint256 allowedIntervalMaxDays,
            ,
            ,
            bool rescheduleAllowed,

        ) = IDAOManager(daoManager).getBnplTerms(a.daoId);
        require(rescheduleAllowed, "RESCHEDULE_NOT_ALLOWED");

        require(
            msg.sender == a.payer || hasRole(DEFAULT_ADMIN_ROLE, msg.sender),
            "UNAUTHORIZED"
        );

        require(
            newIntervalSeconds >= allowedIntervalMinDays * 1 days &&
                newIntervalSeconds <= allowedIntervalMaxDays * 1 days,
            "INTERVAL_OUT_OF_BOUNDS"
        );

        bytes32 oldHash = keccak256(
            abi.encode(a.startTimestamp, a.intervalSeconds)
        );
        a.startTimestamp = newStartTimestamp;
        a.intervalSeconds = newIntervalSeconds;
        bytes32 newHash = keccak256(
            abi.encode(newStartTimestamp, newIntervalSeconds)
        );
        emit BNPLRescheduled(arrangementId, oldHash, newHash, block.timestamp);
    }

    /// @inheritdoc IBNPLManager
    function getArrangement(
        uint256 arrangementId
    )
        external
        view
        returns (
            uint256 id,
            uint256 daoId,
            address payer,
            address recipient,
            uint256 totalAmount,
            uint256 numInstallments,
            uint256[] memory installmentAmounts,
            uint256 startTimestamp,
            uint256 intervalSeconds,
            uint256 lateFeeBps,
            uint8 status
        )
    {
        Arrangement storage a = _arrangements[arrangementId];
        return (
            a.id,
            a.daoId,
            a.payer,
            a.recipient,
            a.totalAmount,
            a.numInstallments,
            a.installmentAmounts,
            a.startTimestamp,
            a.intervalSeconds,
            a.lateFeeBps,
            a.status
        );
    }
}
