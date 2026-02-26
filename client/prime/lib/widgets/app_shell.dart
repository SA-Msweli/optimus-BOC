import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import '../services/auth_service.dart';
import '../screens/did_screen.dart';
import '../screens/bnpl_screen.dart';
import '../screens/loan_screen.dart';
import '../screens/dao_screen.dart';
import '../screens/vault_screen.dart';

/// Truncates a hex address to 0x1234…abcd format.
String _short(String addr) {
  if (addr.length <= 12) return addr;
  return '${addr.substring(0, 6)}…${addr.substring(addr.length - 4)}';
}

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
    final auth = context.watch<AuthService>();

    return Scaffold(
      appBar: AppBar(
        title: Text(
          auth.walletAddress != null
              ? _short(auth.walletAddress!)
              : auth.email ?? 'Optimus',
        ),
        actions: [
          if (auth.walletAddress != null)
            Padding(
              padding: const EdgeInsets.only(right: 8),
              child: Chip(
                avatar: const Icon(
                  Icons.account_balance_wallet,
                  size: 16,
                  color: Colors.white70,
                ),
                label: Text(
                  _short(auth.walletAddress!),
                  style: const TextStyle(color: Colors.white, fontSize: 12),
                ),
                backgroundColor: Colors.white24,
                side: BorderSide.none,
              ),
            ),
          IconButton(
            icon: const Icon(Icons.logout),
            tooltip: 'Sign out',
            onPressed: () async {
              final confirmed = await showDialog<bool>(
                context: context,
                builder: (ctx) => AlertDialog(
                  title: const Text('Sign out?'),
                  content: const Text(
                    'You will need to log in again to access the protocol.',
                  ),
                  actions: [
                    TextButton(
                      onPressed: () => Navigator.pop(ctx, false),
                      child: const Text('Cancel'),
                    ),
                    TextButton(
                      onPressed: () => Navigator.pop(ctx, true),
                      child: const Text('Sign out'),
                    ),
                  ],
                ),
              );
              if (confirmed == true && context.mounted) {
                auth.logout();
              }
            },
          ),
        ],
      ),
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
            selectedIcon: Icon(Icons.shopping_cart, color: Color(0xFF1A73E8)),
            label: 'BNPL',
          ),
          NavigationDestination(
            icon: Icon(Icons.account_balance_outlined),
            selectedIcon: Icon(Icons.account_balance, color: Color(0xFF1A73E8)),
            label: 'Loans',
          ),
          NavigationDestination(
            icon: Icon(Icons.how_to_vote_outlined),
            selectedIcon: Icon(Icons.how_to_vote, color: Color(0xFF1A73E8)),
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
