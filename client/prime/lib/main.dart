import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import 'services/api_client.dart';
import 'services/did_service.dart';
import 'services/bnpl_service.dart';
import 'services/loan_service.dart';
import 'services/dao_service.dart';
import 'services/vault_service.dart';
import 'theme.dart';
import 'widgets/app_shell.dart';

void main() {
  final api = ApiClient();

  runApp(
    MultiProvider(
      providers: [
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
      home: const AppShell(),
    );
  }
}

