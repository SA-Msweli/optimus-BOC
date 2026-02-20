// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {Test} from "forge-std/Test.sol";
import {BNPLManager} from "../src/BNPLManager.sol";
import {DAOManager} from "../src/DAOManager.sol";

/// @title BNPLManager Edge Case Tests
/// @notice Comprehensive edge case testing for BNPL: extreme amounts, refund sequences, late fees, rescheduling.
contract BNPLManagerEdgeCasesTest is Test {
    BNPLManager bnpl;
    DAOManager dao;
    uint256 daoId;

    function setUp() public {
        bnpl = new BNPLManager();
        dao = new DAOManager();
        daoId = dao.createDAO(address(this), 1, 7);
        dao.setBnplTerms(daoId, 3, 1, 90, 500, 5, true, 0);
        bnpl.setDaoManager(address(dao));
        // allow BNPLManager to credit treasury when applying late fees in tests
        dao.grantRole(dao.TREASURY_FUNDER_ROLE(), address(bnpl));
    }

    receive() external payable {}

    /// @notice Test installment division with very small amounts
    function testEdgeCase_TinyInstallments() public {
        uint256 id = bnpl.createBNPL(
            daoId,
            address(0x99),
            3,
            block.timestamp,
            30 days,
            ""
        );
        (, , , , , , uint256[] memory inst, , , , ) = bnpl.getArrangement(id);
        assertEq(inst[0] + inst[1] + inst[2], 3);
        assertTrue(inst[0] > 0);
        assertTrue(inst[1] > 0);
        assertTrue(inst[2] > 0);
    }

    /// @notice Test with maximum policy interval (90 days)
    function testEdgeCase_MaxInterval() public {
        uint256 id = bnpl.createBNPL(
            daoId,
            address(0x99),
            1000,
            block.timestamp,
            90 days,
            ""
        );
        (, , , , , , , , uint256 interval, , ) = bnpl.getArrangement(id);
        assertEq(interval, 90 days);
    }

    /// @notice Test with minimum policy interval (1 day)
    function testEdgeCase_MinInterval() public {
        uint256 id = bnpl.createBNPL(
            daoId,
            address(0x99),
            1000,
            block.timestamp,
            1 days,
            ""
        );
        (, , , , , , , , uint256 interval, , ) = bnpl.getArrangement(id);
        assertEq(interval, 1 days);
    }

    /// @notice Test multiple refunds in sequence
    function testEdgeCase_MultipleRefundsSequential() public {
        uint256 id = bnpl.createBNPL(
            daoId,
            address(0x99),
            300,
            block.timestamp,
            30 days,
            ""
        );
        (, , , , , , uint256[] memory inst, , , , ) = bnpl.getArrangement(id);

        vm.deal(address(this), 1 ether);

        uint256 beforePayment1 = address(this).balance;
        bnpl.makePayment{value: inst[0] + 100}(id, 0);
        uint256 afterPayment1 = address(this).balance;
        assertEq(beforePayment1 - afterPayment1, inst[0]);

        uint256 beforePayment2 = address(this).balance;
        bnpl.makePayment{value: inst[1] + 50}(id, 1);
        uint256 afterPayment2 = address(this).balance;
        assertEq(beforePayment2 - afterPayment2, inst[1]);

        uint256 beforePayment3 = address(this).balance;
        bnpl.makePayment{value: inst[2] + 75}(id, 2);
        uint256 afterPayment3 = address(this).balance;
        assertEq(beforePayment3 - afterPayment3, inst[2]);
    }

    /// @notice Test late fee application with maximum DAO value
    function testEdgeCase_LateFeeWithHighRate() public {
        uint256 newDaoId = dao.createDAO(address(this), 2, 7);
        dao.setBnplTerms(newDaoId, 3, 1, 90, 10000, 5, true, 0);

        uint256 id = bnpl.createBNPL(
            newDaoId,
            address(0x99),
            1000,
            block.timestamp,
            30 days,
            ""
        );
        (, , , , uint256 originalTotal, , uint256[] memory inst, , , , ) = bnpl
            .getArrangement(id);

        uint256 expectedMaxFee = (inst[0] * 10000) / 10000;
        bnpl.applyLateFee(id, 0);

        (, , , , uint256 newTotal, , , , , , ) = bnpl.getArrangement(id);
        assertEq(newTotal, originalTotal + expectedMaxFee);
    }

    /// @notice Test rescheduling to different intervals sequentially
    function testEdgeCase_RescheduleMultipleTimes() public {
        uint256 id = bnpl.createBNPL(
            daoId,
            address(0x99),
            1000,
            block.timestamp,
            30 days,
            ""
        );

        uint256 newStart1 = block.timestamp + 5 days;
        bnpl.reschedule(id, newStart1, 25 days);
        (, , , , , , , uint256 start1, uint256 interval1, , ) = bnpl
            .getArrangement(id);
        assertEq(start1, newStart1);
        assertEq(interval1, 25 days);

        uint256 newStart2 = block.timestamp + 10 days;
        bnpl.reschedule(id, newStart2, 35 days);
        (, , , , , , , uint256 start2, uint256 interval2, , ) = bnpl
            .getArrangement(id);
        assertEq(start2, newStart2);
        assertEq(interval2, 35 days);
    }

    /// @notice Test payment completion edge case with non-uniform installments
    function testEdgeCase_NonUniformInstallmentCompletion() public {
        uint256 totalAmount = 1005;
        uint256 id = bnpl.createBNPL(
            daoId,
            address(0x99),
            totalAmount,
            block.timestamp,
            30 days,
            ""
        );
        (, , , , , , uint256[] memory inst, , , , ) = bnpl.getArrangement(id);

        uint256 total = inst[0] + inst[1] + inst[2];
        assertEq(total, totalAmount);

        vm.deal(address(this), totalAmount + 100);
        for (uint8 i = 0; i < 3; i++) {
            bnpl.makePayment{value: inst[i]}(id, i);
        }

        (, , , , , , , , , , uint8 finalStatus) = bnpl.getArrangement(id);
        assertEq(finalStatus, 2);
    }

    /// @notice Test payment with exactly one wei under requirement
    function testEdgeCase_PaymentOneWeiShort() public {
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

    /// @notice Test boundary: activate from pending without activation function
    function testEdgeCase_FirstPaymentAutoActivation() public {
        uint256 id = bnpl.createBNPL(
            daoId,
            address(0x99),
            1000,
            block.timestamp,
            30 days,
            ""
        );
        (, , , , , , uint256[] memory inst, , , , uint8 statusBefore) = bnpl
            .getArrangement(id);
        assertEq(statusBefore, 0);

        vm.deal(address(this), 1 ether);
        bnpl.makePayment{value: inst[0]}(id, 0);

        (, , , , , , , , , , uint8 statusAfter) = bnpl.getArrangement(id);
        assertEq(statusAfter, 1);
    }

    /// @notice Test late fee accumulation on multiple installments
    function testEdgeCase_MultipleLateFeesAccumulate() public {
        uint256 id = bnpl.createBNPL(
            daoId,
            address(0x99),
            9000,
            block.timestamp,
            30 days,
            ""
        );
        (, , , , uint256 originalTotal, , uint256[] memory inst, , , , ) = bnpl
            .getArrangement(id);

        uint256 fee0 = (inst[0] * 500) / 10000;
        uint256 fee1 = (inst[1] * 500) / 10000;

        bnpl.applyLateFee(id, 0);
        bnpl.applyLateFee(id, 1);

        (, , , , uint256 newTotal, , , , , , ) = bnpl.getArrangement(id);
        assertEq(newTotal, originalTotal + fee0 + fee1);
    }

    /// @notice Test mixed payment and late fee scenarios
    function testEdgeCase_PaymentAfterLateFee() public {
        uint256 id = bnpl.createBNPL(
            daoId,
            address(0x99),
            1500,
            block.timestamp,
            30 days,
            ""
        );
        (, , , , , , uint256[] memory inst, , , , ) = bnpl.getArrangement(id);

        bnpl.applyLateFee(id, 0);

        vm.deal(address(this), 1 ether);
        bnpl.makePayment{value: inst[0]}(id, 0);

        (, , , , , , , , , , uint8 status) = bnpl.getArrangement(id);
        assertEq(status, 1);
    }
}
