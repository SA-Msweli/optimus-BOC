import 'package:flutter/foundation.dart';
import 'api_client.dart';

/// Manages DAO, proposals, voting, BNPL terms, and treasury via backend API.
class DAOService extends ChangeNotifier {
  final ApiClient _api;

  DAOService(this._api);

  Map<String, dynamic>? _currentDAO;
  Map<String, dynamic>? get currentDAO => _currentDAO;

  Map<String, dynamic>? _bnplTerms;
  Map<String, dynamic>? get bnplTerms => _bnplTerms;

  String _treasuryBalance = '0';
  String get treasuryBalance => _treasuryBalance;

  bool _loading = false;
  bool get loading => _loading;

  String? _error;
  String? get error => _error;

  String? _lastTx;
  String? get lastTx => _lastTx;

  // ─── DAO lifecycle ─────────────────────────────────────────────────────

  Future<void> createDAO({
    required int goal,
    required int votingPeriodDays,
  }) async {
    _setLoading();
    try {
      final resp =
          await _api.createDAO(goal: goal, votingPeriodDays: votingPeriodDays);
      _lastTx = resp['tx']?.toString();
      _currentDAO = resp;
    } on ApiException catch (e) {
      _error = e.message;
    } catch (e) {
      _error = e.toString();
    }
    _loading = false;
    notifyListeners();
  }

  Future<void> joinDAO(String daoId, String investment) async {
    _setLoading();
    try {
      final resp = await _api.joinDAO(daoId, investment);
      _lastTx = resp['tx']?.toString();
    } on ApiException catch (e) {
      _error = e.message;
    } catch (e) {
      _error = e.toString();
    }
    _loading = false;
    notifyListeners();
  }

  // ─── Proposals ─────────────────────────────────────────────────────────

  Future<void> propose(String daoId, String data) async {
    _setLoading();
    try {
      final resp = await _api.propose(daoId, data);
      _lastTx = resp['tx']?.toString();
    } on ApiException catch (e) {
      _error = e.message;
    } catch (e) {
      _error = e.toString();
    }
    _loading = false;
    notifyListeners();
  }

  /// Build structured treasury withdrawal proposal data and submit.
  /// The contract decodes `abi.encode(address token, uint256 amount, address recipient)`.
  Future<void> proposeTreasuryWithdrawal(
    String daoId, {
    required String token,
    required String amount,
    required String recipient,
  }) async {
    if (token.isEmpty || amount.isEmpty || recipient.isEmpty) {
      _error = 'Token, amount, and recipient are all required';
      notifyListeners();
      return;
    }
    // Build ABI-encoded data: pad address (20→32 bytes), uint256 (32 bytes), address (20→32 bytes)
    final tokenHex = token.replaceFirst('0x', '').toLowerCase().padLeft(64, '0');
    final amtBig = BigInt.tryParse(amount);
    if (amtBig == null) {
      _error = 'Invalid amount';
      notifyListeners();
      return;
    }
    final amtHex = amtBig.toRadixString(16).padLeft(64, '0');
    final recipientHex = recipient.replaceFirst('0x', '').toLowerCase().padLeft(64, '0');
    final data = '0x$tokenHex$amtHex$recipientHex';
    await propose(daoId, data);
  }

  Future<void> vote(String proposalId, bool support, {String? daoId}) async {
    _setLoading();
    try {
      final resp = await _api.vote(proposalId, support, daoId: daoId);
      _lastTx = resp['tx']?.toString();
    } on ApiException catch (e) {
      _error = e.message;
    } catch (e) {
      _error = e.toString();
    }
    _loading = false;
    notifyListeners();
  }

  Future<void> finalizeProposal(String proposalId) async {
    _setLoading();
    try {
      final resp = await _api.finalizeProposal(proposalId);
      _lastTx = resp['tx']?.toString();
    } on ApiException catch (e) {
      _error = e.message;
    } catch (e) {
      _error = e.toString();
    }
    _loading = false;
    notifyListeners();
  }

  Future<void> executeProposal(String proposalId) async {
    _setLoading();
    try {
      final resp = await _api.executeProposal(proposalId);
      _lastTx = resp['tx']?.toString();
    } on ApiException catch (e) {
      _error = e.message;
    } catch (e) {
      _error = e.toString();
    }
    _loading = false;
    notifyListeners();
  }

  // ─── BNPL Terms ────────────────────────────────────────────────────────

  Future<void> setBnplTerms(
    String daoId, {
    required String numInstallments,
    required String minDays,
    required String maxDays,
    required String lateFeeBps,
    required String graceDays,
    required bool rescheduleAllowed,
    required String minDownBps,
  }) async {
    _setLoading();
    try {
      final resp = await _api.setBnplTerms(
        daoId,
        numInstallments: numInstallments,
        minDays: minDays,
        maxDays: maxDays,
        lateFeeBps: lateFeeBps,
        graceDays: graceDays,
        rescheduleAllowed: rescheduleAllowed,
        minDownBps: minDownBps,
      );
      _lastTx = resp['tx']?.toString();
      await fetchBnplTerms(daoId);
    } on ApiException catch (e) {
      _error = e.message;
    } catch (e) {
      _error = e.toString();
    }
    _loading = false;
    notifyListeners();
  }

  Future<void> fetchBnplTerms(String daoId) async {
    try {
      _bnplTerms = await _api.getBnplTerms(daoId);
      notifyListeners();
    } catch (_) {}
  }

  // ─── Treasury ──────────────────────────────────────────────────────────

  Future<void> fetchTreasuryBalance(String daoId) async {
    try {
      final resp = await _api.getTreasuryBalance(daoId);
      _treasuryBalance = resp['balance']?.toString() ?? '0';
      notifyListeners();
    } catch (_) {}
  }

  Future<void> creditTreasury(String daoId, String amount) async {
    _setLoading();
    try {
      final resp = await _api.creditTreasury(daoId, amount);
      _lastTx = resp['tx']?.toString();
      await fetchTreasuryBalance(daoId);
    } on ApiException catch (e) {
      _error = e.message;
    } catch (e) {
      _error = e.toString();
    }
    _loading = false;
    notifyListeners();
  }

  // ─── Helpers ───────────────────────────────────────────────────────────

  void _setLoading() {
    _loading = true;
    _error = null;
    notifyListeners();
  }

  void clearError() {
    _error = null;
    notifyListeners();
  }

  void setError(String msg) {
    _error = msg;
    notifyListeners();
  }
}
