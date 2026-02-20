// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {Test} from "forge-std/Test.sol";
import {TokenVault} from "../src/TokenVault.sol";
import {DAOManager} from "../src/DAOManager.sol";
import {BNPLManager} from "../src/BNPLManager.sol";
import {LoanManager} from "../src/LoanManager.sol";
import {DIDRegistry} from "../src/DIDRegistry.sol";

contract DeploymentTest is Test {
    TokenVault vault;
    DAOManager dao;
    BNPLManager bnpl;
    LoanManager loan;
    DIDRegistry did;

    function setUp() public {
        vault = new TokenVault();
        dao = new DAOManager();
        bnpl = new BNPLManager();
        loan = new LoanManager();
        did = new DIDRegistry();

        // wire as deployment script should
        dao.setTokenVault(address(vault));
        bnpl.setDaoManager(address(dao));
        vault.grantRole(vault.VAULT_MANAGER_ROLE(), address(dao));
        dao.grantRole(dao.TREASURY_FUNDER_ROLE(), address(bnpl));
    }

    function test_roles_and_wiring_after_deploy() public view {
        // TokenVault should allow DAOManager as VAULT_MANAGER_ROLE
        assertTrue(vault.hasRole(vault.VAULT_MANAGER_ROLE(), address(dao)));

        // DAOManager should have TokenVault address set
        assertEq(dao.tokenVault(), address(vault));

        // BNPLManager should be wired to DAOManager
        assertEq(bnpl.daoManager(), address(dao));

        // DAOManager should have TREASURY_FUNDER_ROLE granted to BNPLManager
        assertTrue(dao.hasRole(dao.TREASURY_FUNDER_ROLE(), address(bnpl)));

        // LoanManager & DIDRegistry should be deployable and grant DEFAULT_ADMIN_ROLE to deployer
        assertTrue(loan.hasRole(loan.DEFAULT_ADMIN_ROLE(), address(this)));
        assertTrue(did.hasRole(did.DEFAULT_ADMIN_ROLE(), address(this)));

        // Deployer should be DEFAULT_ADMIN_ROLE on vault & dao as well
        assertTrue(vault.hasRole(vault.DEFAULT_ADMIN_ROLE(), address(this)));
        assertTrue(dao.hasRole(dao.DEFAULT_ADMIN_ROLE(), address(this)));
    }
}
