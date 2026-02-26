import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import '../services/auth_service.dart';
import '../services/did_service.dart';
import '../theme.dart';
import '../widgets/shared.dart';

/// Screen for DID identity management: create, lookup, risk score, privy linking.
class DIDScreen extends StatefulWidget {
  const DIDScreen({super.key});

  @override
  State<DIDScreen> createState() => _DIDScreenState();
}

class _DIDScreenState extends State<DIDScreen> {
  final _ownerCtrl = TextEditingController();
  final _privyHashCtrl = TextEditingController();
  final _riskScoreCtrl = TextEditingController();
  final _riskHashCtrl = TextEditingController();
  bool _prefilled = false;

  @override
  void dispose() {
    _ownerCtrl.dispose();
    _privyHashCtrl.dispose();
    _riskScoreCtrl.dispose();
    _riskHashCtrl.dispose();
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

    return Scaffold(
      appBar: AppBar(title: const Text('Identity (DID)')),
      body: Consumer<DIDService>(
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
                  _actionBtn(
                    'Fetch Risk Score',
                    Icons.assessment,
                    svc.loading,
                    () => svc.fetchRiskScore(_ownerCtrl.text.trim()),
                  ),
                ],
              ),
              const SizedBox(height: 16),

              // ── Profile display ──
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

              const Divider(height: 32),

              // ── Link Privy ──
              InfoCard(
                title: 'Link Privy Identity',
                child: Column(
                  children: [
                    TextField(
                      controller: _privyHashCtrl,
                      decoration: AppTheme.inputDecoration(
                        'Privy Hash',
                        hint: '0xabc…',
                      ),
                    ),
                    const SizedBox(height: 10),
                    SizedBox(
                      width: double.infinity,
                      child: ElevatedButton.icon(
                        onPressed: svc.loading
                            ? null
                            : () => svc.linkPrivy(
                                _ownerCtrl.text.trim(),
                                _privyHashCtrl.text.trim(),
                              ),
                        icon: const Icon(Icons.link),
                        label: const Text('Link'),
                      ),
                    ),
                  ],
                ),
              ),

              // ── Fetch Privy Hash ──
              _actionBtn(
                'Fetch Privy Hash',
                Icons.tag,
                svc.loading,
                () => svc.fetchPrivyHash(_ownerCtrl.text.trim()),
              ),

              const Divider(height: 32),

              // ── Update Risk Profile ──
              InfoCard(
                title: 'Update Risk Profile',
                child: Column(
                  children: [
                    TextField(
                      controller: _riskScoreCtrl,
                      decoration: AppTheme.inputDecoration(
                        'New Score',
                        hint: '750',
                      ),
                      keyboardType: TextInputType.number,
                    ),
                    const SizedBox(height: 10),
                    TextField(
                      controller: _riskHashCtrl,
                      decoration: AppTheme.inputDecoration(
                        'Risk Profile Hash',
                        hint: '0x…',
                      ),
                    ),
                    const SizedBox(height: 10),
                    SizedBox(
                      width: double.infinity,
                      child: ElevatedButton.icon(
                        onPressed: svc.loading
                            ? null
                            : () => svc.updateRiskProfile(
                                _ownerCtrl.text.trim(),
                                _riskScoreCtrl.text.trim(),
                                _riskHashCtrl.text.trim(),
                              ),
                        icon: const Icon(Icons.update),
                        label: const Text('Update Risk'),
                      ),
                    ),
                  ],
                ),
              ),

              // ── Loading indicator ──
              if (svc.loading) const LoadingOverlay(label: 'Processing…'),
            ],
          );
        },
      ),
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
