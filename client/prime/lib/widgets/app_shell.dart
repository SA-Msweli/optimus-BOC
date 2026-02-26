import 'package:flutter/material.dart';
import '../screens/did_screen.dart';
import '../screens/bnpl_screen.dart';
import '../screens/loan_screen.dart';
import '../screens/dao_screen.dart';
import '../screens/vault_screen.dart';

/// Main navigation shell with a bottom navigation bar.
class AppShell extends StatefulWidget {
  const AppShell({super.key});

  @override
  State<AppShell> createState() => _AppShellState();
}

class _AppShellState extends State<AppShell> {
  int _index = 0;

  static const _pages = <Widget>[
    DIDScreen(),
    BNPLScreen(),
    LoanScreen(),
    DAOScreen(),
    VaultScreen(),
  ];

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: IndexedStack(index: _index, children: _pages),
      bottomNavigationBar: NavigationBar(
        selectedIndex: _index,
        onDestinationSelected: (i) => setState(() => _index = i),
        destinations: const [
          NavigationDestination(
            icon: Icon(Icons.fingerprint),
            selectedIcon: Icon(Icons.fingerprint, color: Color(0xFF1A73E8)),
            label: 'Identity',
          ),
          NavigationDestination(
            icon: Icon(Icons.shopping_cart_outlined),
            selectedIcon:
                Icon(Icons.shopping_cart, color: Color(0xFF1A73E8)),
            label: 'BNPL',
          ),
          NavigationDestination(
            icon: Icon(Icons.account_balance_outlined),
            selectedIcon:
                Icon(Icons.account_balance, color: Color(0xFF1A73E8)),
            label: 'Loans',
          ),
          NavigationDestination(
            icon: Icon(Icons.how_to_vote_outlined),
            selectedIcon:
                Icon(Icons.how_to_vote, color: Color(0xFF1A73E8)),
            label: 'DAO',
          ),
          NavigationDestination(
            icon: Icon(Icons.savings_outlined),
            selectedIcon: Icon(Icons.savings, color: Color(0xFF1A73E8)),
            label: 'Vault',
          ),
        ],
      ),
    );
  }
}
