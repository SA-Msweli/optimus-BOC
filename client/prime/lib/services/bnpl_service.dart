import 'package:flutter/foundation.dart';
import '../models/arrangement.dart';
import 'api_client.dart';

/// Manages BNPL arrangement state via the backend API.
class BNPLService extends ChangeNotifier {
  final ApiClient _api;

  BNPLService(this._api);

  Arrangement? _current;
  Arrangement? get current => _current;

  bool _loading = false;
  bool get loading => _loading;

  String? _error;
  String? get error => _error;

  String? _lastTx;
  String? get lastTx => _lastTx;

  Future<void> createArrangement({
    required String daoId,
    required String recipient,
    required String totalAmount,
    required int startTimestamp,
    required int intervalSeconds,
  }) async {
    _loading = true;
    _error = null;
    notifyListeners();
    try {
      final resp = await _api.createArrangement(
        daoId: daoId,
        recipient: recipient,
        totalAmount: totalAmount,
        startTimestamp: startTimestamp,
        intervalSeconds: intervalSeconds,
      );
      _lastTx = resp['tx']?.toString();
      // Fetch the created arrangement
      final id = resp['arrangement_id']?.toString();
      if (id != null) {
        await fetchArrangement(id);
      }
    } on ApiException catch (e) {
      _error = e.message;
    } catch (e) {
      _error = e.toString();
    }
    _loading = false;
    notifyListeners();
  }

  Future<void> fetchArrangement(String id) async {
    _loading = true;
    _error = null;
    notifyListeners();
    try {
      final resp = await _api.getArrangement(id);
      _current = Arrangement.fromJson(resp);
    } on ApiException catch (e) {
      _error = e.message;
    } catch (e) {
      _error = e.toString();
    }
    _loading = false;
    notifyListeners();
  }

  Future<void> makePayment(String id, int installment, String amount) async {
    _loading = true;
    _error = null;
    notifyListeners();
    try {
      final resp = await _api.makeArrangementPayment(id, installment, amount);
      _lastTx = resp['tx']?.toString();
      await fetchArrangement(id);
    } on ApiException catch (e) {
      _error = e.message;
    } catch (e) {
      _error = e.toString();
    }
    _loading = false;
    notifyListeners();
  }

  Future<void> reschedule(String id, int newStart, int newInterval) async {
    _loading = true;
    _error = null;
    notifyListeners();
    try {
      final resp =
          await _api.rescheduleArrangement(id, newStart, newInterval);
      _lastTx = resp['tx']?.toString();
      await fetchArrangement(id);
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
}
