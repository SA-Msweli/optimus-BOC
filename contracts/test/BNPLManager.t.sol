// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {Test} from "forge-std/Test.sol";
import {BNPLManager} from "../src/BNPLManager.sol";
import {DAOManager} from "../src/DAOManager.sol";

contract BNPLManagerTest is Test {
    BNPLManager bnpl;
    DAOManager dao;
    uint256 daoId;

    function setUp() public {
        bnpl = new BNPLManager();
        dao = new DAOManager();
        daoId = dao.createDAO(address(this), 1, 7);
        dao.setBnplTerms(daoId, 3, 1, 90, 500, 5, true, 0);
        bnpl.setDaoManager(address(dao));
        // Allow BNPLManager to credit DAO treasury when late fees are applied
        dao.grantRole(dao.TREASURY_FUNDER_ROLE(), address(bnpl));
    }

    receive() external payable {}

    function testCreateBNPL_InstallmentsSum() public {
        uint256 id = bnpl.createBNPL(
            daoId,
            address(0x99),
            1000,
            block.timestamp,
            30 days,
            ""
        );
        (, , , , , , uint256[] memory inst, , , , ) = bnpl.getArrangement(id);
        assertEq(
            inst[0] + inst[1] + inst[2],
            1000,
            "Installments sum to total"
        );
    }

    function testCreateBNPL_MultipleAmounts() public {
        for (uint256 amt = 1001; amt <= 1003; amt++) {
            uint256 id = bnpl.createBNPL(
                daoId,
                address(0x99),
                amt,
                block.timestamp,
                30 days,
                ""
            );
            (, , , , , , uint256[] memory inst, , , , ) = bnpl.getArrangement(
                id
            );
            assertEq(inst[0] + inst[1] + inst[2], amt);
        }
    }

    function testCreateBNPL_EnforcesPolicy() public {
        vm.expectRevert();
        bnpl.createBNPL(
            daoId,
            address(0x99),
            1000,
            block.timestamp,
            100 days,
            ""
        );
    }

    function testCreateBNPL_RequiresPolicy() public {
        uint256 newDaoId = dao.createDAO(address(0x5), 2, 7);
        vm.expectRevert("DAO_BNPL_NOT_CONFIGURED");
        bnpl.createBNPL(
            newDaoId,
            address(0x99),
            1000,
            block.timestamp,
            30 days,
            ""
        );
    }

    function testMakePayment_FirstActivates() public {
        uint256 id = bnpl.createBNPL(
            daoId,
            address(0x99),
            1000,
            block.timestamp,
            30 days,
            ""
        );
        (, , , , , , uint256[] memory inst, , , , uint8 status0) = bnpl
            .getArrangement(id);
        assertEq(status0, 0);

        vm.deal(address(this), 1 ether);
        bnpl.makePayment{value: inst[0]}(id, 0);

        (, , , , , , , , , , uint8 status1) = bnpl.getArrangement(id);
        assertEq(status1, 1);
    }

    function testMakePayment_RefundsExcess() public {
        uint256 id = bnpl.createBNPL(
            daoId,
            address(0x99),
            1000,
            block.timestamp,
            30 days,
            ""
        );
        (, , , , , , uint256[] memory inst, , , , ) = bnpl.getArrangement(id);

        vm.deal(address(this), 1 ether);
        uint256 before = address(this).balance;
        bnpl.makePayment{value: inst[0] + 50}(id, 0);
        uint256 after_balance = address(this).balance;

        assertEq(before - after_balance, inst[0]);
    }

    function testMakePayment_RejectsInsufficient() public {
        uint256 id = bnpl.createBNPL(
            daoId,
            address(0x99),
            1000,
            block.timestamp,
            30 days,
            ""
        );
        (, , , , , , uint256[] memory inst, , , , ) = bnpl.getArrangement(id);

        vm.deal(address(this), 1 ether);
        vm.expectRevert("INSUFFICIENT_PAYMENT");
        bnpl.makePayment{value: inst[0] - 1}(id, 0);
    }

    function testMakePayment_SequentialToComplete() public {
        uint256 id = bnpl.createBNPL(
            daoId,
            address(0x99),
            900,
            block.timestamp,
            30 days,
            ""
        );
        (, , , , , , uint256[] memory inst, , , , ) = bnpl.getArrangement(id);
        vm.deal(address(this), 1 ether);

        for (uint8 i = 0; i < 3; i++) {
            bnpl.makePayment{value: inst[i]}(id, i);
        }

        (, , , , , , , , , , uint8 status) = bnpl.getArrangement(id);
        assertEq(status, 2);
    }

    function testActivateBNPL_NeedsFirstPayment() public {
        uint256 id = bnpl.createBNPL(
            daoId,
            address(0x99),
            1000,
            block.timestamp,
            30 days,
            ""
        );
        vm.expectRevert("FIRST_PAYMENT_REQUIRED");
        bnpl.activateBNPL(id);
    }

    function testApplyLateFee() public {
        uint256 id = bnpl.createBNPL(
            daoId,
            address(0x99),
            10000,
            block.timestamp,
            30 days,
            ""
        );
        (, , , , uint256 before, , uint256[] memory inst, , , , ) = bnpl
            .getArrangement(id);

        uint256 treasuryBefore = dao.getTreasuryBalance(daoId);
        uint256 expectedFee = (inst[0] * 500) / 10000;

        bnpl.applyLateFee(id, 0);

        (, , , , uint256 after_total, , , , , , ) = bnpl.getArrangement(id);
        uint256 treasuryAfter = dao.getTreasuryBalance(daoId);

        assertEq(after_total, before + expectedFee);
        assertEq(treasuryAfter, treasuryBefore + expectedFee);
    }

    function testReschedule() public {
        uint256 id = bnpl.createBNPL(
            daoId,
            address(0x99),
            1000,
            block.timestamp,
            30 days,
            ""
        );
        uint256 newStart = block.timestamp + 1 days;
        uint256 newInterval = 31 days;

        bnpl.reschedule(id, newStart, newInterval);
        (, , , , , , , uint256 start, uint256 interval, , ) = bnpl
            .getArrangement(id);

        assertEq(start, newStart);
        assertEq(interval, newInterval);
    }

    function testReschedule_NotAllowedReverts() public {
        uint256 newDaoId = dao.createDAO(address(this), 2, 7);
        dao.setBnplTerms(newDaoId, 3, 1, 90, 500, 5, false, 0);
        bnpl.setDaoManager(address(dao));

        uint256 id = bnpl.createBNPL(
            newDaoId,
            address(0x99),
            1000,
            block.timestamp,
            30 days,
            ""
        );

        vm.expectRevert("RESCHEDULE_NOT_ALLOWED");
        bnpl.reschedule(id, block.timestamp + 1 days, 31 days);
    }

    function testReschedule_UnauthorizedReverts() public {
        uint256 id = bnpl.createBNPL(
            daoId,
            address(0x99),
            1000,
            block.timestamp,
            30 days,
            ""
        );

        vm.prank(address(0x5));
        vm.expectRevert("UNAUTHORIZED");
        bnpl.reschedule(id, block.timestamp + 1 days, 31 days);
    }

    function testMakePayment_InvalidInstallment() public {
        uint256 id = bnpl.createBNPL(
            daoId,
            address(0x99),
            1000,
            block.timestamp,
            30 days,
            ""
        );
        vm.deal(address(this), 1 ether);
        vm.expectRevert("INVALID_INSTALLMENT");
        bnpl.makePayment{value: 100}(id, 10);
    }

    function testMakePayment_NotFound() public {
        vm.expectRevert("ARR_NOT_FOUND");
        bnpl.makePayment{value: 100}(9999, 0);
    }
}
