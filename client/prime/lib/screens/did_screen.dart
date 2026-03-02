import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import '../services/auth_service.dart';
import '../services/did_service.dart';
import '../theme.dart';
import '../widgets/shared.dart';

/// Screen for DID identity management: create, lookup, and view profile.
///
/// Risk profiles are managed exclusively by CRE workflows (RISK_UPDATER_ROLE).
/// Privy credentials are auto-linked by [AuthService] on login.
class DIDScreen extends StatefulWidget {
  const DIDScreen({super.key});

  @override
  State<DIDScreen> createState() => _DIDScreenState();
}

class _DIDScreenState extends State<DIDScreen> {
  final _ownerCtrl = TextEditingController();
  bool _prefilled = false;

  @override
  void dispose() {
    _ownerCtrl.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    // Auto-fill the wallet address from the authenticated Privy session.
    if (!_prefilled) {
      final auth = context.read<AuthService>();
      if (auth.walletAddress != null && _ownerCtrl.text.isEmpty) {
        _ownerCtrl.text = auth.walletAddress!;
        _prefilled = true;
      }
    }

    return Consumer<DIDService>(
      builder: (context, svc, _) {
        return ListView(
          padding: const EdgeInsets.all(16),
          children: [
            // ── Error / Success banners ──
            if (svc.error != null)
              ErrorBanner(
                message: svc.error!,
                onDismiss: () => svc.clearError(),
              ),
            if (svc.lastTx != null)
              SuccessBanner(message: 'TX: ${truncateAddress(svc.lastTx!)}'),

            // ── Owner address input ──
            InfoCard(
              title: 'Wallet Address',
              child: TextField(
                controller: _ownerCtrl,
                decoration: AppTheme.inputDecoration(
                  'Owner (0x…)',
                  hint: '0x08DEB6b37c3659D045a7Fb93C742f33D1f9B3780',
                ),
              ),
            ),

            // ── Action buttons ──
            // "Lookup DID" fetches existence + risk score + privy hash in one call.
            Wrap(
              spacing: 8,
              runSpacing: 8,
              children: [
                _actionBtn(
                  'Create DID',
                  Icons.person_add,
                  svc.loading,
                  () => svc.createDID(_ownerCtrl.text.trim()),
                ),
                _actionBtn(
                  'Lookup DID',
                  Icons.search,
                  svc.loading,
                  () => svc.lookupDID(_ownerCtrl.text.trim()),
                ),
              ],
            ),
            const SizedBox(height: 16),

            // ── Profile display (populated by lookupDID) ──
            if (svc.profile != null) ...[
              InfoCard(
                title: 'DID Profile',
                child: Column(
                  children: [
                    KVRow(
                      label: 'Owner',
                      value: svc.profile!.owner,
                      mono: true,
                    ),
                    KVRow(
                      label: 'Registered',
                      value: svc.profile!.exists ? 'Yes' : 'No',
                    ),
                    KVRow(
                      label: 'Risk Score',
                      value: svc.profile!.riskScore.toString(),
                    ),
                    KVRow(label: 'Risk Tier', value: svc.profile!.riskTier),
                    if (svc.profile!.privyHash != null)
                      KVRow(
                        label: 'Privy Hash',
                        value: svc.profile!.privyHash!,
                        mono: true,
                      ),
                  ],
                ),
              ),
            ],

            const SizedBox(height: 16),

            // ── Info: Privy auto-linked, risk managed by CRE ──
            InfoCard(
              title: 'Identity Notes',
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text(
                    'Privy credential is automatically linked on login.',
                    style: AppTheme.body,
                  ),
                  const SizedBox(height: 4),
                  Text(
                    'Risk score is updated by CRE workflows based on '
                    'loan and BNPL activity.',
                    style: AppTheme.body,
                  ),
                ],
              ),
            ),

            // ── Loading indicator ──
            if (svc.loading) const LoadingOverlay(label: 'Processing…'),
          ],
        );
      },
    );
  }

  Widget _actionBtn(
    String label,
    IconData icon,
    bool loading,
    VoidCallback onPressed,
  ) {
    return ElevatedButton.icon(
      onPressed: loading ? null : onPressed,
      icon: Icon(icon, size: 18),
      label: Text(label),
    );
  }
}
