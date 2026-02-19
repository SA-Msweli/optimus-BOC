// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {Test} from "forge-std/Test.sol";
import {ERC20} from "openzeppelin-contracts/contracts/token/ERC20/ERC20.sol";
import {TokenVault} from "../src/TokenVault.sol";

contract ERC20Mock is ERC20 {
    constructor(uint256 initial) ERC20("Mock", "MCK") {
        _mint(msg.sender, initial);
    }
}

contract TokenVaultTest is Test {
    TokenVault vault;
    ERC20Mock token;

    function setUp() public {
        vault = new TokenVault();
        token = new ERC20Mock(10000);
    }

    function testDepositWithdraw() public {
        token.approve(address(vault), 1000);
        vault.deposit(address(token), 1000);
        assertEq(vault.getBalance(address(token)), 1000);
        vault.withdraw(address(token), 500);
        assertEq(vault.getBalance(address(token)), 500);
    }
}
