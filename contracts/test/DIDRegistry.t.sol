// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {Test} from "forge-std/Test.sol";
import "../src/DIDRegistry.sol";

event DIDCreated(address indexed owner, uint256 createdAt);
event PrivyCredentialLinked(address indexed owner, bytes32 indexed credentialHash, uint256 linkedAt);
event RiskProfileUpdated(address indexed owner, uint256 newScore, bytes32 indexed profileHash);

contract DIDRegistryTest is Test {
    DIDRegistry internal registry;
    address internal admin = address(this);
    address internal owner = address(0x1);
    address internal other = address(0x2);
    address internal updater = address(0x3);
    address internal privyLinker = address(0x4);

    function setUp() public {
        registry = new DIDRegistry();
    }

    function testCreateDID_byOwner() public {
        vm.prank(owner);
        vm.expectEmit(true, false, false, false);
        emit DIDCreated(owner, 0);
        registry.createDID(owner);
        assertTrue(registry.exists(owner));
    }

    function testCreateDID_byAdmin() public {
        vm.expectEmit(true, false, false, false);
        emit DIDCreated(owner, 0);
        registry.createDID(owner);
        assertTrue(registry.exists(owner));
    }

    function testCreateDID_duplicateReverts() public {
        registry.createDID(owner);
        vm.expectRevert(abi.encodeWithSelector(DIDAlreadyExists.selector, owner));
        registry.createDID(owner);
    }

    function testLinkPrivyCredential_byOwner() public {
        registry.createDID(owner);
        bytes32 h = keccak256(abi.encodePacked("cred-1"));
        vm.prank(owner);
        vm.expectEmit(true, true, false, false);
        emit PrivyCredentialLinked(owner, h, 0);
        registry.linkPrivyCredential(owner, h);
        assertEq(registry.getPrivyCredentialHash(owner), h);
    }

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

    function testLinkPrivyCredential_unauthorizedReverts() public {
        registry.createDID(owner);
        bytes32 h = keccak256(abi.encodePacked("cred-3"));
        vm.prank(other);
        vm.expectRevert();
        registry.linkPrivyCredential(owner, h);
    }

    function testUpdateRiskProfile_byRole() public {
        registry.createDID(owner);
        registry.grantRole(registry.RISK_UPDATER_ROLE(), updater);
        vm.prank(updater);
        vm.expectEmit(true, true, false, false);
        emit RiskProfileUpdated(owner, 420, bytes32(keccak256("hash")));
        registry.updateRiskProfile(owner, 420, bytes32(keccak256("hash")));
        assertEq(registry.getRiskProfileScore(owner), 420);
    }

    function testUpdateRiskProfile_boundsRevert() public {
        registry.createDID(owner);
        registry.grantRole(registry.RISK_UPDATER_ROLE(), updater);
        vm.prank(updater);
        vm.expectRevert(abi.encodeWithSelector(InvalidScore.selector, uint256(10001)));
        registry.updateRiskProfile(owner, 10001, bytes32(0));
    }

    function testGetters_nonexistentRevert() public {
        vm.expectRevert(abi.encodeWithSelector(DIDNotFound.selector, owner));
        registry.getPrivyCredentialHash(owner);
    }
}
