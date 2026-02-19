// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {AccessControl} from "openzeppelin-contracts/contracts/access/AccessControl.sol";
import {IDIDRegistry} from "./interfaces/IDIDRegistry.sol";

error DIDAlreadyExists(address owner);
error DIDNotFound(address owner);
error Unauthorized();
error InvalidScore(uint256 score);

/// @title DIDRegistry
/// @notice On‑chain DID registry storing hashed Privy credential handles and a canonical risk score.
/// @dev Emits events consumed by CRE workflows. Stores only hashes (no PII). Access control roles:
///      - DEFAULT_ADMIN_ROLE: administrative actions
///      - PRIVY_LINKER_ROLE: permitted to link Privy credentials
///      - RISK_UPDATER_ROLE: permitted to update risk profiles
contract DIDRegistry is AccessControl, IDIDRegistry {
    bytes32 public constant RISK_UPDATER_ROLE = keccak256("RISK_UPDATER_ROLE");
    bytes32 public constant PRIVY_LINKER_ROLE = keccak256("PRIVY_LINKER_ROLE");

    struct Did {
        bytes32 privyCredentialHash;
        uint256 riskScore;
        bytes32 riskProfileHash;
        uint256 createdAt;
        bool exists;
    }

    mapping(address => Did) private _dids;

    /// @dev Grant DEFAULT_ADMIN_ROLE to deployer.
    constructor() {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
    }

    /// @notice Create a DID record for `owner`.
    /// @dev Callable by the `owner` or `DEFAULT_ADMIN_ROLE`.
    /// @param owner The account that will own the DID.
    function createDID(address owner) external override {
        if (_dids[owner].exists) revert DIDAlreadyExists(owner);
        if (msg.sender != owner && !hasRole(DEFAULT_ADMIN_ROLE, msg.sender))
            revert Unauthorized();

        _dids[owner] = Did({
            privyCredentialHash: bytes32(0),
            riskScore: 0,
            riskProfileHash: bytes32(0),
            createdAt: block.timestamp,
            exists: true
        });

        emit DIDCreated(owner, block.timestamp);
    }

    /// @notice Link a hashed Privy credential pointer for `owner`.
    /// @dev Callable by the `owner`, an account with `PRIVY_LINKER_ROLE`, or admin.
    /// @param owner The DID owner for whom the credential is linked.
    /// @param privyCredentialHash The keccak256 hash of the Privy credential handle.
    function linkPrivyCredential(
        address owner,
        bytes32 privyCredentialHash
    ) external override {
        if (!_dids[owner].exists) revert DIDNotFound(owner);
        if (
            msg.sender != owner &&
            !hasRole(PRIVY_LINKER_ROLE, msg.sender) &&
            !hasRole(DEFAULT_ADMIN_ROLE, msg.sender)
        ) revert Unauthorized();

        _dids[owner].privyCredentialHash = privyCredentialHash;
        emit PrivyCredentialLinked(owner, privyCredentialHash, block.timestamp);
    }

    /// @notice Returns the stored Privy credential hash for `owner`.
    /// @dev Reverts if DID does not exist.
    /// @param owner DID owner address.
    /// @return privyCredentialHash The stored credential hash.
    function getPrivyCredentialHash(
        address owner
    ) external view override returns (bytes32 privyCredentialHash) {
        if (!_dids[owner].exists) revert DIDNotFound(owner);
        return _dids[owner].privyCredentialHash;
    }

    /// @notice Update the canonical risk score for `owner`.
    /// @dev Restricted to `RISK_UPDATER_ROLE` or admin. `newScore` must be <= 10000.
    /// @param owner DID owner address.
    /// @param newScore New integer score (0..10000).
    /// @param riskProfileHash Optional hashed off‑chain profile pointer.
    function updateRiskProfile(
        address owner,
        uint256 newScore,
        bytes32 riskProfileHash
    ) external override {
        if (!_dids[owner].exists) revert DIDNotFound(owner);
        if (
            !hasRole(RISK_UPDATER_ROLE, msg.sender) &&
            !hasRole(DEFAULT_ADMIN_ROLE, msg.sender)
        ) revert Unauthorized();
        if (newScore > 10000) revert InvalidScore(newScore);

        _dids[owner].riskScore = newScore;
        _dids[owner].riskProfileHash = riskProfileHash;

        emit RiskProfileUpdated(owner, newScore, riskProfileHash);
    }

    /// @notice Returns the canonical risk score for `owner`.
    /// @param owner DID owner.
    function getRiskProfileScore(
        address owner
    ) external view override returns (uint256) {
        if (!_dids[owner].exists) revert DIDNotFound(owner);
        return _dids[owner].riskScore;
    }

    /// @notice Returns whether a DID exists for `owner`.
    /// @param owner DID owner address.
    function exists(address owner) external view returns (bool) {
        return _dids[owner].exists;
    }
}
