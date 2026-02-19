// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

/// @notice BNPL arrangement manager interface.
interface IBNPLManager {
    event BNPLCreated(
        uint256 indexed arrangementId,
        uint256 indexed daoId,
        address indexed payer,
        address recipient,
        uint256 totalAmount,
        uint256 numInstallments,
        uint256 createdAt
    );
    event BNPLScheduleChosen(
        uint256 indexed arrangementId,
        uint256 startTimestamp,
        uint256 intervalSeconds
    );
    event BNPLPaymentMade(
        uint256 indexed arrangementId,
        uint8 installmentNumber,
        address indexed payer,
        uint256 amount,
        uint256 timestamp
    );
    event BNPLActivated(uint256 indexed arrangementId, uint256 activatedAt);
    event BNPLLateFeeApplied(
        uint256 indexed arrangementId,
        uint8 installmentNumber,
        uint256 feeAmount,
        uint256 timestamp
    );
    event BNPLRescheduled(
        uint256 indexed arrangementId,
        bytes32 oldScheduleHash,
        bytes32 newScheduleHash,
        uint256 timestamp
    );
    event BNPLCompleted(uint256 indexed arrangementId, uint256 completedAt);

    function createBNPL(
        uint256 daoId,
        address recipient,
        uint256 totalAmount,
        uint256 startTimestamp,
        uint256 intervalSeconds,
        bytes calldata metadata
    ) external returns (uint256 arrangementId);

    function makePayment(
        uint256 arrangementId,
        uint8 installmentNumber
    ) external payable;

    function activateBNPL(uint256 arrangementId) external;

    function applyLateFee(
        uint256 arrangementId,
        uint8 installmentNumber
    ) external;

    function reschedule(
        uint256 arrangementId,
        uint256 newStartTimestamp,
        uint256 newIntervalSeconds
    ) external;

    function getArrangement(
        uint256 arrangementId
    )
        external
        view
        returns (
            uint256 id,
            uint256 daoId,
            address payer,
            address recipient,
            uint256 totalAmount,
            uint256 numInstallments,
            uint256[] memory installmentAmounts,
            uint256 startTimestamp,
            uint256 intervalSeconds,
            uint256 lateFeeBps,
            uint8 status
        );
}
