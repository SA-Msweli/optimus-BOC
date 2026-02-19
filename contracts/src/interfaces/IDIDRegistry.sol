// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

/// @notice Minimal on‑chain DID registry interface (Privy pointer + risk score)
interface IDIDRegistry {
    /* Events (indexed fields for CRE triggers / fast queries) */
    event DIDCreated(address indexed owner, uint256 createdAt);
    event PrivyCredentialLinked(
        address indexed owner,
        bytes32 indexed credentialHash,
        uint256 linkedAt
    );
    event RiskProfileUpdated(
        address indexed owner,
        uint256 newScore,
        bytes32 indexed profileHash
    );

    /* Lifecycle */

    function createDID(address owner) external;

    /* Privy linkage (store hashed pointer only) */
    function linkPrivyCredential(
        address owner,
        bytes32 privyCredentialHash
    ) external;

    function getPrivyCredentialHash(
        address owner
    ) external view returns (bytes32);

    /* Risk scoring */
    function updateRiskProfile(
        address owner,
        uint256 newScore,
        bytes32 riskProfileHash
    ) external;

    function getRiskProfileScore(address owner) external view returns (uint256);
}
