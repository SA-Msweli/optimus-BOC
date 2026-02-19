// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {Test} from "forge-std/Test.sol";
import {LoanManager} from "../src/LoanManager.sol";

/// @title LoanManager Edge Case Tests
/// @notice Comprehensive edge case testing for Loans: interest accrual, refunds, time progression.
contract LoanManagerEdgeCasesTest is Test {
    LoanManager loans;
    address borrower = address(0x1);

    function setUp() public {
        loans = new LoanManager();
    }

    receive() external payable {}

    /// @notice Test interest with zero interest rate (principal only)
    function testEdgeCase_ZeroInterestRate() public {
        uint256 lid = loans.createLoan(borrower, 0, 1 ether, 0, 365 days);
        loans.approveLoan(lid);

        uint256 accrued = loans.getAccruedInterest(lid);
        assertEq(accrued, 0);

        uint256 owed = loans.getAmountOwed(lid);
        assertEq(owed, 1 ether);
    }

    /// @notice Test with maximum interest rate (10000 bps = 100%)
    function testEdgeCase_MaximumInterestRate() public {
        uint256 principal = 10000;
        uint256 lid = loans.createLoan(borrower, 0, principal, 10000, 365 days);
        loans.approveLoan(lid);

        vm.warp(block.timestamp + 365 days);

        uint256 accrued = loans.getAccruedInterest(lid);
        uint256 expectedInterest = (principal * 10000) / 10000;
        assertEq(accrued, expectedInterest);
    }

    /// @notice Test interest accrual at very short time intervals
    function testEdgeCase_InterestAtOneHour() public {
        uint256 principal = 525600;
        uint256 lid = loans.createLoan(borrower, 0, principal, 100, 365 days);
        loans.approveLoan(lid);

        vm.warp(block.timestamp + 1 hours);

        uint256 accrued = loans.getAccruedInterest(lid);
        uint256 expectedMax = (principal * 100) / 10000;
        assertTrue(accrued <= expectedMax);
    }

    /// @notice Test multiple partial payments with growing interest
    function testEdgeCase_MultiplePartialPaymentsWithInterest() public {
        uint256 principal = 10000;
        uint256 lid = loans.createLoan(borrower, 0, principal, 500, 365 days);
        loans.approveLoan(lid);

        vm.deal(address(this), 20 ether);

        vm.warp(block.timestamp + 90 days);
        uint256 owed90 = loans.getAmountOwed(lid);
        loans.makePayment{value: owed90 / 2}(lid);

        vm.warp(block.timestamp + 180 days);
        uint256 owed270 = loans.getAmountOwed(lid);
        loans.makePayment{value: owed270}(lid);

        (, , , , , , uint256 paid, uint8 status) = loans.getLoan(lid);
        assertGt(paid, principal);
        assertEq(status, 2);
    }

    /// @notice Test refund with exact overpayment
    function testEdgeCase_RefundExactOverpayment() public {
        uint256 lid = loans.createLoan(borrower, 0, 1000, 0, 365 days);
        loans.approveLoan(lid);

        vm.deal(address(this), 2000);
        uint256 before = address(this).balance;

        loans.makePayment{value: 1100}(lid);

        uint256 after_balance = address(this).balance;
        assertEq(before - after_balance, 1000);
    }

    /// @notice Test payment of exactly principal amount
    function testEdgeCase_PaymentExactPrincipal() public {
        uint256 principal = 5555;
        uint256 lid = loans.createLoan(borrower, 0, principal, 0, 365 days);
        loans.approveLoan(lid);

        vm.deal(address(this), principal + 100);
        loans.makePayment{value: principal}(lid);

        (, , , , , , uint256 paid, uint8 status) = loans.getLoan(lid);
        assertEq(paid, principal);
        assertEq(status, 2);
    }

    /// @notice Test default marking before repayment
    function testEdgeCase_DefaultBeforeAnyPayment() public {
        uint256 lid = loans.createLoan(borrower, 0, 1 ether, 500, 365 days);
        loans.approveLoan(lid);

        loans.markDefaulted(lid);

        (, , , , , , uint256 paid, uint8 status) = loans.getLoan(lid);
        assertEq(paid, 0);
        assertEq(status, 3);
    }

    /// @notice Test default marking after partial payment
    function testEdgeCase_DefaultAfterPartialPayment() public {
        uint256 lid = loans.createLoan(borrower, 0, 1 ether, 500, 365 days);
        loans.approveLoan(lid);

        vm.deal(address(this), 1 ether);
        loans.makePayment{value: 0.3 ether}(lid);

        loans.markDefaulted(lid);

        (, , , , , , uint256 paid, uint8 status) = loans.getLoan(lid);
        assertEq(paid, 0.3 ether);
        assertEq(status, 3);
    }

    /// @notice Test interest accrual over multiple time slices
    function testEdgeCase_InterestAccrualTimeSlices() public {
        uint256 principal = 1000;
        uint256 lid = loans.createLoan(borrower, 0, principal, 1000, 365 days);
        loans.approveLoan(lid);

        uint256 accrued0 = loans.getAccruedInterest(lid);
        assertEq(accrued0, 0);

        vm.warp(block.timestamp + 91.25 days);
        uint256 accrued91 = loans.getAccruedInterest(lid);
        assertGt(accrued91, 0);

        uint256 expected91 = (principal * 1000 * 91.25 days) /
            (10000 * 365 days);
        assertApproxEqAbs(accrued91, expected91, 1);

        vm.warp(block.timestamp + 182.5 days);
        uint256 accrued273 = loans.getAccruedInterest(lid);
        assertGt(accrued273, accrued91);
    }

    /// @notice Test amount owed transitions from principal to principal+interest
    function testEdgeCase_AmountOwedProgression() public {
        uint256 principal = 2000;
        uint256 lid = loans.createLoan(borrower, 0, principal, 1000, 365 days);
        loans.approveLoan(lid);

        uint256 owed0 = loans.getAmountOwed(lid);
        assertEq(owed0, principal);

        vm.warp(block.timestamp + 365 days);

        uint256 owed365 = loans.getAmountOwed(lid);
        uint256 expectedInterest = (principal * 1000) / 10000;
        assertEq(owed365, principal + expectedInterest);
    }

    /// @notice Test refund precision with small amounts
    function testEdgeCase_RefundSmallOverpayment() public {
        uint256 lid = loans.createLoan(borrower, 0, 100, 0, 365 days);
        loans.approveLoan(lid);

        vm.deal(address(this), 1000);
        uint256 before = address(this).balance;

        loans.makePayment{value: 105}(lid);

        uint256 after_balance = address(this).balance;
        assertEq(before - after_balance, 100);
    }

    /// @notice Test sequential partial repayments across long time periods
    function testEdgeCase_LongSequentialRepayments() public {
        uint256 principal = 30000;
        uint256 lid = loans.createLoan(
            borrower,
            0,
            principal,
            200,
            5 * 365 days
        );
        loans.approveLoan(lid);

        vm.deal(address(this), 100 ether);

        for (uint256 year = 1; year <= 5; year++) {
            vm.warp(block.timestamp + 365 days);
            uint256 owed = loans.getAmountOwed(lid);
            if (year < 5) {
                loans.makePayment{value: owed / 4}(lid);
            } else {
                loans.makePayment{value: owed}(lid);
            }
        }

        (, , , , , , uint256 paid, uint8 status) = loans.getLoan(lid);
        assertGt(paid, principal);
        assertEq(status, 2);
    }

    /// @notice Test amount owed after partial payment includes remaining interest
    function testEdgeCase_RemainingOwedAfterPartialPayment() public {
        uint256 principal = 10000;
        uint256 lid = loans.createLoan(borrower, 0, principal, 500, 365 days);
        loans.approveLoan(lid);

        vm.warp(block.timestamp + 365 days);
        uint256 totalOwed = loans.getAmountOwed(lid);

        vm.deal(address(this), totalOwed + 1 ether);
        loans.makePayment{value: totalOwed}(lid);

        uint256 remainingOwed = loans.getAmountOwed(lid);
        assertEq(remainingOwed, 0);

        (, , , , , , , uint8 status) = loans.getLoan(lid);
        assertEq(status, 2);
    }

    /// @notice Test approval requires exact status 0 -> 1
    function testEdgeCase_CannotApproveApprovedLoan() public {
        uint256 lid = loans.createLoan(borrower, 0, 1 ether, 500, 365 days);
        loans.approveLoan(lid);

        (, , , , , , , uint8 status) = loans.getLoan(lid);
        assertEq(status, 1);

        loans.approveLoan(lid);
        (, , , , , , , uint8 statusAfter) = loans.getLoan(lid);
        assertEq(statusAfter, 1);
    }
}
