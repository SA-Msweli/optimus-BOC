import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import '../services/auth_service.dart';
import '../services/loan_service.dart';
import '../theme.dart';
import '../widgets/shared.dart';

/// Screen for loan management: create, lookup, approve, pay.
///
/// Removed: DAO ID field (contract ignores the second param),
///          Mark Defaulted button (CRE loan_default_monitor only).
class LoanScreen extends StatefulWidget {
  const LoanScreen({super.key});

  @override
  State<LoanScreen> createState() => _LoanScreenState();
}

class _LoanScreenState extends State<LoanScreen> {
  // ── Lookup ──
  final _idCtrl = TextEditingController();

  // ── Create ──
  final _borrowerCtrl = TextEditingController();
  final _principalCtrl = TextEditingController();
  final _rateBpsCtrl = TextEditingController();
  final _durationCtrl = TextEditingController();

  // ── Payment ──
  final _payAmountCtrl = TextEditingController();

  bool _showCreate = false;
  bool _prefilled = false;

  @override
  void dispose() {
    _idCtrl.dispose();
    _borrowerCtrl.dispose();
    _principalCtrl.dispose();
    _rateBpsCtrl.dispose();
    _durationCtrl.dispose();
    _payAmountCtrl.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    // Auto-fill borrower from authenticated Privy wallet (spec: caller = borrower).
    if (!_prefilled) {
      final auth = context.read<AuthService>();
      if (auth.walletAddress != null && _borrowerCtrl.text.isEmpty) {
        _borrowerCtrl.text = auth.walletAddress!;
        _prefilled = true;
      }
    }

    return Consumer<LoanService>(
      builder: (context, svc, _) {
        return ListView(
          padding: const EdgeInsets.all(16),
          children: [
            if (svc.error != null)
              ErrorBanner(
                message: svc.error!,
                onDismiss: () => svc.clearError(),
              ),
            if (svc.lastTx != null)
              SuccessBanner(message: 'TX: ${truncateAddress(svc.lastTx!)}'),

            // ── Lookup ──
            InfoCard(
              title: 'Lookup Loan',
              child: Row(
                children: [
                  Expanded(
                    child: TextField(
                      controller: _idCtrl,
                      decoration: AppTheme.inputDecoration(
                        'Loan ID',
                        hint: '1',
                      ),
                      keyboardType: TextInputType.number,
                    ),
                  ),
                  const SizedBox(width: 8),
                  ElevatedButton(
                    onPressed: svc.loading
                        ? null
                        : () {
                            final id = _idCtrl.text.trim();
                            svc.fetchLoan(id);
                            svc.fetchAccruedInterest(id);
                            svc.fetchAmountOwed(id);
                          },
                    child: const Text('Fetch'),
                  ),
                ],
              ),
            ),

            // ── Loan details ──
            if (svc.current != null) ...[
              InfoCard(
                title: 'Loan #${svc.current!.loanId}',
                child: Column(
                  children: [
                    KVRow(
                      label: 'Borrower',
                      value: truncateAddress(svc.current!.borrower),
                      mono: true,
                    ),
                    KVRow(
                      label: 'Principal (wei)',
                      value: svc.current!.principal,
                    ),
                    KVRow(
                      label: 'Principal (ETH)',
                      value: svc.current!.principalEth,
                    ),
                    KVRow(
                      label: 'Interest Rate',
                      value: svc.current!.interestRatePercent,
                    ),
                    KVRow(
                      label: 'Start',
                      value: _formatTimestamp(svc.current!.startTime),
                    ),
                    KVRow(
                      label: 'End',
                      value: _formatTimestamp(svc.current!.endTime),
                    ),
                    KVRow(
                      label: 'Amount Paid',
                      value: svc.current!.amountPaidEth,
                    ),
                    KVRow(
                      label: 'Status',
                      value: _loanStatus(svc.current!.status),
                    ),
                    const Divider(),
                    KVRow(
                      label: 'Accrued Interest',
                      value: svc.accruedInterest,
                    ),
                    KVRow(label: 'Total Owed', value: svc.amountOwed),
                  ],
                ),
              ),

              // ── Make Payment ──
              InfoCard(
                title: 'Make Payment',
                child: Column(
                  children: [
                    TextField(
                      controller: _payAmountCtrl,
                      decoration: AppTheme.inputDecoration(
                        'Payment Amount (wei)',
                        hint: svc.amountOwed == '0'
                            ? '1000000000000000000'
                            : svc.amountOwed,
                      ),
                      keyboardType: TextInputType.number,
                    ),
                    const SizedBox(height: 8),
                    SizedBox(
                      width: double.infinity,
                      child: ElevatedButton.icon(
                        onPressed: svc.loading
                            ? null
                            : () => svc.makePayment(
                                svc.current!.loanId,
                                _payAmountCtrl.text.trim(),
                              ),
                        icon: const Icon(Icons.payment, size: 18),
                        label: const Text('Pay'),
                      ),
                    ),
                  ],
                ),
              ),

              // ── Other actions ──
              Wrap(
                spacing: 8,
                runSpacing: 8,
                children: [
                  ElevatedButton.icon(
                    onPressed: svc.loading
                        ? null
                        : () => svc.approveLoan(svc.current!.loanId),
                    icon: const Icon(Icons.check_circle, size: 18),
                    label: const Text('Approve'),
                  ),
                  OutlinedButton.icon(
                    onPressed: svc.loading
                        ? null
                        : () {
                            svc.fetchAccruedInterest(svc.current!.loanId);
                            svc.fetchAmountOwed(svc.current!.loanId);
                          },
                    icon: const Icon(Icons.refresh, size: 18),
                    label: const Text('Refresh Amounts'),
                  ),
                ],
              ),
            ],

            const Divider(height: 32),

            // ── Toggle create ──
            TextButton.icon(
              onPressed: () => setState(() => _showCreate = !_showCreate),
              icon: Icon(_showCreate ? Icons.close : Icons.add),
              label: Text(_showCreate ? 'Cancel' : 'New Loan'),
            ),

            if (_showCreate) _buildCreateForm(svc),

            if (svc.loading) const LoadingOverlay(label: 'Processing…'),
          ],
        );
      },
    );
  }

  Widget _buildCreateForm(LoanService svc) {
    return InfoCard(
      title: 'Create Loan',
      child: Column(
        children: [
          TextField(
            controller: _borrowerCtrl,
            decoration: AppTheme.inputDecoration('Borrower (0x…)', hint: '0x…'),
          ),
          const SizedBox(height: 8),
          TextField(
            controller: _principalCtrl,
            decoration: AppTheme.inputDecoration(
              'Principal (wei)',
              hint: '1000000000000000000',
            ),
            keyboardType: TextInputType.number,
          ),
          const SizedBox(height: 8),
          TextField(
            controller: _rateBpsCtrl,
            decoration: AppTheme.inputDecoration(
              'Interest Rate (bps)',
              hint: '500 = 5.00%',
            ),
            keyboardType: TextInputType.number,
          ),
          const SizedBox(height: 8),
          TextField(
            controller: _durationCtrl,
            decoration: AppTheme.inputDecoration(
              'Duration (seconds)',
              hint: '2592000 = 30 days',
            ),
            keyboardType: TextInputType.number,
          ),
          const SizedBox(height: 12),
          SizedBox(
            width: double.infinity,
            child: ElevatedButton.icon(
              onPressed: svc.loading
                  ? null
                  : () => svc.createLoan(
                      borrower: _borrowerCtrl.text.trim(),
                      principal: _principalCtrl.text.trim(),
                      interestRateBps: _rateBpsCtrl.text.trim(),
                      durationSeconds: _durationCtrl.text.trim(),
                    ),
              icon: const Icon(Icons.add_card),
              label: const Text('Create Loan'),
            ),
          ),
        ],
      ),
    );
  }

  String _formatTimestamp(int ts) {
    if (ts == 0) return 'N/A';
    final dt = DateTime.fromMillisecondsSinceEpoch(ts * 1000);
    return '${dt.year}-${_pad(dt.month)}-${_pad(dt.day)} '
        '${_pad(dt.hour)}:${_pad(dt.minute)}';
  }

  String _pad(int v) => v.toString().padLeft(2, '0');

  String _loanStatus(String raw) {
    final code = int.tryParse(raw);
    if (code == null) return raw;
    switch (code) {
      case 0:
        return 'PENDING';
      case 1:
        return 'APPROVED';
      case 2:
        return 'REPAID';
      case 3:
        return 'DEFAULTED';
      default:
        return 'UNKNOWN ($code)';
    }
  }
}
