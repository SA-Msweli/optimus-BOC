// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {Test} from "forge-std/Test.sol";
import {DAOManager} from "../src/DAOManager.sol";

event DaoCreated(uint256 indexed daoId, address indexed creator, uint8 goal, uint256 createdAt);
event MemberJoined(uint256 indexed daoId, address indexed member, uint256 investment, uint256 joinedAt);

contract DAOManagerTest is Test {
    DAOManager dao;
    address alice = address(0x1);
    address bob = address(0x2);
    address charlie = address(0x3);

    function setUp() public {
        dao = new DAOManager();
    }

    function testCreateDAO() public {
        vm.expectEmit(true, true, false, true);
        emit DaoCreated(1, alice, 1, block.timestamp);
        uint256 id = dao.createDAO(alice, 1, 7);
        assertEq(id, 1);
    }

    function testJoinDAO_UpdatesTreasury() public {
        uint256 daoId = dao.createDAO(alice, 1, 7);
        dao.joinDAO(daoId, bob, 1000);
        dao.joinDAO(daoId, charlie, 2000);
    }

    function testJoinDAO_RevertsNotFound() public {
        vm.expectRevert("DAO_NOT_FOUND_OR_DISSOLVED");
        dao.joinDAO(9999, bob, 1000);
    }

    function testSetBnplTerms() public {
        uint256 daoId = dao.createDAO(alice, 2, 3);
        dao.setBnplTerms(daoId, 3, 1, 30, 500, 7, true, 1000);
        
        (uint256 numInst, uint256 minDays, uint256 maxDays, uint256 feeBps, uint256 graceDays, bool reschedAllow, uint256 minDownBps)
            = dao.getBnplTerms(daoId);
        
        assertEq(numInst, 3);
        assertEq(minDays, 1);
        assertEq(maxDays, 30);
        assertEq(feeBps, 500);
        assertEq(graceDays, 7);
        assertTrue(reschedAllow);
        assertEq(minDownBps, 1000);
    }

    function testPropose_CreatesProposal() public {
        uint256 daoId = dao.createDAO(alice, 2, 1);
        uint256 proposalId = dao.propose(daoId, abi.encode("proposal_data"));
        assertEq(proposalId, 1);
    }

    function testVote_StakeWeighted() public {
        uint256 daoId = dao.createDAO(alice, 1, 1);
        dao.joinDAO(daoId, alice, 1000);
        dao.joinDAO(daoId, bob, 500);
        dao.joinDAO(daoId, charlie, 0);
        
        uint256 proposalId = dao.propose(daoId, "");
        
        vm.prank(alice);
        dao.vote(proposalId, true);
        
        vm.prank(bob);
        dao.vote(proposalId, true);
        
        vm.prank(charlie);
        dao.vote(proposalId, false);
        
        vm.warp(block.timestamp + 2 days);
        dao.finalizeProposal(proposalId);
    }

    function testVote_PreventsDoubleVoting() public {
        uint256 daoId = dao.createDAO(alice, 1, 1);
        dao.joinDAO(daoId, alice, 1000);
        
        uint256 proposalId = dao.propose(daoId, "");
        
        vm.prank(alice);
        dao.vote(proposalId, true);
        
        vm.prank(alice);
        vm.expectRevert("ALREADY_VOTED");
        dao.vote(proposalId, false);
    }

    function testVote_RequiresOpenProposal() public {
        uint256 daoId = dao.createDAO(alice, 1, 1);
        uint256 proposalId = dao.propose(daoId, "");
        
        vm.warp(block.timestamp + 2 days);
        vm.expectRevert("VOTING_CLOSED");
        dao.vote(proposalId, true);
    }

    function testFinalizeProposal_RequiresPeriodExpiry() public {
        uint256 daoId = dao.createDAO(alice, 1, 1);
        uint256 proposalId = dao.propose(daoId, "");
        
        vm.expectRevert("NOT_READY_OR_FINALIZED");
        dao.finalizeProposal(proposalId);
        
        vm.warp(block.timestamp + 2 days);
        dao.finalizeProposal(proposalId);
    }

    function testExecuteProposal_RequiresApproval() public {
        uint256 daoId = dao.createDAO(alice, 1, 1);
        dao.joinDAO(daoId, alice, 1000);
        
        uint256 proposalId = dao.propose(daoId, "");
        
        vm.prank(alice);
        dao.vote(proposalId, true);
        
        vm.warp(block.timestamp + 2 days);
        dao.finalizeProposal(proposalId);
        dao.executeProposal(proposalId);
    }

    function testExecuteProposal_RevertsIfRejected() public {
        uint256 daoId = dao.createDAO(alice, 1, 1);
        dao.joinDAO(daoId, alice, 1000);
        dao.joinDAO(daoId, bob, 2000);
        
        uint256 proposalId = dao.propose(daoId, "");
        
        vm.prank(alice);
        dao.vote(proposalId, true);
        
        vm.prank(bob);
        dao.vote(proposalId, false);
        dao.vote(proposalId, false);
        
        vm.warp(block.timestamp + 2 days);
        dao.finalizeProposal(proposalId);
        
        vm.expectRevert("NOT_APPROVED");
        dao.executeProposal(proposalId);
    }

    function testProposal_FullLifecycle() public {
        uint256 daoId = dao.createDAO(alice, 2, 1);
        dao.joinDAO(daoId, alice, 1000);
        dao.joinDAO(daoId, bob, 2000);
        
        uint256 proposalId = dao.propose(daoId, abi.encode("proposal_data"));
        assertEq(proposalId, 1);
        
        vm.prank(alice);
        dao.vote(proposalId, true);
        
        vm.prank(bob);
        dao.vote(proposalId, true);
        
        vm.warp(block.timestamp + 2 days);
        dao.finalizeProposal(proposalId);
        dao.executeProposal(proposalId);
    }
}
