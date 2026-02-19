// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {ITokenVault} from "./interfaces/ITokenVault.sol";
import {AccessControl} from "openzeppelin-contracts/contracts/access/AccessControl.sol";
import {ReentrancyGuard} from "openzeppelin-contracts/contracts/utils/ReentrancyGuard.sol";
import {IERC20} from "openzeppelin-contracts/contracts/token/ERC20/IERC20.sol";

/// @title TokenVault
/// @notice Multi‑asset vault with ERC‑20 safe transfers, access control, and reentrancy protection.
contract TokenVault is AccessControl, ReentrancyGuard, ITokenVault {
    mapping(address => uint256) private _balances;

    /// @dev Grant DEFAULT_ADMIN_ROLE to deployer.
    constructor() {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
    }

    /// @notice Deposit increases the vault accounting balance for `token`.
    function deposit(address token, uint256 amount) external override {
        require(amount > 0, "ZERO_AMOUNT");
        bool ok = IERC20(token).transferFrom(msg.sender, address(this), amount);
        require(ok, "ERC20_TRANSFER_FROM_FAILED");
        _balances[token] += amount;
        emit Deposit(msg.sender, token, amount, _balances[token]);
    }

    /// @notice Withdraw reduces the vault accounting balance for `token`.
    function withdraw(
        address token,
        uint256 amount
    ) external override onlyRole(DEFAULT_ADMIN_ROLE) nonReentrant {
        require(_balances[token] >= amount, "INSUFFICIENT_BALANCE");
        _balances[token] -= amount;
        bool ok = IERC20(token).transfer(msg.sender, amount);
        require(ok, "ERC20_TRANSFER_FAILED");
        emit Withdrawal(msg.sender, token, amount, _balances[token]);
    }

    /// @notice Returns the stored balance (accounting) for `token`.
    function getBalance(
        address token
    ) external view override returns (uint256) {
        return _balances[token];
    }
}
