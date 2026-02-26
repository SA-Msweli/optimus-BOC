/// Mirrors the Go `models.DAO` struct.
class DAO {
  final String daoId;
  final String creator;
  final String goal;
  final int votingPeriodSeconds;
  final String treasuryBalance;
  final String totalInvestments;

  const DAO({
    required this.daoId,
    required this.creator,
    required this.goal,
    required this.votingPeriodSeconds,
    required this.treasuryBalance,
    required this.totalInvestments,
  });

  String get goalLabel {
    switch (goal) {
      case '0':
        return 'General';
      case '1':
        return 'Lending';
      case '2':
        return 'BNPL';
      default:
        return 'Type $goal';
    }
  }

  String get votingPeriodDays =>
      '${(votingPeriodSeconds / 86400).round()} days';

  factory DAO.fromJson(Map<String, dynamic> json) => DAO(
        daoId: json['dao_id']?.toString() ??
            json['DaoID']?.toString() ??
            '',
        creator: json['creator']?.toString() ??
            json['Creator']?.toString() ??
            '',
        goal: json['goal']?.toString() ??
            json['Goal']?.toString() ??
            '0',
        votingPeriodSeconds: _parseInt(
            json['voting_period_seconds'] ?? json['VotingPeriodSeconds']),
        treasuryBalance: json['treasury_balance']?.toString() ??
            json['TreasuryBalance']?.toString() ??
            '0',
        totalInvestments: json['total_investments']?.toString() ??
            json['TotalInvestments']?.toString() ??
            '0',
      );
}

int _parseInt(dynamic v) {
  if (v is int) return v;
  if (v is String) return int.tryParse(v) ?? 0;
  return 0;
}
