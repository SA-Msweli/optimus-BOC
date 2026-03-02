/// Application-wide configuration.
///
/// All values point at the **real** deployed infrastructure – no mocks.
/// Change [backendUrl] when pointing at a different environment.
class AppConfig {
  AppConfig._();

  /// Backend API base URL (EC2 instance running the Optimus protocol server).
  static const String backendUrl = 'http://13.60.166.148';

  /// Privy application ID (the value shown in the Privy dashboard).
  static const String privyAppId = 'cmiyiw1s40005jx0dho2x9bn4';

  /// Privy **mobile** client ID (required by the Flutter SDK).
  ///
  /// This is different from the app ID and the web client ID.
  /// Obtained from the Privy dashboard under API Keys → Client IDs.
  static const String privyClientId =
      'client-WY6TQxeEvQSz7ybBuvhBfaDPMrZZSPPQgK29q9PjUy1vz';

  /// Custom URL scheme used for Privy OAuth redirects.
  /// Must match the scheme registered in AndroidManifest.xml / Info.plist.
  static const String oauthScheme = 'optimus-prime';

  /// Ethereum Sepolia JSON-RPC endpoint.
  static const String rpcUrl =
      'https://sepolia.infura.io/v3/88588baac9d8467f87fd40c7ef3905d7';

  /// On-chain contract addresses (Sepolia).
  static const String bnplManagerAddress =
      '0x4d99Dc2e504c15496319339E822C4a8EAfe3e2ba';
  static const String loanManagerAddress =
      '0xbB0D4067488edf4a007822407e2486412dC8D39D';
  static const String daoManagerAddress =
      '0x561289A9B8439E3fb288a33b3c39C78E0923Cd2b';
  static const String didRegistryAddress =
      '0x0E9D8959bCD99e7AFD7C693e51781058A998b756';
  static const String tokenVaultAddress =
      '0x4C704D51fc47cfe582F8c5477de3AE398B344907';

  /// Chain ID for Sepolia.
  static const int chainId = 11155111;
}
