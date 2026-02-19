// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {Test} from "forge-std/Test.sol";
import "../src/DIDRegistry.sol";

event DIDCreated(address indexed owner, uint256 createdAt);
event PrivyCredentialLinked(address indexed owner, bytes32 indexed credentialHash, uint256 linkedAt);
event RiskProfileUpdated(address indexed owner, uint256 newScore, bytes32 indexed profileHash);

/// @title DIDRegistryTest
/// @notice Unit tests for DIDRegistry covering lifecycle, Privy linkage and risk scoring.
/// @dev Mirrors AccessControl-style tests: role grants, permission checks and expected reverts.
contract DIDRegistryTest is Test {
    DIDRegistry internal registry;
    address internal admin = address(this);
    address internal owner = address(0x1);
    address internal other = address(0x2);
    address internal updater = address(0x3);
    address internal privyLinker = address(0x4);

    /// @dev Deploy DIDRegistry with deployer as DEFAULT_ADMIN_ROLE.
    function setUp() public {
        registry = new DIDRegistry();
    }

    /// @notice Owner can create their DID; existence is stored.
    function testCreateDID_byOwner() public {
        vm.prank(owner);
        vm.expectEmit(true, false, false, false);
        emit DIDCreated(owner, 0);
        registry.createDID(owner);
        assertTrue(registry.exists(owner));
    }

    /// @notice Admin may create a DID for another account.
    function testCreateDID_byAdmin() public {
        vm.expectEmit(true, false, false, false);
        emit DIDCreated(owner, 0);
        registry.createDID(owner);
        assertTrue(registry.exists(owner));
    }

    /// @notice Duplicate creation reverts with DIDAlreadyExists.
    function testCreateDID_duplicateReverts() public {
        registry.createDID(owner);
        vm.expectRevert(abi.encodeWithSelector(DIDAlreadyExists.selector, owner));
        registry.createDID(owner);
    }

    /// @notice Owner can link a Privy credential; event emitted.
    function testLinkPrivyCredential_byOwner() public {
        registry.createDID(owner);
        bytes32 h = keccak256(abi.encodePacked("cred-1"));
        vm.prank(owner);
        vm.expectEmit(true, true, false, false);
        emit PrivyCredentialLinked(owner, h, 0);
        registry.linkPrivyCredential(owner, h);
        assertEq(registry.getPrivyCredentialHash(owner), h);
    }

    /// @notice Account with PRIVY_LINKER_ROLE may link credentials for others.
    function testLinkPrivyCredential_byRole() public {
        registry.createDID(owner);
        registry.grantRole(registry.PRIVY_LINKER_ROLE(), privyLinker);
        bytes32 h = keccak256(abi.encodePacked("cred-2"));
        vm.prank(privyLinker);
        vm.expectEmit(true, true, false, false);
        emit PrivyCredentialLinked(owner, h, 0);
        registry.linkPrivyCredential(owner, h);
        assertEq(registry.getPrivyCredentialHash(owner), h);
    }

    /// @notice Unauthorized callers cannot link credentials.
    function testLinkPrivyCredential_unauthorizedReverts() public {
        registry.createDID(owner);
        bytes32 h = keccak256(abi.encodePacked("cred-3"));
        vm.prank(other);
        vm.expectRevert();
        registry.linkPrivyCredential(owner, h);
    }

    /// @notice RISK_UPDATER_ROLE may update risk scores and emit the event.
    function testUpdateRiskProfile_byRole() public {
        registry.createDID(owner);
        registry.grantRole(registry.RISK_UPDATER_ROLE(), updater);
        vm.prank(updater);
        vm.expectEmit(true, true, false, false);
        emit RiskProfileUpdated(owner, 420, bytes32(keccak256("hash")));
        registry.updateRiskProfile(owner, 420, bytes32(keccak256("hash")));
        assertEq(registry.getRiskProfileScore(owner), 420);
    }

    /// @notice Scores outside allowed bounds revert with InvalidScore.
    function testUpdateRiskProfile_boundsRevert() public {
        registry.createDID(owner);
        registry.grantRole(registry.RISK_UPDATER_ROLE(), updater);
        vm.prank(updater);
        vm.expectRevert(abi.encodeWithSelector(InvalidScore.selector, uint256(10001)));
        registry.updateRiskProfile(owner, 10001, bytes32(0));
    }

    /// @notice Querying a non-existent DID reverts with DIDNotFound.
    function testGetters_nonexistentRevert() public {
        vm.expectRevert(abi.encodeWithSelector(DIDNotFound.selector, owner));
        registry.getPrivyCredentialHash(owner);
    }
}
