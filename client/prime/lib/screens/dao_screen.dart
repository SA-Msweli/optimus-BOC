import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import '../services/dao_service.dart';
import '../theme.dart';
import '../widgets/shared.dart';

/// Screen for DAO governance: create DAO, join, propose, vote, finalize,
/// execute, set BNPL terms, and manage treasury.
///
/// Proposal data is built from structured fields (token, amount, recipient)
/// for treasury withdrawal proposals, or raw hex for generic proposals.
class DAOScreen extends StatefulWidget {
  const DAOScreen({super.key});

  @override
  State<DAOScreen> createState() => _DAOScreenState();
}

class _DAOScreenState extends State<DAOScreen>
    with SingleTickerProviderStateMixin {
  late TabController _tabs;

  // ── Create DAO ──
  int _selectedGoal = 0;
  final _votingPeriodCtrl = TextEditingController();

  // ── Join DAO ──
  final _joinDaoIdCtrl = TextEditingController();
  final _investmentCtrl = TextEditingController();

  // ── Proposals ──
  final _proposeDaoIdCtrl = TextEditingController();
  // Structured proposal fields (for treasury withdrawals)
  final _propTokenCtrl = TextEditingController();
  final _propAmountCtrl = TextEditingController();
  final _propRecipientCtrl = TextEditingController();
  final _propRawDataCtrl = TextEditingController();
  bool _useRawProposal = false;

  // ── Vote ──
  final _voteProposalIdCtrl = TextEditingController();
  final _voteDaoIdCtrl = TextEditingController();
  bool _voteSupport = true;

  // ── Finalize / Execute ──
  final _actionProposalIdCtrl = TextEditingController();

  // ── BNPL Terms ──
  final _termsDaoIdCtrl = TextEditingController();
  final _numInstCtrl = TextEditingController();
  final _minDaysCtrl = TextEditingController();
  final _maxDaysCtrl = TextEditingController();
  final _lateFeeBpsCtrl = TextEditingController();
  final _graceDaysCtrl = TextEditingController();
  final _minDownBpsCtrl = TextEditingController();
  bool _rescheduleAllowed = true;

  // ── Treasury ──
  final _treasuryDaoIdCtrl = TextEditingController();
  final _creditAmountCtrl = TextEditingController();

  @override
  void initState() {
    super.initState();
    _tabs = TabController(length: 4, vsync: this);
  }

  @override
  void dispose() {
    _tabs.dispose();
    _votingPeriodCtrl.dispose();
    _joinDaoIdCtrl.dispose();
    _investmentCtrl.dispose();
    _proposeDaoIdCtrl.dispose();
    _propTokenCtrl.dispose();
    _propAmountCtrl.dispose();
    _propRecipientCtrl.dispose();
    _propRawDataCtrl.dispose();
    _voteProposalIdCtrl.dispose();
    _voteDaoIdCtrl.dispose();
    _actionProposalIdCtrl.dispose();
    _termsDaoIdCtrl.dispose();
    _numInstCtrl.dispose();
    _minDaysCtrl.dispose();
    _maxDaysCtrl.dispose();
    _lateFeeBpsCtrl.dispose();
    _graceDaysCtrl.dispose();
    _minDownBpsCtrl.dispose();
    _treasuryDaoIdCtrl.dispose();
    _creditAmountCtrl.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Consumer<DAOService>(
      builder: (context, svc, _) {
        return Column(
          children: [
            Material(
              color: AppTheme.primary,
              child: TabBar(
                controller: _tabs,
                labelColor: Colors.white,
                unselectedLabelColor: Colors.white70,
                indicatorColor: Colors.white,
                tabs: const [
                  Tab(text: 'DAO'),
                  Tab(text: 'Proposals'),
                  Tab(text: 'Terms'),
                  Tab(text: 'Treasury'),
                ],
              ),
            ),
            Expanded(
              child: TabBarView(
                controller: _tabs,
                children: [
                  _buildDAOTab(svc),
                  _buildProposalTab(svc),
                  _buildTermsTab(svc),
                  _buildTreasuryTab(svc),
                ],
              ),
            ),
          ],
        );
      },
    );
  }

  // ═══════════════════════════════════════════════════════════════════════
  // TAB 1: Create & Join DAO
  // ═══════════════════════════════════════════════════════════════════════

  Widget _buildDAOTab(DAOService svc) {
    return ListView(
      padding: const EdgeInsets.all(16),
      children: [
        if (svc.error != null)
          ErrorBanner(message: svc.error!, onDismiss: () => svc.clearError()),
        if (svc.lastTx != null)
          SuccessBanner(message: 'TX: ${truncateAddress(svc.lastTx!)}'),

        // Create DAO (no name field — contract stores goal enum only)
        InfoCard(
          title: 'Create DAO',
          child: Column(
            children: [
              DropdownButtonFormField<int>(
                initialValue: _selectedGoal,
                decoration: AppTheme.inputDecoration('Goal'),
                items: const [
                  DropdownMenuItem(value: 0, child: Text('Savings')),
                  DropdownMenuItem(value: 1, child: Text('Lending')),
                  DropdownMenuItem(value: 2, child: Text('Investment')),
                ],
                onChanged: (v) => setState(() => _selectedGoal = v ?? 0),
              ),
              const SizedBox(height: 8),
              TextField(
                controller: _votingPeriodCtrl,
                decoration: AppTheme.inputDecoration(
                  'Voting Period (days)',
                  hint: '7',
                ),
                keyboardType: TextInputType.number,
              ),
              const SizedBox(height: 12),
              SizedBox(
                width: double.infinity,
                child: ElevatedButton.icon(
                  onPressed: svc.loading
                      ? null
                      : () => svc.createDAO(
                          goal: _selectedGoal,
                          votingPeriodDays:
                              int.tryParse(_votingPeriodCtrl.text) ?? 7,
                        ),
                  icon: const Icon(Icons.group_add),
                  label: const Text('Create DAO'),
                ),
              ),
            ],
          ),
        ),

        // DAO create result
        if (svc.currentDAO != null) ...[
          InfoCard(
            title: 'Created DAO',
            child: Column(
              children: [
                KVRow(
                  label: 'DAO ID',
                  value: svc.currentDAO!['dao_id']?.toString() ?? 'N/A',
                ),
                KVRow(
                  label: 'TX',
                  value: svc.currentDAO!['tx']?.toString() ?? 'N/A',
                  mono: true,
                ),
              ],
            ),
          ),
        ],

        const Divider(height: 24),

        // Register Member (Join DAO)
        InfoCard(
          title: 'Register Member',
          child: Column(
            children: [
              TextField(
                controller: _joinDaoIdCtrl,
                decoration: AppTheme.inputDecoration('DAO ID', hint: '1'),
                keyboardType: TextInputType.number,
              ),
              const SizedBox(height: 8),
              TextField(
                controller: _investmentCtrl,
                decoration: AppTheme.inputDecoration(
                  'Voting Weight (accounting units)',
                  hint: '1000 — determines vote weight, not ETH transfer',
                ),
                keyboardType: TextInputType.number,
              ),
              const SizedBox(height: 12),
              SizedBox(
                width: double.infinity,
                child: ElevatedButton.icon(
                  onPressed: svc.loading
                      ? null
                      : () => svc.joinDAO(
                          _joinDaoIdCtrl.text.trim(),
                          _investmentCtrl.text.trim(),
                        ),
                  icon: const Icon(Icons.person_add),
                  label: const Text('Register'),
                ),
              ),
            ],
          ),
        ),

        if (svc.loading) const LoadingOverlay(label: 'Processing…'),
      ],
    );
  }

  // ═══════════════════════════════════════════════════════════════════════
  // TAB 2: Proposals (propose, vote, finalize, execute)
  // ═══════════════════════════════════════════════════════════════════════

  Widget _buildProposalTab(DAOService svc) {
    return ListView(
      padding: const EdgeInsets.all(16),
      children: [
        if (svc.error != null)
          ErrorBanner(message: svc.error!, onDismiss: () => svc.clearError()),
        if (svc.lastTx != null)
          SuccessBanner(message: 'TX: ${truncateAddress(svc.lastTx!)}'),

        // Propose — structured form for treasury proposals
        InfoCard(
          title: 'Create Proposal',
          child: Column(
            children: [
              TextField(
                controller: _proposeDaoIdCtrl,
                decoration: AppTheme.inputDecoration('DAO ID', hint: '1'),
                keyboardType: TextInputType.number,
              ),
              const SizedBox(height: 8),
              SwitchListTile(
                title: Text(
                  _useRawProposal ? 'Raw Hex Data' : 'Treasury Withdrawal',
                  style: AppTheme.subheading,
                ),
                subtitle: Text(
                  _useRawProposal
                      ? 'Enter ABI-encoded bytes directly'
                      : 'Withdraw tokens from vault to recipient',
                ),
                value: _useRawProposal,
                onChanged: (v) => setState(() => _useRawProposal = v),
              ),
              const SizedBox(height: 8),
              if (!_useRawProposal) ...[
                TextField(
                  controller: _propTokenCtrl,
                  decoration: AppTheme.inputDecoration(
                    'Token Address (0x…)',
                    hint: '0x… (ERC-20 token to withdraw)',
                  ),
                ),
                const SizedBox(height: 8),
                TextField(
                  controller: _propAmountCtrl,
                  decoration: AppTheme.inputDecoration(
                    'Amount (wei)',
                    hint: '1000000000000000000',
                  ),
                  keyboardType: TextInputType.number,
                ),
                const SizedBox(height: 8),
                TextField(
                  controller: _propRecipientCtrl,
                  decoration: AppTheme.inputDecoration(
                    'Recipient Address (0x…)',
                    hint: '0x… (receives the tokens)',
                  ),
                ),
              ] else ...[
                TextField(
                  controller: _propRawDataCtrl,
                  decoration: AppTheme.inputDecoration(
                    'Proposal Data (hex)',
                    hint: '0x…',
                  ),
                  maxLines: 3,
                ),
              ],
              const SizedBox(height: 12),
              SizedBox(
                width: double.infinity,
                child: ElevatedButton.icon(
                  onPressed: svc.loading
                      ? null
                      : () {
                          final daoId = _proposeDaoIdCtrl.text.trim();
                          if (_useRawProposal) {
                            svc.propose(daoId, _propRawDataCtrl.text.trim());
                          } else {
                            svc.proposeTreasuryWithdrawal(
                              daoId,
                              token: _propTokenCtrl.text.trim(),
                              amount: _propAmountCtrl.text.trim(),
                              recipient: _propRecipientCtrl.text.trim(),
                            );
                          }
                        },
                  icon: const Icon(Icons.description),
                  label: const Text('Propose'),
                ),
              ),
            ],
          ),
        ),

        const Divider(height: 24),

        // Vote (DAO ID is required for membership check)
        InfoCard(
          title: 'Cast Vote',
          child: Column(
            children: [
              TextField(
                controller: _voteProposalIdCtrl,
                decoration: AppTheme.inputDecoration('Proposal ID', hint: '1'),
                keyboardType: TextInputType.number,
              ),
              const SizedBox(height: 8),
              TextField(
                controller: _voteDaoIdCtrl,
                decoration: AppTheme.inputDecoration(
                  'DAO ID',
                  hint: 'Required for membership verification',
                ),
                keyboardType: TextInputType.number,
              ),
              const SizedBox(height: 8),
              SwitchListTile(
                title: Text(
                  _voteSupport ? 'FOR' : 'AGAINST',
                  style: AppTheme.subheading.copyWith(
                    color: _voteSupport ? AppTheme.secondary : AppTheme.error,
                  ),
                ),
                value: _voteSupport,
                onChanged: (v) => setState(() => _voteSupport = v),
                activeThumbColor: AppTheme.secondary,
              ),
              const SizedBox(height: 8),
              SizedBox(
                width: double.infinity,
                child: ElevatedButton.icon(
                  onPressed: svc.loading
                      ? null
                      : () {
                          final daoId = _voteDaoIdCtrl.text.trim();
                          if (daoId.isEmpty) {
                            svc.setError(
                              'DAO ID is required to verify membership',
                            );
                            return;
                          }
                          svc.vote(
                            _voteProposalIdCtrl.text.trim(),
                            _voteSupport,
                            daoId: daoId,
                          );
                        },
                  icon: const Icon(Icons.how_to_vote),
                  label: const Text('Vote'),
                ),
              ),
            ],
          ),
        ),

        const Divider(height: 24),

        // Finalize / Execute
        InfoCard(
          title: 'Finalize / Execute',
          child: Column(
            children: [
              TextField(
                controller: _actionProposalIdCtrl,
                decoration: AppTheme.inputDecoration('Proposal ID', hint: '1'),
                keyboardType: TextInputType.number,
              ),
              const SizedBox(height: 12),
              Row(
                children: [
                  Expanded(
                    child: ElevatedButton.icon(
                      onPressed: svc.loading
                          ? null
                          : () => svc.finalizeProposal(
                              _actionProposalIdCtrl.text.trim(),
                            ),
                      icon: const Icon(Icons.gavel, size: 18),
                      label: const Text('Finalize'),
                    ),
                  ),
                  const SizedBox(width: 8),
                  Expanded(
                    child: ElevatedButton.icon(
                      onPressed: svc.loading
                          ? null
                          : () => svc.executeProposal(
                              _actionProposalIdCtrl.text.trim(),
                            ),
                      icon: const Icon(Icons.play_arrow, size: 18),
                      label: const Text('Execute'),
                    ),
                  ),
                ],
              ),
            ],
          ),
        ),

        if (svc.loading) const LoadingOverlay(label: 'Processing…'),
      ],
    );
  }

  // ═══════════════════════════════════════════════════════════════════════
  // TAB 3: BNPL Terms
  // ═══════════════════════════════════════════════════════════════════════

  Widget _buildTermsTab(DAOService svc) {
    return ListView(
      padding: const EdgeInsets.all(16),
      children: [
        if (svc.error != null)
          ErrorBanner(message: svc.error!, onDismiss: () => svc.clearError()),
        if (svc.lastTx != null)
          SuccessBanner(message: 'TX: ${truncateAddress(svc.lastTx!)}'),

        // Fetch terms
        InfoCard(
          title: 'View BNPL Terms',
          child: Row(
            children: [
              Expanded(
                child: TextField(
                  controller: _termsDaoIdCtrl,
                  decoration: AppTheme.inputDecoration('DAO ID', hint: '1'),
                  keyboardType: TextInputType.number,
                ),
              ),
              const SizedBox(width: 8),
              ElevatedButton(
                onPressed: svc.loading
                    ? null
                    : () => svc.fetchBnplTerms(_termsDaoIdCtrl.text.trim()),
                child: const Text('Fetch'),
              ),
            ],
          ),
        ),

        if (svc.bnplTerms != null)
          InfoCard(
            title: 'Current Terms',
            child: Column(
              children: [
                for (final e in svc.bnplTerms!.entries)
                  KVRow(label: e.key, value: e.value.toString()),
              ],
            ),
          ),

        const Divider(height: 24),

        // Set terms
        InfoCard(
          title: 'Set BNPL Terms',
          child: Column(
            children: [
              TextField(
                controller: _numInstCtrl,
                decoration: AppTheme.inputDecoration(
                  'Num Installments',
                  hint: '4',
                ),
                keyboardType: TextInputType.number,
              ),
              const SizedBox(height: 8),
              TextField(
                controller: _minDaysCtrl,
                decoration: AppTheme.inputDecoration('Min Days', hint: '7'),
                keyboardType: TextInputType.number,
              ),
              const SizedBox(height: 8),
              TextField(
                controller: _maxDaysCtrl,
                decoration: AppTheme.inputDecoration('Max Days', hint: '90'),
                keyboardType: TextInputType.number,
              ),
              const SizedBox(height: 8),
              TextField(
                controller: _lateFeeBpsCtrl,
                decoration: AppTheme.inputDecoration(
                  'Late Fee (bps)',
                  hint: '200',
                ),
                keyboardType: TextInputType.number,
              ),
              const SizedBox(height: 8),
              TextField(
                controller: _graceDaysCtrl,
                decoration: AppTheme.inputDecoration('Grace Days', hint: '3'),
                keyboardType: TextInputType.number,
              ),
              const SizedBox(height: 8),
              TextField(
                controller: _minDownBpsCtrl,
                decoration: AppTheme.inputDecoration(
                  'Min Down Payment (bps)',
                  hint: '1000',
                ),
                keyboardType: TextInputType.number,
              ),
              const SizedBox(height: 8),
              SwitchListTile(
                title: const Text('Reschedule Allowed'),
                value: _rescheduleAllowed,
                onChanged: (v) => setState(() => _rescheduleAllowed = v),
              ),
              const SizedBox(height: 12),
              SizedBox(
                width: double.infinity,
                child: ElevatedButton.icon(
                  onPressed: svc.loading
                      ? null
                      : () => svc.setBnplTerms(
                          _termsDaoIdCtrl.text.trim(),
                          numInstallments: _numInstCtrl.text.trim(),
                          minDays: _minDaysCtrl.text.trim(),
                          maxDays: _maxDaysCtrl.text.trim(),
                          lateFeeBps: _lateFeeBpsCtrl.text.trim(),
                          graceDays: _graceDaysCtrl.text.trim(),
                          rescheduleAllowed: _rescheduleAllowed,
                          minDownBps: _minDownBpsCtrl.text.trim(),
                        ),
                  icon: const Icon(Icons.save),
                  label: const Text('Set Terms'),
                ),
              ),
            ],
          ),
        ),

        if (svc.loading) const LoadingOverlay(label: 'Processing…'),
      ],
    );
  }

  // ═══════════════════════════════════════════════════════════════════════
  // TAB 4: Treasury
  // ═══════════════════════════════════════════════════════════════════════

  Widget _buildTreasuryTab(DAOService svc) {
    return ListView(
      padding: const EdgeInsets.all(16),
      children: [
        if (svc.error != null)
          ErrorBanner(message: svc.error!, onDismiss: () => svc.clearError()),
        if (svc.lastTx != null)
          SuccessBanner(message: 'TX: ${truncateAddress(svc.lastTx!)}'),

        InfoCard(
          title: 'DAO Treasury',
          child: Column(
            children: [
              TextField(
                controller: _treasuryDaoIdCtrl,
                decoration: AppTheme.inputDecoration('DAO ID', hint: '1'),
                keyboardType: TextInputType.number,
              ),
              const SizedBox(height: 8),
              ElevatedButton(
                onPressed: svc.loading
                    ? null
                    : () => svc.fetchTreasuryBalance(
                        _treasuryDaoIdCtrl.text.trim(),
                      ),
                child: const Text('Fetch Balance'),
              ),
              const SizedBox(height: 12),
              KVRow(label: 'Balance (wei)', value: svc.treasuryBalance),
            ],
          ),
        ),

        const Divider(height: 24),

        InfoCard(
          title: 'Credit Treasury',
          child: Column(
            children: [
              TextField(
                controller: _creditAmountCtrl,
                decoration: AppTheme.inputDecoration(
                  'Amount (wei)',
                  hint: '1000000000000000000',
                ),
                keyboardType: TextInputType.number,
              ),
              const SizedBox(height: 12),
              SizedBox(
                width: double.infinity,
                child: ElevatedButton.icon(
                  onPressed: svc.loading
                      ? null
                      : () => svc.creditTreasury(
                          _treasuryDaoIdCtrl.text.trim(),
                          _creditAmountCtrl.text.trim(),
                        ),
                  icon: const Icon(Icons.add_card),
                  label: const Text('Credit'),
                ),
              ),
            ],
          ),
        ),

        if (svc.loading) const LoadingOverlay(label: 'Processing…'),
      ],
    );
  }
}
