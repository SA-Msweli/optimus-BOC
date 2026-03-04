import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import '../services/vault_service.dart';
import '../theme.dart';
import '../widgets/shared.dart';

/// Screen for token vault operations: deposit, withdraw, check balance.
///
/// The token field should be an ERC-20 token address (NOT the vault contract
/// address). Users deposit/withdraw a specific ERC-20 token into the vault.
class VaultScreen extends StatefulWidget {
  const VaultScreen({super.key});

  @override
  State<VaultScreen> createState() => _VaultScreenState();
}

class _VaultScreenState extends State<VaultScreen> {
  final _tokenCtrl = TextEditingController();
  final _amountCtrl = TextEditingController();

  @override
  void dispose() {
    _tokenCtrl.dispose();
    _amountCtrl.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Consumer<VaultService>(
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

            // ── Token address ──
            InfoCard(
              title: 'ERC-20 Token Address',
              child: TextField(
                controller: _tokenCtrl,
                decoration: AppTheme.inputDecoration(
                  'ERC-20 Token (0x…)',
                  hint: 'Enter the token contract address to deposit/withdraw',
                ),
              ),
            ),

            // ── Balance ──
            InfoCard(
              title: 'Vault Balance',
              child: Column(
                children: [
                  KVRow(label: 'Balance (wei)', value: svc.balance),
                  const SizedBox(height: 8),
                  SizedBox(
                    width: double.infinity,
                    child: ElevatedButton.icon(
                      onPressed: svc.loading
                          ? null
                          : () => svc.fetchBalance(_tokenCtrl.text.trim()),
                      icon: const Icon(Icons.refresh, size: 18),
                      label: const Text('Refresh Balance'),
                    ),
                  ),
                ],
              ),
            ),

            const Divider(height: 24),

            // ── Amount input ──
            InfoCard(
              title: 'Transaction Amount',
              child: TextField(
                controller: _amountCtrl,
                decoration: AppTheme.inputDecoration(
                  'Amount (wei)',
                  hint: '1000000000000000000',
                ),
                keyboardType: TextInputType.number,
              ),
            ),

            // ── Deposit / Withdraw ──
            Row(
              children: [
                Expanded(
                  child: ElevatedButton.icon(
                    onPressed: svc.loading
                        ? null
                        : () => svc.deposit(
                            _tokenCtrl.text.trim(),
                            _amountCtrl.text.trim(),
                          ),
                    icon: const Icon(Icons.arrow_downward, size: 18),
                    label: const Text('Deposit'),
                    style: ElevatedButton.styleFrom(
                      backgroundColor: AppTheme.secondary,
                    ),
                  ),
                ),
                const SizedBox(width: 12),
                Expanded(
                  child: ElevatedButton.icon(
                    onPressed: svc.loading
                        ? null
                        : () => svc.withdraw(
                            _tokenCtrl.text.trim(),
                            _amountCtrl.text.trim(),
                          ),
                    icon: const Icon(Icons.arrow_upward, size: 18),
                    label: const Text('Withdraw'),
                    style: ElevatedButton.styleFrom(
                      backgroundColor: AppTheme.warning,
                    ),
                  ),
                ),
              ],
            ),

            if (svc.loading)
              const Padding(
                padding: EdgeInsets.only(top: 24),
                child: LoadingOverlay(label: 'Processing…'),
              ),
          ],
        );
      },
    );
  }
}
