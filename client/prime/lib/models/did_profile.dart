/// Local aggregate of DID identity data built from multiple API calls.
///
/// The backend does not return all these fields in one call:
/// - GET /did/{owner}       → { "exists": true }
/// - GET /did/{owner}/risk  → { "risk_score": "720" }
/// - GET /did/{owner}/privy → { "hash": "0x…" }
class DIDProfile {
  final String owner;
  final bool exists;
  final int riskScore;
  final String? privyHash;

  const DIDProfile({
    required this.owner,
    this.exists = false,
    this.riskScore = 0,
    this.privyHash,
  });

  String get riskTier {
    if (riskScore >= 800) return 'EXCELLENT';
    if (riskScore >= 700) return 'GOOD';
    if (riskScore >= 600) return 'FAIR';
    if (riskScore >= 500) return 'POOR';
    return 'VERY_POOR';
  }

  DIDProfile copyWith({
    String? owner,
    bool? exists,
    int? riskScore,
    String? privyHash,
  }) => DIDProfile(
    owner: owner ?? this.owner,
    exists: exists ?? this.exists,
    riskScore: riskScore ?? this.riskScore,
    privyHash: privyHash ?? this.privyHash,
  );
}
