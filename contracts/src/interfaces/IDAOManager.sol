// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

/// @notice DAO management interface for Optimus (governance, membership, BNPL policy).
/// @dev Events are indexed for efficient CRE event triggers.
interface IDAOManager {
    event DaoCreated(
        uint256 indexed daoId,
        address indexed creator,
        uint8 goal,
        uint256 createdAt
    );
    event MemberJoined(
        uint256 indexed daoId,
        address indexed member,
        uint256 investment,
        uint256 joinedAt
    );
    event ProposalOpened(
        uint256 indexed proposalId,
        uint256 indexed daoId,
        uint256 expiry,
        bytes data
    );
    event VoteCast(
        uint256 indexed proposalId,
        address indexed voter,
        bool support,
        uint256 weight
    );
    event ProposalFinalized(
        uint256 indexed proposalId,
        bool approved,
        uint256 finalizedAt
    );
    event ProposalExecuted(uint256 indexed proposalId, uint256 executedAt);
    event BnplTermsUpdated(
        uint256 indexed daoId,
        uint256 numInstallments,
        uint256 allowedIntervalMinDays,
        uint256 allowedIntervalMaxDays,
        uint256 lateFeeBps,
        uint256 gracePeriodDays,
        bool rescheduleAllowed
    );

    /// @notice Create a new DAO and return its id.
    function createDAO(
        address creator,
        uint8 goal,
        uint64 votingPeriodDays
    ) external returns (uint256 daoId);

    /// @notice Join `daoId` by supplying an investment amount.
    function joinDAO(
        uint256 daoId,
        address member,
        uint256 investmentAmount
    ) external;

    /// @notice Open a governance proposal for `daoId` with encoded `proposalData`.
    function propose(
        uint256 daoId,
        bytes calldata proposalData
    ) external returns (uint256 proposalId);

    /// @notice Cast a vote on `proposalId`.
    function vote(uint256 proposalId, bool support) external;

    /// @notice Finalize voting for `proposalId` (compute outcome and emit event).
    function finalizeProposal(uint256 proposalId) external;

    /// @notice Execute a finalized and approved proposal. Implementation details are application specific.
    function executeProposal(uint256 proposalId) external;

    /// @notice Get BNPL terms configured for `daoId`.
    function getBnplTerms(
        uint256 daoId
    )
        external
        view
        returns (
            uint256 numInstallments,
            uint256 allowedIntervalMinDays,
            uint256 allowedIntervalMaxDays,
            uint256 lateFeeBps,
            uint256 gracePeriodDays,
            bool rescheduleAllowed,
            uint256 minDownPaymentBps
        );

    /// @notice Set BNPL terms for `daoId`.
    function setBnplTerms(
        uint256 daoId,
        uint256 numInstallments,
        uint256 allowedIntervalMinDays,
        uint256 allowedIntervalMaxDays,
        uint256 lateFeeBps,
        uint256 gracePeriodDays,
        bool rescheduleAllowed,
        uint256 minDownPaymentBps
    ) external;
}
