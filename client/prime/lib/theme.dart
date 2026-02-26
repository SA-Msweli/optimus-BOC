import 'package:flutter/material.dart';

/// Optimus protocol brand palette and theme data.
class AppTheme {
  AppTheme._();

  // ─── Brand colours ─────────────────────────────────────────────────────

  static const Color primary = Color(0xFF1A73E8); // vibrant blue
  static const Color secondary = Color(0xFF34A853); // green – success / money
  static const Color error = Color(0xFFEA4335); // red – errors / defaults
  static const Color warning = Color(0xFFFBBC04); // amber – late fees / warnings
  static const Color surface = Color(0xFFF8F9FA); // near-white background
  static const Color onSurface = Color(0xFF202124); // dark text
  static const Color cardColor = Colors.white;

  // ─── Text styles ───────────────────────────────────────────────────────

  static const TextStyle heading = TextStyle(
    fontSize: 22,
    fontWeight: FontWeight.w700,
    color: onSurface,
  );

  static const TextStyle subheading = TextStyle(
    fontSize: 16,
    fontWeight: FontWeight.w600,
    color: onSurface,
  );

  static const TextStyle body = TextStyle(
    fontSize: 14,
    fontWeight: FontWeight.w400,
    color: onSurface,
  );

  static const TextStyle caption = TextStyle(
    fontSize: 12,
    fontWeight: FontWeight.w400,
    color: Color(0xFF5F6368),
  );

  static const TextStyle monoValue = TextStyle(
    fontSize: 14,
    fontWeight: FontWeight.w500,
    fontFamily: 'monospace',
    color: onSurface,
  );

  // ─── Widget helpers ────────────────────────────────────────────────────

  static BoxDecoration cardDecoration = BoxDecoration(
    color: cardColor,
    borderRadius: BorderRadius.circular(12),
    boxShadow: const [
      BoxShadow(color: Color(0x0F000000), blurRadius: 8, offset: Offset(0, 2)),
    ],
  );

  static InputDecoration inputDecoration(String label, {String? hint}) =>
      InputDecoration(
        labelText: label,
        hintText: hint,
        border: OutlineInputBorder(borderRadius: BorderRadius.circular(8)),
        contentPadding:
            const EdgeInsets.symmetric(horizontal: 14, vertical: 12),
      );

  // ─── ThemeData ─────────────────────────────────────────────────────────

  static ThemeData get light => ThemeData(
        useMaterial3: true,
        colorSchemeSeed: primary,
        brightness: Brightness.light,
        scaffoldBackgroundColor: surface,
        cardColor: cardColor,
        appBarTheme: const AppBarTheme(
          backgroundColor: primary,
          foregroundColor: Colors.white,
          elevation: 0,
        ),
        elevatedButtonTheme: ElevatedButtonThemeData(
          style: ElevatedButton.styleFrom(
            backgroundColor: primary,
            foregroundColor: Colors.white,
            shape:
                RoundedRectangleBorder(borderRadius: BorderRadius.circular(8)),
            padding: const EdgeInsets.symmetric(horizontal: 24, vertical: 14),
          ),
        ),
        inputDecorationTheme: InputDecorationTheme(
          border: OutlineInputBorder(borderRadius: BorderRadius.circular(8)),
          contentPadding:
              const EdgeInsets.symmetric(horizontal: 14, vertical: 12),
        ),
        chipTheme: ChipThemeData(
          backgroundColor: surface,
          shape:
              RoundedRectangleBorder(borderRadius: BorderRadius.circular(20)),
        ),
      );
}
