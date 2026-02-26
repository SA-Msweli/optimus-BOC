import 'package:flutter/foundation.dart';
import 'api_client.dart';

/// Manages token vault deposits, withdrawals, and balance queries.
class VaultService extends ChangeNotifier {
  final ApiClient _api;

  VaultService(this._api);

  String _balance = '0';
  String get balance => _balance;

  bool _loading = false;
  bool get loading => _loading;

  String? _error;
  String? get error => _error;

  String? _lastTx;
  String? get lastTx => _lastTx;

  Future<void> deposit(String token, String amount) async {
    _loading = true;
    _error = null;
    notifyListeners();
    try {
      final resp = await _api.deposit(token, amount);
      _lastTx = resp['tx']?.toString();
      await fetchBalance(token);
    } on ApiException catch (e) {
      _error = e.message;
    } catch (e) {
      _error = e.toString();
    }
    _loading = false;
    notifyListeners();
  }

  Future<void> withdraw(String token, String amount) async {
    _loading = true;
    _error = null;
    notifyListeners();
    try {
      final resp = await _api.withdraw(token, amount);
      _lastTx = resp['tx']?.toString();
      await fetchBalance(token);
    } on ApiException catch (e) {
      _error = e.message;
    } catch (e) {
      _error = e.toString();
    }
    _loading = false;
    notifyListeners();
  }

  Future<void> fetchBalance(String token) async {
    try {
      final resp = await _api.getVaultBalance(token);
      _balance = resp['balance']?.toString() ?? '0';
      notifyListeners();
    } catch (_) {}
  }

  void clearError() {
    _error = null;
    notifyListeners();
  }
}
