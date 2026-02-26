import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import '../services/auth_service.dart';
import '../theme.dart';
import '../widgets/shared.dart';

/// Full-screen login / onboarding screen powered by Privy.
///
/// Supports two flows:
/// 1. **Email OTP** – user enters email → receives code → enters code.
/// 2. **Social login** – one-tap Google / Apple sign-in.
class LoginScreen extends StatefulWidget {
  const LoginScreen({super.key});

  @override
  State<LoginScreen> createState() => _LoginScreenState();
}

class _LoginScreenState extends State<LoginScreen> {
  final _emailCtrl = TextEditingController();
  final _otpCtrl = TextEditingController();

  @override
  void dispose() {
    _emailCtrl.dispose();
    _otpCtrl.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SafeArea(
        child: Consumer<AuthService>(
          builder: (context, auth, _) {
            return Center(
              child: SingleChildScrollView(
                padding: const EdgeInsets.symmetric(horizontal: 28),
                child: Column(
                  mainAxisSize: MainAxisSize.min,
                  children: [
                    // ── Brand header ──
                    const Icon(
                      Icons.account_balance_wallet,
                      size: 72,
                      color: AppTheme.primary,
                    ),
                    const SizedBox(height: 16),
                    Text('Optimus Protocol', style: AppTheme.heading),
                    const SizedBox(height: 6),
                    Text(
                      'Decentralised BNPL · Lending · Governance',
                      style: AppTheme.caption,
                      textAlign: TextAlign.center,
                    ),
                    const SizedBox(height: 36),

                    // ── Error banner ──
                    if (auth.error != null)
                      ErrorBanner(
                        message: auth.error!,
                        onDismiss: auth.clearError,
                      ),

                    // ── OTP Step 1: email input ──
                    if (auth.status != AuthStatus.awaitingOtp) ...[
                      InfoCard(
                        title: 'Sign in with Email',
                        child: Column(
                          children: [
                            TextField(
                              controller: _emailCtrl,
                              keyboardType: TextInputType.emailAddress,
                              textInputAction: TextInputAction.go,
                              decoration: AppTheme.inputDecoration(
                                'Email address',
                                hint: 'you@example.com',
                              ),
                              onSubmitted: (_) => _sendOtp(auth),
                            ),
                            const SizedBox(height: 14),
                            SizedBox(
                              width: double.infinity,
                              child: ElevatedButton(
                                onPressed: auth.loading
                                    ? null
                                    : () => _sendOtp(auth),
                                child: auth.loading
                                    ? const SizedBox(
                                        height: 20,
                                        width: 20,
                                        child: CircularProgressIndicator(
                                          strokeWidth: 2,
                                          color: Colors.white,
                                        ),
                                      )
                                    : const Text('Continue'),
                              ),
                            ),
                          ],
                        ),
                      ),
                    ],

                    // ── OTP Step 2: code verification ──
                    if (auth.status == AuthStatus.awaitingOtp) ...[
                      InfoCard(
                        title: 'Enter verification code',
                        child: Column(
                          children: [
                            Text(
                              'A 6-digit code was sent to ${auth.email}',
                              style: AppTheme.body,
                            ),
                            const SizedBox(height: 14),
                            TextField(
                              controller: _otpCtrl,
                              keyboardType: TextInputType.number,
                              textInputAction: TextInputAction.done,
                              decoration: AppTheme.inputDecoration(
                                'Verification code',
                                hint: '123456',
                              ),
                              onSubmitted: (_) => _verifyOtp(auth),
                            ),
                            const SizedBox(height: 14),
                            Row(
                              children: [
                                Expanded(
                                  child: OutlinedButton(
                                    onPressed: auth.loading
                                        ? null
                                        : () {
                                            _otpCtrl.clear();
                                            auth.sendOtp(
                                              _emailCtrl.text.trim(),
                                            );
                                          },
                                    child: const Text('Resend code'),
                                  ),
                                ),
                                const SizedBox(width: 12),
                                Expanded(
                                  child: ElevatedButton(
                                    onPressed: auth.loading
                                        ? null
                                        : () => _verifyOtp(auth),
                                    child: auth.loading
                                        ? const SizedBox(
                                            height: 20,
                                            width: 20,
                                            child: CircularProgressIndicator(
                                              strokeWidth: 2,
                                              color: Colors.white,
                                            ),
                                          )
                                        : const Text('Verify'),
                                  ),
                                ),
                              ],
                            ),
                            const SizedBox(height: 8),
                            TextButton(
                              onPressed: () {
                                _otpCtrl.clear();
                                auth.logout(); // reset to email entry
                              },
                              child: const Text('Use a different email'),
                            ),
                          ],
                        ),
                      ),
                    ],

                    // ── Divider ──
                    if (auth.status != AuthStatus.awaitingOtp) ...[
                      const SizedBox(height: 8),
                      Row(
                        children: [
                          const Expanded(child: Divider()),
                          Padding(
                            padding: const EdgeInsets.symmetric(horizontal: 14),
                            child: Text('or', style: AppTheme.caption),
                          ),
                          const Expanded(child: Divider()),
                        ],
                      ),
                      const SizedBox(height: 8),

                      // ── Social logins ──
                      SizedBox(
                        width: double.infinity,
                        child: OutlinedButton.icon(
                          onPressed: auth.loading ? null : auth.loginWithGoogle,
                          icon: const Icon(Icons.g_mobiledata, size: 24),
                          label: const Text('Continue with Google'),
                          style: OutlinedButton.styleFrom(
                            padding: const EdgeInsets.symmetric(vertical: 14),
                          ),
                        ),
                      ),
                      const SizedBox(height: 10),
                      SizedBox(
                        width: double.infinity,
                        child: OutlinedButton.icon(
                          onPressed: auth.loading ? null : auth.loginWithApple,
                          icon: const Icon(Icons.apple, size: 22),
                          label: const Text('Continue with Apple'),
                          style: OutlinedButton.styleFrom(
                            padding: const EdgeInsets.symmetric(vertical: 14),
                          ),
                        ),
                      ),
                    ],

                    const SizedBox(height: 32),
                    Text(
                      'By continuing you agree to our Terms of Service.',
                      style: AppTheme.caption,
                      textAlign: TextAlign.center,
                    ),
                  ],
                ),
              ),
            );
          },
        ),
      ),
    );
  }

  // ── Actions ────────────────────────────────────────────────────────────

  void _sendOtp(AuthService auth) {
    final email = _emailCtrl.text.trim();
    if (email.isEmpty || !email.contains('@')) return;
    auth.sendOtp(email);
  }

  void _verifyOtp(AuthService auth) {
    final code = _otpCtrl.text.trim();
    if (code.length < 6) return;
    auth.verifyOtp(code);
  }
}
