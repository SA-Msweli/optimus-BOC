/// Mirrors the Go `models.Arrangement` struct.
class Arrangement {
  final String arrangementId;
  final String daoId;
  final String payer;
  final String recipient;
  final String totalAmount;
  final int numInstallments;
  final int startTimestamp;
  final int intervalSeconds;
  final String status;

  const Arrangement({
    required this.arrangementId,
    required this.daoId,
    required this.payer,
    required this.recipient,
    required this.totalAmount,
    required this.numInstallments,
    required this.startTimestamp,
    required this.intervalSeconds,
    required this.status,
  });

  /// Human-readable total in ETH.
  String get totalEth {
    final wei = BigInt.tryParse(totalAmount) ?? BigInt.zero;
    final eth = wei / BigInt.from(10).pow(18);
    final remainder = wei % BigInt.from(10).pow(18);
    return '${eth}.${remainder.toString().padLeft(18, '0').substring(0, 4)}';
  }

  /// Per-installment amount (total / numInstallments).
  String get installmentAmount {
    if (numInstallments <= 0) return totalAmount;
    final wei = BigInt.tryParse(totalAmount) ?? BigInt.zero;
    final part = wei ~/ BigInt.from(numInstallments);
    return part.toString();
  }

  factory Arrangement.fromJson(Map<String, dynamic> json) => Arrangement(
        arrangementId: json['arrangement_id']?.toString() ??
            json['ArrangementID']?.toString() ??
            json['ID']?.toString() ??
            '',
        daoId: json['dao_id']?.toString() ??
            json['DaoID']?.toString() ??
            json['DaoId']?.toString() ??
            '',
        payer: json['payer']?.toString() ??
            json['Payer']?.toString() ??
            '',
        recipient: json['recipient']?.toString() ??
            json['Recipient']?.toString() ??
            '',
        totalAmount: json['total_amount']?.toString() ??
            json['TotalAmount']?.toString() ??
            '0',
        numInstallments: _parseInt(
            json['num_installments'] ?? json['NumInstallments']),
        startTimestamp: _parseInt(
            json['start_timestamp'] ?? json['StartTimestamp']),
        intervalSeconds: _parseInt(
            json['interval_seconds'] ?? json['IntervalSeconds']),
        status: json['status']?.toString() ??
            json['Status']?.toString() ??
            'UNKNOWN',
      );
}

int _parseInt(dynamic v) {
  if (v is int) return v;
  if (v is String) return int.tryParse(v) ?? 0;
  return 0;
}
