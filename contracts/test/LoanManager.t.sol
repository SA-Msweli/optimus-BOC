// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {Test} from "forge-std/Test.sol";
import {LoanManager} from "../src/LoanManager.sol";

contract LoanManagerTest is Test {
    LoanManager loans;
    address borrower = address(0x1);

    function setUp() public {
        loans = new LoanManager();
    }

    receive() external payable {}

    function testCreateLoan_StoresCorrectValues() public {
        uint256 principal = 1 ether;
        uint256 interestBps = 500;
        uint256 duration = 365 days;

        uint256 lid = loans.createLoan(
            borrower,
            0,
            principal,
            interestBps,
            duration
        );

        (
            uint256 id,
            address b,
            uint256 p,
            uint256 ibps,
            uint256 start,
            uint256 end,
            uint256 paid,
            uint8 status
        ) = loans.getLoan(lid);

        assertEq(id, lid);
        assertEq(b, borrower);
        assertEq(p, principal);
        assertEq(ibps, interestBps);
        assertEq(start, block.timestamp);
        assertEq(end, block.timestamp + duration);
        assertEq(paid, 0);
        assertEq(status, 0);
    }

    function testCreateLoan_RevertsZeroPrincipal() public {
        vm.expectRevert("ZERO_PRINCIPAL");
        loans.createLoan(borrower, 0, 0, 500, 365 days);
    }

    function testCreateLoan_RevertsZeroDuration() public {
        vm.expectRevert("ZERO_DURATION");
        loans.createLoan(borrower, 0, 1 ether, 500, 0);
    }

    function testApproveLoan_TransitionsStatus() public {
        uint256 lid = loans.createLoan(borrower, 0, 1 ether, 500, 365 days);
        loans.approveLoan(lid);

        (, , , , , , , uint8 status) = loans.getLoan(lid);
        assertEq(status, 1);
    }

    function testAccruedInterest_CalculatesCorrectly() public {
        uint256 principal = 10000;
        uint256 interestBps = 500;

        uint256 lid = loans.createLoan(
            borrower,
            0,
            principal,
            interestBps,
            365 days
        );

        uint256 accrued0 = loans.getAccruedInterest(lid);
        assertEq(accrued0, 0);

        vm.warp(block.timestamp + 182.5 days);
        uint256 accrued180 = loans.getAccruedInterest(lid);
        assertGt(accrued180, 0);

        uint256 expected = (principal * interestBps * 182.5 days) /
            (10000 * 365 days);
        assertApproxEqAbs(accrued180, expected, 1);
    }

    function testAmountOwed_IncludesPrincipalAndInterest() public {
        uint256 principal = 10000;
        uint256 interestBps = 1000;

        uint256 lid = loans.createLoan(
            borrower,
            0,
            principal,
            interestBps,
            365 days
        );
        loans.approveLoan(lid);

        uint256 owed0 = loans.getAmountOwed(lid);
        assertEq(owed0, principal);

        vm.warp(block.timestamp + 365 days);
        uint256 owedFull = loans.getAmountOwed(lid);
        uint256 expectedInterest = (principal * interestBps) / 10000;
        assertEq(owedFull, principal + expectedInterest);
    }

    function testMakePayment_RevertsIfNotApproved() public {
        uint256 lid = loans.createLoan(borrower, 0, 1 ether, 500, 365 days);
        vm.deal(address(this), 1 ether);

        vm.expectRevert("LOAN_NOT_APPROVED");
        loans.makePayment{value: 1 ether}(lid);
    }

    function testMakePayment_AcceptsExactAmount() public {
        uint256 lid = loans.createLoan(borrower, 0, 1 ether, 0, 365 days);
        loans.approveLoan(lid);

        vm.deal(address(this), 1 ether);
        loans.makePayment{value: 1 ether}(lid);

        (, , , , , , uint256 paid, uint8 status) = loans.getLoan(lid);
        assertEq(paid, 1 ether);
        assertEq(status, 2);
    }

    function testMakePayment_RefundsExcessPayment() public {
        uint256 lid = loans.createLoan(borrower, 0, 1 ether, 0, 365 days);
        loans.approveLoan(lid);

        vm.deal(address(this), 2 ether);
        uint256 before = address(this).balance;

        loans.makePayment{value: 1.5 ether}(lid);

        uint256 after_balance = address(this).balance;
        assertEq(before - after_balance, 1 ether);
    }

    function testMakePayment_IncludesAccruedInterest() public {
        uint256 principal = 10000;
        uint256 interestBps = 1000;

        uint256 lid = loans.createLoan(
            borrower,
            0,
            principal,
            interestBps,
            365 days
        );
        loans.approveLoan(lid);

        vm.warp(block.timestamp + 365 days);

        uint256 owedFull = loans.getAmountOwed(lid);
        vm.deal(address(this), owedFull + 1 ether);

        loans.makePayment{value: owedFull}(lid);

        (, , , , , , uint256 paid, uint8 status) = loans.getLoan(lid);
        assertEq(paid, owedFull);
        assertEq(status, 2);
    }

    function testMakePayment_PartialPayments() public {
        uint256 lid = loans.createLoan(borrower, 0, 10 ether, 0, 365 days);
        loans.approveLoan(lid);

        vm.deal(address(this), 10 ether);

        loans.makePayment{value: 3 ether}(lid);
        (, , , , , , uint256 paid1, uint8 status1) = loans.getLoan(lid);
        assertEq(paid1, 3 ether);
        assertEq(status1, 1);

        loans.makePayment{value: 7 ether}(lid);
        (, , , , , , uint256 paid2, uint8 status2) = loans.getLoan(lid);
        assertEq(paid2, 10 ether);
        assertEq(status2, 2);
    }

    function testMarkDefaulted_TransitionsStatus() public {
        uint256 lid = loans.createLoan(borrower, 0, 1 ether, 500, 365 days);
        loans.markDefaulted(lid);

        (, , , , , , , uint8 status) = loans.getLoan(lid);
        assertEq(status, 3);
    }

    function testMarkDefaulted_RevertsIfNotFound() public {
        vm.expectRevert("LOAN_NOT_FOUND");
        loans.markDefaulted(9999);
    }

    function testLoan_FullRepaymenetWithInterest() public {
        uint256 lid = loans.createLoan(borrower, 0, 1 ether, 500, 365 days);
        loans.approveLoan(lid);
        uint256 owed = loans.getAmountOwed(lid);

        vm.deal(address(this), owed);
        loans.makePayment{value: owed}(lid);

        (, , , , , , uint256 paid, uint8 status) = loans.getLoan(lid);
        assertEq(paid, owed);
        assertEq(status, 2);
    }
}
