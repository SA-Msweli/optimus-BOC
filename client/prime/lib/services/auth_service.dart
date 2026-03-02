import 'dart:convert';

import 'package:crypto/crypto.dart';
import 'package:flutter/foundation.dart';
import 'package:privy_flutter/privy_flutter.dart';

import '../config.dart';
import 'api_client.dart';

/// Authentication states the app can be in.
enum AuthStatus {
  /// SDK not yet initialised / checking persisted session.
  unknown,

  /// No active session – show login screen.
  unauthenticated,

  /// OTP sent – waiting for the user to enter the code.
  awaitingOtp,

  /// Fully authenticated – user has a Privy session + embedded wallet.
  authenticated,
}

/// Wraps the Privy Flutter SDK (v0.6.0) and exposes a [ChangeNotifier]-friendly
/// authentication surface for the rest of the Optimus client.
///
/// Responsibilities:
/// 1. Initialise Privy SDK on app start.
/// 2. Manage email-OTP and social login flows.
/// 3. Create / surface the user's embedded wallet address.
/// 4. Provide a JWT access-token getter for [ApiClient].
/// 5. Auto-link the Privy credential hash to the user's on-chain DID.
class AuthService extends ChangeNotifier {
  AuthService(this._api);

  final ApiClient _api;

  // ── Privy SDK instance ─────────────────────────────────────────────────
  Privy? _privy;

  /// Internal accessor that asserts the SDK has been initialised.
  Privy get _sdk {
    assert(_privy != null, 'Privy SDK has not been initialised');
    return _privy!;
  }

  // ── Observable state ───────────────────────────────────────────────────
  AuthStatus _status = AuthStatus.unknown;
  AuthStatus get status => _status;

  PrivyUser? _user;
  PrivyUser? get user => _user;

  /// The embedded-wallet address produced by Privy (used as the DID owner).
  String? _walletAddress;
  String? get walletAddress => _walletAddress;

  /// Privy user id (e.g. `did:privy:clxyz…`).
  String? _privyUserId;
  String? get privyUserId => _privyUserId;

  String? _email;
  String? get email => _email;

  String? _error;
  String? get error => _error;

  bool _loading = false;
  bool get loading => _loading;

  // ── Lifecycle ──────────────────────────────────────────────────────────

  /// Call once from `main()` **before** `runApp`.
  Future<void> init() async {
    debugPrint('[AuthService] initializing Privy');
    debugPrint('  appId=${AppConfig.privyAppId}');
    debugPrint(
      '  appClientId=${AppConfig.privyClientId.isEmpty ? "<empty>" : AppConfig.privyClientId}',
    );

    // ── Validate configuration (works in both debug AND release) ─────────
    if (AppConfig.privyAppId.isEmpty) {
      _error = 'Privy app ID is not configured. Update AppConfig.';
      _status = AuthStatus.unauthenticated;
      notifyListeners();
      return;
    }
    if (AppConfig.privyClientId.isEmpty) {
      _error = 'Privy client ID is not configured. Update AppConfig.';
      _status = AuthStatus.unauthenticated;
      notifyListeners();
      return;
    }

    // ── Create SDK instance ─────────────────────────────────────────────
    try {
      _privy = Privy.init(
        config: PrivyConfig(
          appId: AppConfig.privyAppId,
          appClientId: AppConfig.privyClientId,
        ),
      );
    } catch (e) {
      debugPrint('[AuthService] Privy.init failed: $e');
      _error = 'Failed to initialise authentication: $e';
      _status = AuthStatus.unauthenticated;
      notifyListeners();
      return;
    }

    // ── Check for existing session ──────────────────────────────────────
    try {
      await _sdk.getAuthState();

      final existingUser = await _sdk.getUser();
      if (existingUser != null) {
        await _handleAuthenticated(existingUser);
      } else {
        _status = AuthStatus.unauthenticated;
      }
    } catch (e) {
      debugPrint('[AuthService] post-init error: $e');
      _status = AuthStatus.unauthenticated;
    }

    notifyListeners();
  }

  // ── Email OTP flow ─────────────────────────────────────────────────────

  /// Step 1 – send a one-time code to [emailAddress].
  Future<void> sendOtp(String emailAddress) async {
    _setLoading();
    try {
      final result = await _sdk.email.sendCode(emailAddress);
      result.fold(
        onSuccess: (_) {
          _email = emailAddress;
          _status = AuthStatus.awaitingOtp;
        },
        onFailure: (error) {
          _error = 'Failed to send OTP: $error';
        },
      );
    } catch (e) {
      _error = 'Failed to send OTP: $e';
    }
    _loading = false;
    notifyListeners();
  }

  /// Step 2 – verify the OTP code the user received.
  Future<void> verifyOtp(String code) async {
    if (_email == null) {
      _error = 'No email address. Please request a new OTP.';
      notifyListeners();
      return;
    }
    _setLoading();
    try {
      final result = await _sdk.email.loginWithCode(email: _email!, code: code);
      PrivyUser? loggedInUser;
      result.fold(
        onSuccess: (user) => loggedInUser = user,
        onFailure: (error) {
          _error = 'OTP verification failed: $error';
        },
      );
      if (loggedInUser != null) {
        await _handleAuthenticated(loggedInUser!);
      }
    } catch (e) {
      _error = 'OTP verification failed: $e';
    }
    _loading = false;
    notifyListeners();
  }

  // ── Social / OAuth login ───────────────────────────────────────────────

  Future<void> loginWithGoogle() => _oauthLogin(OAuthProvider.google);
  Future<void> loginWithApple() => _oauthLogin(OAuthProvider.apple);

  Future<void> _oauthLogin(OAuthProvider provider) async {
    _setLoading();
    try {
      final result = await _sdk.oAuth.login(
        provider: provider,
        appUrlScheme: AppConfig.oauthScheme,
      );
      PrivyUser? loggedInUser;
      result.fold(
        onSuccess: (user) => loggedInUser = user,
        onFailure: (error) {
          _error = 'Social login failed: $error';
        },
      );
      if (loggedInUser != null) {
        await _handleAuthenticated(loggedInUser!);
      }
    } catch (e) {
      _error = 'Social login failed: $e';
    }
    _loading = false;
    notifyListeners();
  }

  // ── Token ──────────────────────────────────────────────────────────────

  /// Returns a fresh Privy JWT access-token, or `null` if not authenticated.
  Future<String?> getAccessToken() async {
    try {
      final user = await _sdk.getUser();
      if (user == null) return null;
      final result = await user.getAccessToken();
      String? token;
      result.fold(onSuccess: (t) => token = t, onFailure: (_) {});
      return token;
    } catch (_) {
      return null;
    }
  }

  // ── Logout ─────────────────────────────────────────────────────────────

  Future<void> logout() async {
    try {
      if (_privy != null) await _privy!.logout();
    } catch (_) {}
    _api.clearTokenProvider();
    _user = null;
    _walletAddress = null;
    _privyUserId = null;
    _email = null;
    _status = AuthStatus.unauthenticated;
    notifyListeners();
  }

  // ── Helpers ────────────────────────────────────────────────────────────

  void clearError() {
    _error = null;
    notifyListeners();
  }

  void _setLoading() {
    _loading = true;
    _error = null;
    notifyListeners();
  }

  /// Called after any successful login path.
  Future<void> _handleAuthenticated(PrivyUser privyUser) async {
    _user = privyUser;
    _privyUserId = privyUser.id;
    _status = AuthStatus.authenticated;

    // Surface the embedded-wallet address (create one if needed).
    try {
      final wallets = privyUser.embeddedEthereumWallets;
      if (wallets.isNotEmpty) {
        _walletAddress = wallets.first.address;
      } else {
        // Create an embedded wallet for the user on first login.
        final result = await privyUser.createEthereumWallet();
        result.fold(
          onSuccess: (wallet) => _walletAddress = wallet.address,
          onFailure: (error) {
            debugPrint('[AuthService] wallet create error: $error');
          },
        );
      }
    } catch (e) {
      debugPrint('[AuthService] embedded wallet error: $e');
    }

    // Wire the auth token into the ApiClient so all subsequent calls are
    // authenticated.
    _api.setTokenProvider(getAccessToken);

    // Auto-link the Privy identity to the on-chain DID (fire-and-forget).
    _tryLinkPrivyToDid();
  }

  /// Best-effort: ensure the user's DID exists and their Privy credential
  /// hash is linked on-chain.
  Future<void> _tryLinkPrivyToDid() async {
    if (_walletAddress == null || _privyUserId == null) return;
    try {
      // SHA-256 of the Privy user id – matches what the contract stores.
      final hash = '0x${sha256.convert(utf8.encode(_privyUserId!)).toString()}';

      // 1. Ensure a DID exists for the wallet address.
      try {
        final existing = await _api.lookupDID(_walletAddress!);
        if (existing['exists'] != true) {
          await _api.createDID(_walletAddress!);
        }
      } catch (_) {
        // DID lookup may 404 – safe to ignore, createDID will handle it.
        try {
          await _api.createDID(_walletAddress!);
        } catch (_) {}
      }

      // 2. Link the Privy hash.
      await _api.linkPrivy(_walletAddress!, hash);
    } catch (e) {
      debugPrint('[AuthService] DID link failed (non-fatal): $e');
    }
  }
}
