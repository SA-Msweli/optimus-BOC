import 'package:flutter/foundation.dart';
import '../models/did_profile.dart';
import 'api_client.dart';

/// Manages DID / identity state for the current user session.
///
/// Builds a local [DIDProfile] aggregate from multiple backend endpoints:
/// - GET /did/{owner}       → { "exists": bool }
/// - GET /did/{owner}/risk  → { "risk_score": "720" }
/// - GET /did/{owner}/privy → { "hash": "0x…" }
class DIDService extends ChangeNotifier {
  final ApiClient _api;

  DIDService(this._api);

  DIDProfile? _profile;
  DIDProfile? get profile => _profile;

  bool _loading = false;
  bool get loading => _loading;

  String? _error;
  String? get error => _error;

  String? _lastTx;
  String? get lastTx => _lastTx;

  /// Create a new DID on-chain for [owner], then refresh the local profile.
  Future<void> createDID(String owner) async {
    _setLoading();
    try {
      final resp = await _api.createDID(owner);
      _lastTx = resp['tx']?.toString();
      // Refresh local aggregate
      await _refreshProfile(owner);
      return;
    } on ApiException catch (e) {
      _error = e.message;
    } catch (e) {
      _error = e.toString();
    }
    _loading = false;
    notifyListeners();
  }

  /// Look up a DID and build a local profile aggregate from multiple calls.
  Future<void> lookupDID(String owner) async {
    _setLoading();
    try {
      await _refreshProfile(owner);
    } on ApiException catch (e) {
      _error = e.message;
      _profile = null;
    } catch (e) {
      _error = e.toString();
      _profile = null;
    }
    _loading = false;
    notifyListeners();
  }

  /// Fetch only the risk score for [owner] and update the local profile.
  Future<void> fetchRiskScore(String owner) async {
    _setLoading();
    try {
      final resp = await _api.getRiskScore(owner);
      final score =
          int.tryParse(resp['risk_score']?.toString() ?? '0') ?? 0;
      if (_profile != null) {
        _profile = _profile!.copyWith(riskScore: score);
      } else {
        _profile = DIDProfile(owner: owner, riskScore: score);
      }
    } on ApiException catch (e) {
      _error = e.message;
    } catch (e) {
      _error = e.toString();
    }
    _loading = false;
    notifyListeners();
  }

  /// Update the on-chain risk profile, then refresh local score.
  Future<void> updateRiskProfile(
      String owner, String newScore, String riskProfileHash) async {
    _setLoading();
    try {
      final resp =
          await _api.updateRiskProfile(owner, newScore, riskProfileHash);
      _lastTx = resp['tx']?.toString();
      await fetchRiskScore(owner);
      return;
    } on ApiException catch (e) {
      _error = e.message;
    } catch (e) {
      _error = e.toString();
    }
    _loading = false;
    notifyListeners();
  }

  /// Link a Privy identity hash to [owner].
  Future<void> linkPrivy(String owner, String hash) async {
    _setLoading();
    try {
      final resp = await _api.linkPrivy(owner, hash);
      _lastTx = resp['tx']?.toString();
    } on ApiException catch (e) {
      _error = e.message;
    } catch (e) {
      _error = e.toString();
    }
    _loading = false;
    notifyListeners();
  }

  /// Fetch the Privy hash linked to [owner] and update the local profile.
  Future<void> fetchPrivyHash(String owner) async {
    _setLoading();
    try {
      final resp = await _api.getPrivyHash(owner);
      final hash = resp['hash']?.toString();
      if (_profile != null) {
        _profile = _profile!.copyWith(privyHash: hash);
      }
    } on ApiException catch (e) {
      _error = e.message;
    } catch (e) {
      _error = e.toString();
    }
    _loading = false;
    notifyListeners();
  }

  void clearError() {
    _error = null;
    notifyListeners();
  }

  // ─── Internal ──────────────────────────────────────────────────────────

  void _setLoading() {
    _loading = true;
    _error = null;
    notifyListeners();
  }

  /// Builds the local aggregate from exists + risk + privy calls.
  Future<void> _refreshProfile(String owner) async {
    // 1. Check existence
    final existsResp = await _api.lookupDID(owner);
    final exists = existsResp['exists'] == true;

    var prof = DIDProfile(owner: owner, exists: exists);

    if (exists) {
      // 2. Fetch risk score (may fail if never set)
      try {
        final riskResp = await _api.getRiskScore(owner);
        final score =
            int.tryParse(riskResp['risk_score']?.toString() ?? '0') ?? 0;
        prof = prof.copyWith(riskScore: score);
      } catch (_) {}

      // 3. Fetch privy hash (may fail if never linked)
      try {
        final privyResp = await _api.getPrivyHash(owner);
        prof = prof.copyWith(privyHash: privyResp['hash']?.toString());
      } catch (_) {}
    }

    _profile = prof;
  }
}
    _error = null;
    notifyListeners();
  }
}
