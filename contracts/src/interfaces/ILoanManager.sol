// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

/// @notice Loan manager interface (create, approve, repay, default).
interface ILoanManager {
    event LoanCreated(
        uint256 indexed loanId,
        address indexed borrower,
        uint256 principal,
        uint256 interestRateBps,
        uint256 startTime,
        uint256 endTime
    );
    event LoanApproved(uint256 indexed loanId, address indexed approver);
    event PaymentMade(
        uint256 indexed loanId,
        address indexed payer,
        uint256 amount,
        uint256 amountPaid,
        uint256 remaining
    );
    event LoanDefaulted(uint256 indexed loanId);

    function createLoan(
        address borrower,
        uint256 daoId,
        uint256 principal,
        uint256 interestRateBps,
        uint256 durationSeconds
    ) external returns (uint256 loanId);

    function approveLoan(uint256 loanId) external;

    function makePayment(uint256 loanId) external payable;

    function markDefaulted(uint256 loanId) external;

    /// @notice Returns the current accrued interest (since loan start) for `loanId`.
    function getAccruedInterest(uint256 loanId) external view returns (uint256);

    /// @notice Returns the total amount currently owed (principal + accrued interest - paid).
    function getAmountOwed(uint256 loanId) external view returns (uint256);

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
        );
}
