// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {AccessControl} from "openzeppelin-contracts/contracts/access/AccessControl.sol";
import {ReentrancyGuard} from "openzeppelin-contracts/contracts/utils/ReentrancyGuard.sol";
import {ILoanManager} from "./interfaces/ILoanManager.sol";

contract LoanManager is AccessControl, ReentrancyGuard, ILoanManager {
    bytes32 public constant LOAN_ADMIN_ROLE = keccak256("LOAN_ADMIN_ROLE");

    struct Loan {
        uint256 id;
        address borrower;
        uint256 principal;
        uint256 interestRateBps;
        uint256 startTime;
        uint256 endTime;
        uint256 amountPaid;
        /// @notice status: 0 = pending, 1 = approved, 2 = repaid, 3 = defaulted
        uint8 status;
    }

    uint256 private _nextLoanId = 1;
    mapping(uint256 => Loan) private _loans;

    constructor() {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
    }

    /// @inheritdoc ILoanManager
    function createLoan(
        address borrower,
        uint256,
        uint256 principal,
        uint256 interestRateBps,
        uint256 durationSeconds
    ) external returns (uint256 loanId) {
        require(principal > 0, "ZERO_PRINCIPAL");
        require(durationSeconds > 0, "ZERO_DURATION");
        loanId = _nextLoanId++;
        Loan storage l = _loans[loanId];
        l.id = loanId;
        l.borrower = borrower;
        l.principal = principal;
        l.interestRateBps = interestRateBps;
        l.startTime = block.timestamp;
        l.endTime = block.timestamp + durationSeconds;
        l.amountPaid = 0;
        l.status = 0;
        emit LoanCreated(
            loanId,
            borrower,
            principal,
            interestRateBps,
            l.startTime,
            l.endTime
        );
        return loanId;
    }

    /// @inheritdoc ILoanManager
    function approveLoan(uint256 loanId) external {
        Loan storage l = _loans[loanId];
        require(l.id != 0, "LOAN_NOT_FOUND");
        l.status = 1;
        emit LoanApproved(loanId, msg.sender);
    }

    function _accruedInterest(Loan storage l) internal view returns (uint256) {
        if (l.id == 0) return 0;
        uint256 elapsed = block.timestamp > l.startTime
            ? block.timestamp - l.startTime
            : 0;
        uint256 secondsInYear = 365 days;
        return
            (l.principal * l.interestRateBps * elapsed) /
            (10000 * secondsInYear);
    }

    /// @inheritdoc ILoanManager
    function getAccruedInterest(
        uint256 loanId
    ) external view returns (uint256) {
        Loan storage l = _loans[loanId];
        return _accruedInterest(l);
    }

    /// @inheritdoc ILoanManager
    function getAmountOwed(uint256 loanId) external view returns (uint256) {
        Loan storage l = _loans[loanId];
        uint256 accrued = _accruedInterest(l);
        uint256 total = l.principal + accrued;
        if (total <= l.amountPaid) return 0;
        return total - l.amountPaid;
    }

    /// @inheritdoc ILoanManager
    function makePayment(uint256 loanId) external payable nonReentrant {
        Loan storage l = _loans[loanId];
        require(l.id != 0, "LOAN_NOT_FOUND");
        require(l.status == 1, "LOAN_NOT_APPROVED");
        require(msg.value > 0, "ZERO_PAYMENT");

        uint256 accrued = _accruedInterest(l);
        uint256 totalOwed = l.principal + accrued;
        uint256 remaining = totalOwed > l.amountPaid
            ? totalOwed - l.amountPaid
            : 0;
        uint256 toApply = msg.value <= remaining ? msg.value : remaining;

        l.amountPaid += toApply;
        uint256 refund = msg.value - toApply;

        emit PaymentMade(
            loanId,
            msg.sender,
            toApply,
            l.amountPaid,
            remaining > toApply ? remaining - toApply : 0
        );

        if (refund > 0) {
            (bool sent, ) = msg.sender.call{value: refund}("");
            require(sent, "REFUND_FAILED");
        }

        if (l.amountPaid >= totalOwed) l.status = 2;
    }

    /// @inheritdoc ILoanManager
    function markDefaulted(uint256 loanId) external {
        Loan storage l = _loans[loanId];
        require(l.id != 0, "LOAN_NOT_FOUND");
        l.status = 3;
        emit LoanDefaulted(loanId);
    }

    /// @inheritdoc ILoanManager
    function getLoan(
        uint256 loanId
    )
        external
        view
        returns (
            uint256 id,
            address borrower,
            uint256 principal,
            uint256 interestRateBps,
            uint256 startTime,
            uint256 endTime,
            uint256 amountPaid,
            uint8 status
        )
    {
        Loan storage l = _loans[loanId];
        return (
            l.id,
            l.borrower,
            l.principal,
            l.interestRateBps,
            l.startTime,
            l.endTime,
            l.amountPaid,
            l.status
        );
    }
}
