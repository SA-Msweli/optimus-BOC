// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {AccessControl} from "openzeppelin-contracts/contracts/access/AccessControl.sol";
import {ReentrancyGuard} from "openzeppelin-contracts/contracts/utils/ReentrancyGuard.sol";
import {IERC20} from "openzeppelin-contracts/contracts/token/ERC20/IERC20.sol";
import {IDAOManager} from "./interfaces/IDAOManager.sol";
import {ITokenVault} from "./interfaces/ITokenVault.sol";

contract DAOManager is AccessControl, ReentrancyGuard, IDAOManager {
    bytes32 public constant DAO_ADMIN_ROLE = keccak256("DAO_ADMIN_ROLE");
    bytes32 public constant TREASURY_FUNDER_ROLE =
        keccak256("TREASURY_FUNDER_ROLE");

    struct BnplTerms {
        uint256 numInstallments;
        uint256 allowedIntervalMinDays;
        uint256 allowedIntervalMaxDays;
        uint256 lateFeeBps;
        uint256 gracePeriodDays;
        bool rescheduleAllowed;
        uint256 minDownPaymentBps;
    }

    struct Dao {
        address creator;
        uint8 goal;
        uint64 votingPeriodDays;
        uint256 treasuryBalance;
        uint256 memberCount;
        uint256 createdAt;
        bool isDissolved;
        uint256 quorumBps;
    }

    struct Proposal {
        uint256 daoId;
        bytes data;
        uint256 expiry;
        uint256 yesWeight;
        uint256 noWeight;
        bool finalized;
        bool executed;
    }

    uint256 private _nextDaoId = 1;
    uint256 private _nextProposalId = 1;

    mapping(uint256 => Dao) private _daos;
    mapping(uint256 => BnplTerms) private _bnplTerms;
    mapping(uint256 => Proposal) private _proposals;
    mapping(uint256 => mapping(address => uint256)) private _memberInvestments;
    mapping(uint256 => mapping(address => bool)) private _hasVoted;

    address public tokenVault;

    /// @dev Grant DEFAULT_ADMIN_ROLE to deployer.
    constructor() {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
    }

    function setTokenVault(
        address _tokenVault
    ) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(_tokenVault != address(0), "INVALID_VAULT");
        tokenVault = _tokenVault;
    }

    /// @inheritdoc IDAOManager
    function createDAO(
        address creator,
        uint8 goal,
        uint64 votingPeriodDays
    ) external returns (uint256 daoId) {
        daoId = _nextDaoId++;
        _daos[daoId] = Dao({
            creator: creator,
            goal: goal,
            votingPeriodDays: votingPeriodDays,
            treasuryBalance: 0,
            memberCount: 0,
            createdAt: block.timestamp,
            isDissolved: false,
            quorumBps: 0
        });
        emit DaoCreated(daoId, creator, goal, block.timestamp);
        return daoId;
    }

    /// @inheritdoc IDAOManager
    function joinDAO(
        uint256 daoId,
        address member,
        uint256 investmentAmount
    ) external {
        Dao storage d = _daos[daoId];
        require(
            d.createdAt != 0 && !d.isDissolved,
            "DAO_NOT_FOUND_OR_DISSOLVED"
        );
        d.memberCount += 1;
        d.treasuryBalance += investmentAmount;
        _memberInvestments[daoId][member] += investmentAmount;
        emit MemberJoined(daoId, member, investmentAmount, block.timestamp);
    }

    /// @inheritdoc IDAOManager
    function propose(
        uint256 daoId,
        bytes calldata proposalData
    ) external returns (uint256 proposalId) {
        Proposal storage p = _proposals[_nextProposalId];
        p.daoId = daoId;
        p.data = proposalData;
        p.expiry =
            block.timestamp +
            uint256(_daos[daoId].votingPeriodDays) *
            1 days;
        proposalId = _nextProposalId++;
        emit ProposalOpened(proposalId, daoId, p.expiry, proposalData);
        return proposalId;
    }

    /// @inheritdoc IDAOManager
    function vote(uint256 proposalId, bool support) external {
        Proposal storage p = _proposals[proposalId];
        require(block.timestamp <= p.expiry && !p.finalized, "VOTING_CLOSED");
        require(!_hasVoted[proposalId][msg.sender], "ALREADY_VOTED");
        uint256 stake = _memberInvestments[p.daoId][msg.sender];
        uint256 weight = stake == 0 ? 1 : stake;
        _hasVoted[proposalId][msg.sender] = true;
        if (support) p.yesWeight += weight;
        else p.noWeight += weight;
        emit VoteCast(proposalId, msg.sender, support, weight);
    }

    /// @inheritdoc IDAOManager
    function finalizeProposal(uint256 proposalId) external {
        Proposal storage p = _proposals[proposalId];
        require(
            block.timestamp > p.expiry && !p.finalized,
            "NOT_READY_OR_FINALIZED"
        );
        p.finalized = true;
        bool approved = p.yesWeight > p.noWeight;
        emit ProposalFinalized(proposalId, approved, block.timestamp);
    }

    /// @notice Execute a finalized proposal. Supports treasury withdrawals via encoded proposal data.
    /// @dev Proposal data can encode (token, amount, recipient) for treasury operations. Uses ReentrancyGuard.
    function executeProposal(uint256 proposalId) external nonReentrant {
        Proposal storage p = _proposals[proposalId];
        require(p.finalized, "NOT_FINALIZED");
        require(!p.executed, "ALREADY_EXECUTED");
        require(p.yesWeight > p.noWeight, "NOT_APPROVED");
        p.executed = true;
        emit ProposalExecuted(proposalId, block.timestamp);

        if (tokenVault != address(0) && p.data.length == 96) {
            (address token, uint256 amount, address recipient) = abi.decode(
                p.data,
                (address, uint256, address)
            );
            if (token != address(0) && amount > 0 && recipient != address(0)) {
                ITokenVault(tokenVault).withdraw(token, amount);
                bool ok = IERC20(token).transfer(recipient, amount);
                require(ok, "TRANSFER_FAILED");
                emit TreasuryWithdrawn(p.daoId, recipient, token, amount);
            }
        }
    }

    /// @inheritdoc IDAOManager
    function getBnplTerms(
        uint256 daoId
    )
        external
        view
        returns (
            uint256 numInstallments,
            uint256 allowedIntervalMinDays,
            uint256 allowedIntervalMaxDays,
            uint256 lateFeeBps,
            uint256 gracePeriodDays,
            bool rescheduleAllowed,
            uint256 minDownPaymentBps
        )
    {
        BnplTerms storage t = _bnplTerms[daoId];
        return (
            t.numInstallments,
            t.allowedIntervalMinDays,
            t.allowedIntervalMaxDays,
            t.lateFeeBps,
            t.gracePeriodDays,
            t.rescheduleAllowed,
            t.minDownPaymentBps
        );
    }

    /// @inheritdoc IDAOManager
    function setBnplTerms(
        uint256 daoId,
        uint256 numInstallments,
        uint256 allowedIntervalMinDays,
        uint256 allowedIntervalMaxDays,
        uint256 lateFeeBps,
        uint256 gracePeriodDays,
        bool rescheduleAllowed,
        uint256 minDownPaymentBps
    ) external {
        require(_daos[daoId].createdAt != 0, "DAO_NOT_FOUND");
        _bnplTerms[daoId] = BnplTerms({
            numInstallments: numInstallments,
            allowedIntervalMinDays: allowedIntervalMinDays,
            allowedIntervalMaxDays: allowedIntervalMaxDays,
            lateFeeBps: lateFeeBps,
            gracePeriodDays: gracePeriodDays,
            rescheduleAllowed: rescheduleAllowed,
            minDownPaymentBps: minDownPaymentBps
        });
        emit BnplTermsUpdated(
            daoId,
            numInstallments,
            allowedIntervalMinDays,
            allowedIntervalMaxDays,
            lateFeeBps,
            gracePeriodDays,
            rescheduleAllowed
        );
    }

    /// @notice Credit DAO treasury accounting. Callable by privileged funders (BNPLManager, CRE, admin).
    function creditTreasury(uint256 daoId, uint256 amount) external {
        require(_daos[daoId].createdAt != 0, "DAO_NOT_FOUND");
        require(
            hasRole(TREASURY_FUNDER_ROLE, msg.sender) ||
                hasRole(DEFAULT_ADMIN_ROLE, msg.sender),
            "NOT_AUTHORIZED"
        );
        _daos[daoId].treasuryBalance += amount;
        emit TreasuryDeposited(
            daoId,
            msg.sender,
            amount,
            _daos[daoId].treasuryBalance
        );
    }

    /// @notice Returns the DAO's accounting treasury balance.
    function getTreasuryBalance(uint256 daoId) external view returns (uint256) {
        return _daos[daoId].treasuryBalance;
    }
}
