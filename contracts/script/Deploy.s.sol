// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "forge-std/Script.sol";
import "src/TokenVault.sol";
import "src/DAOManager.sol";
import "src/BNPLManager.sol";
import "src/LoanManager.sol";
import "src/DIDRegistry.sol";

/// @notice Simple Foundry deployment script that deploys core Optimus contracts
/// and wires treasury/roles required for DAO<->Vault<->BNPL interactions.
contract Deploy is Script {
    function run() external returns (address[] memory addrs) {
        uint256 pk = vm.envUint("DEPLOYER_KEY");
        vm.startBroadcast(pk);

        TokenVault vault = new TokenVault();
        DAOManager dao = new DAOManager();
        BNPLManager bnpl = new BNPLManager();
        LoanManager loan = new LoanManager();
        DIDRegistry did = new DIDRegistry();

        // wire vault <> dao
        dao.setTokenVault(address(vault));

        // wire dao <> bnpl
        bnpl.setDaoManager(address(dao));

        // grant roles so DAOManager can pull from vault and BNPLManager can credit treasury
        vault.grantRole(vault.VAULT_MANAGER_ROLE(), address(dao));
        dao.grantRole(dao.TREASURY_FUNDER_ROLE(), address(bnpl));

        // Log deployed addresses for convenience
        console.log("TokenVault:", address(vault));
        console.log("DAOManager:", address(dao));
        console.log("BNPLManager:", address(bnpl));
        console.log("LoanManager:", address(loan));
        console.log("DIDRegistry:", address(did));

        vm.stopBroadcast();

        addrs = new address[](5);
        addrs[0] = address(vault);
        addrs[1] = address(dao);
        addrs[2] = address(bnpl);
        addrs[3] = address(loan);
        addrs[4] = address(did);
    }
}
