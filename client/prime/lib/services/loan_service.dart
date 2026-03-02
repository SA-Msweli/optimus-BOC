import 'package:flutter/foundation.dart';
import '../models/loan.dart';
import 'api_client.dart';

/// Manages loan state via the backend API.
class LoanService extends ChangeNotifier {
  final ApiClient _api;

  LoanService(this._api);

  Loan? _current;
  Loan? get current => _current;

  String _accruedInterest = '0';
  String get accruedInterest => _accruedInterest;

  String _amountOwed = '0';
  String get amountOwed => _amountOwed;

  bool _loading = false;
  bool get loading => _loading;

  String? _error;
  String? get error => _error;

  String? _lastTx;
  String? get lastTx => _lastTx;

  Future<void> createLoan({
    required String borrower,
    required String principal,
    required String interestRateBps,
    required String durationSeconds,
  }) async {
    _loading = true;
    _error = null;
    notifyListeners();
    try {
      final resp = await _api.createLoan(
        borrower: borrower,
        principal: principal,
        interestRateBps: interestRateBps,
        durationSeconds: durationSeconds,
      );
      _lastTx = resp['tx']?.toString();
      final id = resp['loan_id']?.toString();
      if (id != null) {
        await fetchLoan(id);
      }
    } on ApiException catch (e) {
      _error = e.message;
    } catch (e) {
      _error = e.toString();
    }
    _loading = false;
    notifyListeners();
  }

  Future<void> fetchLoan(String id) async {
    _loading = true;
    _error = null;
    notifyListeners();
    try {
      final resp = await _api.getLoan(id);
      _current = Loan.fromJson(resp);
    } on ApiException catch (e) {
      _error = e.message;
    } catch (e) {
      _error = e.toString();
    }
    _loading = false;
    notifyListeners();
  }

  Future<void> approveLoan(String id) async {
    _loading = true;
    _error = null;
    notifyListeners();
    try {
      final resp = await _api.approveLoan(id);
      _lastTx = resp['tx']?.toString();
      await fetchLoan(id);
    } on ApiException catch (e) {
      _error = e.message;
    } catch (e) {
      _error = e.toString();
    }
    _loading = false;
    notifyListeners();
  }

  Future<void> makePayment(String id, String amount) async {
    _loading = true;
    _error = null;
    notifyListeners();
    try {
      final resp = await _api.makeLoanPayment(id, amount);
      _lastTx = resp['tx']?.toString();
      await fetchLoan(id);
    } on ApiException catch (e) {
      _error = e.message;
    } catch (e) {
      _error = e.toString();
    }
    _loading = false;
    notifyListeners();
  }

  Future<void> fetchAccruedInterest(String id) async {
    try {
      final resp = await _api.getAccruedInterest(id);
      _accruedInterest = resp['accrued_interest']?.toString() ?? '0';
      notifyListeners();
    } catch (_) {}
  }

  Future<void> fetchAmountOwed(String id) async {
    try {
      final resp = await _api.getAmountOwed(id);
      _amountOwed = resp['amount_owed']?.toString() ?? '0';
      notifyListeners();
    } catch (_) {}
  }

  void clearError() {
    _error = null;
    notifyListeners();
  }
}
