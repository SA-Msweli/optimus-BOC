import 'dart:convert';
import 'package:http/http.dart' as http;
import '../config.dart';

/// Callback that returns a fresh JWT access-token (or null).
typedef TokenProvider = Future<String?> Function();

/// Low-level HTTP client that talks to the Optimus backend API.
///
/// Every method returns the decoded JSON body. On non-2xx responses an
/// [ApiException] is thrown with the server error message.
///
/// After authentication, call [setTokenProvider] so that every request
/// carries a `Authorization: Bearer <token>` header automatically.
class ApiClient {
  final http.Client _http;
  final String _base;

  /// When set, every request will include the Privy JWT.
  TokenProvider? _tokenProvider;

  ApiClient({http.Client? client, String? baseUrl})
    : _http = client ?? http.Client(),
      _base = baseUrl ?? AppConfig.backendUrl;

  /// Called by [AuthService] once the user is authenticated.
  void setTokenProvider(TokenProvider provider) {
    _tokenProvider = provider;
  }

  /// Remove the token provider on logout.
  void clearTokenProvider() {
    _tokenProvider = null;
  }

  /// Build the standard headers, injecting the Bearer token when available.
  Future<Map<String, String>> _authHeaders({bool json = false}) async {
    final headers = <String, String>{};
    if (json) headers['Content-Type'] = 'application/json';
    if (_tokenProvider != null) {
      final token = await _tokenProvider!();
      if (token != null) {
        headers['Authorization'] = 'Bearer $token';
      }
    }
    return headers;
  }

  // ─── Health ────────────────────────────────────────────────────────────

  Future<bool> healthCheck() async {
    final r = await _get('/health');
    return r == 'OK' || (r is Map && r['status'] == 'OK');
  }

  // ─── DID ───────────────────────────────────────────────────────────────

  Future<Map<String, dynamic>> createDID(String owner) =>
      _postJson('/did', {'owner': owner});

  Future<Map<String, dynamic>> lookupDID(String owner) =>
      _getJson('/did/$owner');

  Future<Map<String, dynamic>> linkPrivy(String owner, String hash) =>
      _postJson('/did/$owner/link', {'hash': hash});

  Future<Map<String, dynamic>> getPrivyHash(String owner) =>
      _getJson('/did/$owner/privy');

  Future<Map<String, dynamic>> updateRiskProfile(
    String owner,
    String newScore,
    String riskProfileHash,
  ) => _postJson('/did/$owner/risk', {
    'new_score': newScore,
    'risk_profile_hash': riskProfileHash,
  });

  Future<Map<String, dynamic>> getRiskScore(String owner) =>
      _getJson('/did/$owner/risk');

  // ─── BNPL ──────────────────────────────────────────────────────────────

  Future<Map<String, dynamic>> createArrangement({
    required String daoId,
    required String recipient,
    required String totalAmount,
    required int startTimestamp,
    required int intervalSeconds,
  }) => _postJson('/bnpl/arrangements', {
    'dao_id': daoId,
    'recipient': recipient,
    'total_amount': totalAmount,
    'start_timestamp': startTimestamp,
    'interval_seconds': intervalSeconds,
  });

  Future<Map<String, dynamic>> getArrangement(String id) =>
      _getJson('/bnpl/arrangements/$id');

  Future<Map<String, dynamic>> makeArrangementPayment(
    String id,
    int installment,
    String amount,
  ) =>
      _postJson('/bnpl/arrangements/$id/payment', {
        'installment': installment,
        'amount': amount,
      });

  Future<Map<String, dynamic>> activateArrangement(String id) =>
      _postJson('/bnpl/arrangements/$id/activate', {});

  Future<Map<String, dynamic>> applyLateFee(String id, int installment) =>
      _postJson('/bnpl/arrangements/$id/latefee', {'installment': installment});

  Future<Map<String, dynamic>> rescheduleArrangement(
    String id,
    int newStart,
    int newInterval,
  ) => _postJson('/bnpl/arrangements/$id/reschedule', {
    'new_start_timestamp': newStart,
    'new_interval_seconds': newInterval,
  });

  // ─── Loans ─────────────────────────────────────────────────────────────

  Future<Map<String, dynamic>> createLoan({
    required String borrower,
    required String principal,
    required String interestRateBps,
    required String durationSeconds,
  }) => _postJson('/loan', {
    'borrower': borrower,
    'dao_id': '0',
    'principal': principal,
    'interest_rate_bps': interestRateBps,
    'duration_seconds': durationSeconds,
  });

  Future<Map<String, dynamic>> getLoan(String id) => _getJson('/loan/$id');

  Future<Map<String, dynamic>> approveLoan(String id) =>
      _postJson('/loan/$id/approve', {});

  Future<Map<String, dynamic>> makeLoanPayment(String id, String amount) =>
      _postJson('/loan/$id/payment', {'amount': amount});

  Future<Map<String, dynamic>> markLoanDefaulted(String id) =>
      _postJson('/loan/$id/default', {});

  Future<Map<String, dynamic>> getAccruedInterest(String id) =>
      _getJson('/loan/$id/interest');

  Future<Map<String, dynamic>> getAmountOwed(String id) =>
      _getJson('/loan/$id/owed');

  // ─── DAO ───────────────────────────────────────────────────────────────

  Future<Map<String, dynamic>> createDAO({
    required int goal,
    required int votingPeriodDays,
  }) =>
      _postJson('/dao', {'goal': goal, 'voting_period_days': votingPeriodDays});

  Future<Map<String, dynamic>> joinDAO(String daoId, String investment) =>
      _postJson('/dao/$daoId/join', {'investment': investment});

  Future<Map<String, dynamic>> propose(String daoId, String data) =>
      _postJson('/dao/$daoId/propose', {'data': data});

  Future<Map<String, dynamic>> vote(
    String proposalId,
    bool support, {
    String? daoId,
  }) => _postJson('/dao/proposal/$proposalId/vote', {
    'support': support,
    'dao_id': ?daoId,
  });

  Future<Map<String, dynamic>> finalizeProposal(String proposalId) =>
      _postJson('/dao/proposal/$proposalId/finalize', {});

  Future<Map<String, dynamic>> executeProposal(String proposalId) =>
      _postJson('/dao/proposal/$proposalId/execute', {});

  Future<Map<String, dynamic>> setBnplTerms(
    String daoId, {
    required String numInstallments,
    required String minDays,
    required String maxDays,
    required String lateFeeBps,
    required String graceDays,
    required bool rescheduleAllowed,
    required String minDownBps,
  }) => _postJson('/dao/$daoId/bnpl-terms', {
    'num_installments': numInstallments,
    'min_days': minDays,
    'max_days': maxDays,
    'late_fee_bps': lateFeeBps,
    'grace_days': graceDays,
    'reschedule_allowed': rescheduleAllowed,
    'min_down_bps': minDownBps,
  });

  Future<Map<String, dynamic>> getBnplTerms(String daoId) =>
      _getJson('/dao/$daoId/bnpl-terms');

  Future<Map<String, dynamic>> getTreasuryBalance(String daoId) =>
      _getJson('/dao/$daoId/treasury');

  Future<Map<String, dynamic>> creditTreasury(String daoId, String amount) =>
      _postJson('/dao/$daoId/treasury/credit', {'amount': amount});

  // ─── Token Vault ───────────────────────────────────────────────────────

  Future<Map<String, dynamic>> deposit(String token, String amount) =>
      _postJson('/vault/deposit', {'token': token, 'amount': amount});

  Future<Map<String, dynamic>> withdraw(String token, String amount) =>
      _postJson('/vault/withdraw', {'token': token, 'amount': amount});

  Future<Map<String, dynamic>> getVaultBalance(String token) =>
      _getJson('/vault/balance/$token');

  // ─── Internals ─────────────────────────────────────────────────────────

  Future<dynamic> _get(String path) async {
    final uri = Uri.parse('$_base$path');
    final headers = await _authHeaders();
    final resp = await _http.get(uri, headers: headers);
    if (resp.statusCode >= 200 && resp.statusCode < 300) {
      try {
        return jsonDecode(resp.body);
      } catch (_) {
        return resp.body;
      }
    }
    throw ApiException(resp.statusCode, resp.body);
  }

  Future<Map<String, dynamic>> _getJson(String path) async {
    final body = await _get(path);
    return body is Map<String, dynamic> ? body : {'raw': body.toString()};
  }

  Future<Map<String, dynamic>> _postJson(
    String path,
    Map<String, dynamic> body,
  ) async {
    final uri = Uri.parse('$_base$path');
    final headers = await _authHeaders(json: true);
    final resp = await _http.post(
      uri,
      headers: headers,
      body: jsonEncode(body),
    );
    if (resp.statusCode >= 200 && resp.statusCode < 300) {
      return jsonDecode(resp.body) as Map<String, dynamic>;
    }
    throw ApiException(resp.statusCode, resp.body);
  }
}

/// Exception thrown on non-2xx API responses.
class ApiException implements Exception {
  final int statusCode;
  final String body;
  ApiException(this.statusCode, this.body);

  String get message {
    try {
      final decoded = jsonDecode(body);
      if (decoded is Map && decoded.containsKey('error')) {
        return decoded['error'].toString();
      }
    } catch (_) {}
    return body;
  }

  @override
  String toString() => 'ApiException($statusCode): $message';
}
