// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

/// @notice Simple TokenVault interface used by Optimus for deposits/withdrawals.
interface ITokenVault {
    event Deposit(
        address indexed from,
        address indexed token,
        uint256 amount,
        uint256 newBalance
    );
    event Withdrawal(
        address indexed to,
        address indexed token,
        uint256 amount,
        uint256 newBalance
    );
    event TreasuryWithdrawn(
        uint256 indexed daoId,
        address to,
        address token,
        uint256 amount
    );

    function deposit(address token, uint256 amount) external;

    function withdraw(address token, uint256 amount) external;

    function getBalance(address token) external view returns (uint256);
}
