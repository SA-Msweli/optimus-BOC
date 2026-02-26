import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import 'services/api_client.dart';
import 'services/auth_service.dart';
import 'services/did_service.dart';
import 'services/bnpl_service.dart';
import 'services/loan_service.dart';
import 'services/dao_service.dart';
import 'services/vault_service.dart';
import 'screens/login_screen.dart';
import 'theme.dart';
import 'widgets/app_shell.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();

  final api = ApiClient();
  final auth = AuthService(api);

  // Initialise the Privy SDK and check for a persisted session.
  await auth.init();

  runApp(
    MultiProvider(
      providers: [
        ChangeNotifierProvider.value(value: auth),
        ChangeNotifierProvider(create: (_) => DIDService(api)),
        ChangeNotifierProvider(create: (_) => BNPLService(api)),
        ChangeNotifierProvider(create: (_) => LoanService(api)),
        ChangeNotifierProvider(create: (_) => DAOService(api)),
        ChangeNotifierProvider(create: (_) => VaultService(api)),
      ],
      child: const OptimusApp(),
    ),
  );
}

class OptimusApp extends StatelessWidget {
  const OptimusApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Optimus Protocol',
      theme: AppTheme.light,
      debugShowCheckedModeBanner: false,
      home: const _AuthGate(),
    );
  }
}

/// Routes to [LoginScreen] or [AppShell] depending on Privy auth state.
class _AuthGate extends StatelessWidget {
  const _AuthGate();

  @override
  Widget build(BuildContext context) {
    return Consumer<AuthService>(
      builder: (context, auth, _) {
        switch (auth.status) {
          case AuthStatus.unknown:
            // Splash / loading while SDK initialises.
            return const Scaffold(
              body: Center(child: CircularProgressIndicator()),
            );
          case AuthStatus.authenticated:
            return const AppShell();
          case AuthStatus.unauthenticated:
          case AuthStatus.awaitingOtp:
            return const LoginScreen();
        }
      },
    );
  }
}
