/// Mirrors the Go `models.Loan` struct.
class Loan {
  final String loanId;
  final String borrower;
  final String daoAddress;
  final String principal;
  final int interestRateBps;
  final int startTime;
  final int endTime;
  final String amountPaid;
  final String status;

  const Loan({
    required this.loanId,
    required this.borrower,
    required this.daoAddress,
    required this.principal,
    required this.interestRateBps,
    required this.startTime,
    required this.endTime,
    required this.amountPaid,
    required this.status,
  });

  /// Annual interest rate as a percentage string (e.g. "5.00%").
  String get interestRatePercent =>
      '${(interestRateBps / 100).toStringAsFixed(2)}%';

  /// Human-readable principal in ETH.
  String get principalEth => _weiToEth(principal);

  /// Human-readable amount paid in ETH.
  String get amountPaidEth => _weiToEth(amountPaid);

  factory Loan.fromJson(Map<String, dynamic> json) => Loan(
        loanId: json['loan_id']?.toString() ??
            json['LoanID']?.toString() ??
            json['loanId']?.toString() ??
            json['Id']?.toString() ??
            '',
        borrower: json['borrower']?.toString() ??
            json['Borrower']?.toString() ??
            '',
        daoAddress: json['dao_address']?.toString() ??
            json['DaoAddress']?.toString() ??
            json['DaoId']?.toString() ??
            '',
        principal: json['principal']?.toString() ??
            json['Principal']?.toString() ??
            '0',
        interestRateBps: _parseInt(
            json['interest_rate_bps'] ?? json['InterestRateBps']),
        startTime: _parseInt(json['start_time'] ?? json['StartTime']),
        endTime: _parseInt(json['end_time'] ?? json['EndTime']),
        amountPaid: json['amount_paid']?.toString() ??
            json['AmountPaid']?.toString() ??
            '0',
        status: json['status']?.toString() ??
            json['Status']?.toString() ??
            'UNKNOWN',
      );
}

String _weiToEth(String wei) {
  final w = BigInt.tryParse(wei) ?? BigInt.zero;
  final eth = w ~/ BigInt.from(10).pow(18);
  final rem = w % BigInt.from(10).pow(18);
  return '$eth.${rem.toString().padLeft(18, '0').substring(0, 4)}';
}

int _parseInt(dynamic v) {
  if (v is int) return v;
  if (v is String) return int.tryParse(v) ?? 0;
  return 0;
}
