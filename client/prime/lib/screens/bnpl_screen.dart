import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import '../services/auth_service.dart';
import '../services/bnpl_service.dart';
import '../theme.dart';
import '../widgets/shared.dart';

/// Screen for BNPL arrangement management: create, lookup, pay, reschedule.
///
/// Removed: Activate (redundant – auto-activates on first payment),
///          Late Fee (CRE-only via bnpl_late_fee cron workflow).
class BNPLScreen extends StatefulWidget {
  const BNPLScreen({super.key});

  @override
  State<BNPLScreen> createState() => _BNPLScreenState();
}

class _BNPLScreenState extends State<BNPLScreen> {
  // ── Lookup ──
  final _idCtrl = TextEditingController();

  // ── Create ──
  final _daoIdCtrl = TextEditingController();
  final _recipientCtrl = TextEditingController();
  final _totalCtrl = TextEditingController();
  final _startCtrl = TextEditingController();
  final _intervalCtrl = TextEditingController();

  // ── Payment ──
  final _installmentCtrl = TextEditingController();
  final _payAmountCtrl = TextEditingController();

  // ── Reschedule ──
  final _newStartCtrl = TextEditingController();
  final _newIntervalCtrl = TextEditingController();

  bool _showCreate = false;
  bool _prefilled = false;

  @override
  void dispose() {
    _idCtrl.dispose();
    _daoIdCtrl.dispose();
    _recipientCtrl.dispose();
    _totalCtrl.dispose();
    _startCtrl.dispose();
    _intervalCtrl.dispose();
    _installmentCtrl.dispose();
    _payAmountCtrl.dispose();
    _newStartCtrl.dispose();
    _newIntervalCtrl.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    // Auto-fill the recipient from the authenticated Privy wallet.
    if (!_prefilled) {
      final auth = context.read<AuthService>();
      if (auth.walletAddress != null && _recipientCtrl.text.isEmpty) {
        _recipientCtrl.text = auth.walletAddress!;
        _prefilled = true;
      }
    }

    return Consumer<BNPLService>(
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
              title: 'Lookup Arrangement',
              child: Row(
                children: [
                  Expanded(
                    child: TextField(
                      controller: _idCtrl,
                      decoration: AppTheme.inputDecoration(
                        'Arrangement ID',
                        hint: '1',
                      ),
                      keyboardType: TextInputType.number,
                    ),
                  ),
                  const SizedBox(width: 8),
                  ElevatedButton(
                    onPressed: svc.loading
                        ? null
                        : () => svc.fetchArrangement(_idCtrl.text.trim()),
                    child: const Text('Fetch'),
                  ),
                ],
              ),
            ),

            // ── Arrangement details ──
            if (svc.current != null) ...[
              InfoCard(
                title: 'Arrangement #${svc.current!.arrangementId}',
                child: Column(
                  children: [
                    KVRow(label: 'DAO ID', value: svc.current!.daoId),
                    KVRow(
                      label: 'Payer',
                      value: truncateAddress(svc.current!.payer),
                      mono: true,
                    ),
                    KVRow(
                      label: 'Recipient',
                      value: truncateAddress(svc.current!.recipient),
                      mono: true,
                    ),
                    KVRow(
                      label: 'Total (wei)',
                      value: svc.current!.totalAmount,
                    ),
                    KVRow(label: 'Total (ETH)', value: svc.current!.totalEth),
                    KVRow(
                      label: 'Installments',
                      value: svc.current!.numInstallments.toString(),
                    ),
                    KVRow(
                      label: 'Per Installment',
                      value: svc.current!.installmentAmount,
                    ),
                    KVRow(
                      label: 'Start',
                      value: _formatTimestamp(svc.current!.startTimestamp),
                    ),
                    KVRow(
                      label: 'Interval',
                      value: '${svc.current!.intervalSeconds}s',
                    ),
                    KVRow(
                      label: 'Status',
                      value: _statusLabel(svc.current!.status),
                    ),
                  ],
                ),
              ),

              // ── Make Payment ──
              InfoCard(
                title: 'Make Payment',
                child: Column(
                  children: [
                    TextField(
                      controller: _installmentCtrl,
                      decoration: AppTheme.inputDecoration(
                        'Installment #',
                        hint: '0',
                      ),
                      keyboardType: TextInputType.number,
                    ),
                    const SizedBox(height: 8),
                    TextField(
                      controller: _payAmountCtrl,
                      decoration: AppTheme.inputDecoration(
                        'Payment Amount (wei)',
                        hint: svc.current!.installmentAmount,
                      ),
                      keyboardType: TextInputType.number,
                    ),
                    const SizedBox(height: 8),
                    SizedBox(
                      width: double.infinity,
                      child: ElevatedButton.icon(
                        onPressed: svc.loading
                            ? null
                            : () {
                                final inst =
                                    int.tryParse(_installmentCtrl.text) ?? 0;
                                svc.makePayment(
                                  svc.current!.arrangementId,
                                  inst,
                                  _payAmountCtrl.text.trim(),
                                );
                              },
                        icon: const Icon(Icons.payment, size: 18),
                        label: const Text('Pay Installment'),
                      ),
                    ),
                  ],
                ),
              ),

              // ── Reschedule ──
              InfoCard(
                title: 'Reschedule',
                child: Column(
                  children: [
                    TextField(
                      controller: _newStartCtrl,
                      decoration: AppTheme.inputDecoration(
                        'New Start (unix)',
                        hint: '1700000000',
                      ),
                      keyboardType: TextInputType.number,
                    ),
                    const SizedBox(height: 8),
                    TextField(
                      controller: _newIntervalCtrl,
                      decoration: AppTheme.inputDecoration(
                        'New Interval (sec)',
                        hint: '86400',
                      ),
                      keyboardType: TextInputType.number,
                    ),
                    const SizedBox(height: 8),
                    SizedBox(
                      width: double.infinity,
                      child: ElevatedButton.icon(
                        onPressed: svc.loading
                            ? null
                            : () => svc.reschedule(
                                svc.current!.arrangementId,
                                int.tryParse(_newStartCtrl.text) ?? 0,
                                int.tryParse(_newIntervalCtrl.text) ?? 0,
                              ),
                        icon: const Icon(Icons.schedule, size: 18),
                        label: const Text('Reschedule'),
                      ),
                    ),
                  ],
                ),
              ),
            ],

            const Divider(height: 32),

            // ── Create toggle ──
            TextButton.icon(
              onPressed: () => setState(() => _showCreate = !_showCreate),
              icon: Icon(_showCreate ? Icons.close : Icons.add),
              label: Text(_showCreate ? 'Cancel' : 'New Arrangement'),
            ),

            if (_showCreate) _buildCreateForm(svc),

            if (svc.loading) const LoadingOverlay(label: 'Processing…'),
          ],
        );
      },
    );
  }

  Widget _buildCreateForm(BNPLService svc) {
    return InfoCard(
      title: 'Create BNPL Arrangement',
      child: Column(
        children: [
          TextField(
            controller: _daoIdCtrl,
            decoration: AppTheme.inputDecoration(
              'DAO ID',
              hint: 'Sponsoring DAO',
            ),
            keyboardType: TextInputType.number,
          ),
          const SizedBox(height: 8),
          TextField(
            controller: _recipientCtrl,
            decoration: AppTheme.inputDecoration(
              'Recipient Address',
              hint: '0x…',
            ),
          ),
          const SizedBox(height: 8),
          TextField(
            controller: _totalCtrl,
            decoration: AppTheme.inputDecoration(
              'Total Amount (wei)',
              hint: 'e.g. 1000000000000000000 = 1 ETH',
            ),
            keyboardType: TextInputType.number,
          ),
          const SizedBox(height: 8),
          TextField(
            controller: _startCtrl,
            decoration: AppTheme.inputDecoration(
              'Start Date (unix timestamp)',
              hint: '${DateTime.now().millisecondsSinceEpoch ~/ 1000}',
            ),
            keyboardType: TextInputType.number,
          ),
          const SizedBox(height: 8),
          TextField(
            controller: _intervalCtrl,
            decoration: AppTheme.inputDecoration(
              'Payment Interval (seconds)',
              hint: '86400 = 1 day',
            ),
            keyboardType: TextInputType.number,
          ),
          const SizedBox(height: 12),
          SizedBox(
            width: double.infinity,
            child: ElevatedButton.icon(
              onPressed: svc.loading
                  ? null
                  : () => svc.createArrangement(
                      daoId: _daoIdCtrl.text.trim(),
                      recipient: _recipientCtrl.text.trim(),
                      totalAmount: _totalCtrl.text.trim(),
                      startTimestamp: int.tryParse(_startCtrl.text) ?? 0,
                      intervalSeconds: int.tryParse(_intervalCtrl.text) ?? 0,
                    ),
              icon: const Icon(Icons.add_shopping_cart),
              label: const Text('Create'),
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

  String _statusLabel(String raw) {
    final code = int.tryParse(raw);
    if (code == null) return raw;
    switch (code) {
      case 0:
        return 'PENDING';
      case 1:
        return 'ACTIVE';
      case 2:
        return 'COMPLETED';
      default:
        return 'UNKNOWN ($code)';
    }
  }
}
