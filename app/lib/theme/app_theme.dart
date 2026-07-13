import 'package:flutter/material.dart';

/// Central place for the app's visual identity.
///
/// [AppTheme.light] / [AppTheme.dark] are wired up in `main.dart` via
/// `MaterialApp(theme: ..., darkTheme: ..., themeMode: ...)`.
///
/// [AppColors] holds the "semantic" colors (easy/medium/hard highlight
/// chips, suggestion boxes, etc.) that aren't part of Flutter's standard
/// ColorScheme but still need a light and a dark variant.
class AppTheme {
  AppTheme._();

  static const _seed = Color(0xFF2563EB); // Tailwind blue-600

  static ThemeData get light => _base(
    ColorScheme.fromSeed(seedColor: _seed, brightness: Brightness.light),
  );

  static ThemeData get dark => _base(
    ColorScheme.fromSeed(seedColor: _seed, brightness: Brightness.dark),
  );

  static ThemeData _base(ColorScheme scheme) {
    return ThemeData(
      useMaterial3: true,
      colorScheme: scheme,
      scaffoldBackgroundColor: scheme.surface,

      appBarTheme: AppBarTheme(
        backgroundColor: scheme.surface,
        foregroundColor: scheme.onSurface,
        centerTitle: true,
        elevation: 0,
        scrolledUnderElevation: 2,
        surfaceTintColor: scheme.surfaceTint,
      ),

      cardTheme: CardThemeData(
        color: scheme.surfaceContainerLow,
        surfaceTintColor: scheme.surfaceTint,
        elevation: scheme.brightness == Brightness.dark ? 0 : 3,
        shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(16)),
      ),

      inputDecorationTheme: InputDecorationTheme(
        filled: true,
        fillColor: scheme.surfaceContainerHighest.withValues(alpha: .4),
        border: OutlineInputBorder(
          borderRadius: BorderRadius.circular(12),
          borderSide: BorderSide(color: scheme.outlineVariant),
        ),
        enabledBorder: OutlineInputBorder(
          borderRadius: BorderRadius.circular(12),
          borderSide: BorderSide(color: scheme.outlineVariant),
        ),
        focusedBorder: OutlineInputBorder(
          borderRadius: BorderRadius.circular(12),
          borderSide: BorderSide(color: scheme.primary, width: 2),
        ),
      ),

      filledButtonTheme: FilledButtonThemeData(
        style: FilledButton.styleFrom(
          shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(12)),
        ),
      ),

      chipTheme: ChipThemeData(
        backgroundColor: scheme.surfaceContainerHighest,
        labelStyle: TextStyle(color: scheme.onSurface),
        side: BorderSide(color: scheme.outlineVariant),
      ),

      bottomSheetTheme: BottomSheetThemeData(
        backgroundColor: scheme.surfaceContainerLow,
        surfaceTintColor: scheme.surfaceTint,
      ),

      dividerTheme: DividerThemeData(color: scheme.outlineVariant),
    );
  }
}

/// Semantic colors that live outside the ColorScheme (difficulty tags,
/// highlight backgrounds, callout boxes) — each with a light and dark value.
class AppColors {
  AppColors._();

  static bool _isDark(BuildContext context) =>
      Theme.of(context).brightness == Brightness.dark;

  // Easy / Medium / Hard accent colors used for badges & icons.
  static Color easy(BuildContext context) =>
      _isDark(context) ? const Color(0xFF4ADE80) : const Color(0xFF16A34A);

  static Color medium(BuildContext context) =>
      _isDark(context) ? const Color(0xFFFBBF24) : const Color(0xFFD97706);

  static Color hard(BuildContext context) =>
      _isDark(context) ? const Color(0xFFF87171) : const Color(0xFFDC2626);

  // Highlighted-word chip backgrounds (script view).
  static Color highlightBg(BuildContext context, String difficulty) {
    final dark = _isDark(context);
    switch (difficulty.toLowerCase()) {
      case "easy":
        return dark ? const Color(0xFF14532D) : const Color(0xFFDCFCE7);
      case "medium":
        return dark ? const Color(0xFF78350F) : const Color(0xFFFEF3C7);
      case "hard":
        return dark ? const Color(0xFF7F1D1D) : const Color(0xFFFEE2E2);
      default:
        return dark ? const Color(0xFF1E3A8A) : const Color(0xFFDBEAFE);
    }
  }

  // Highlighted-word chip text (script view) + difficulty badge chips.
  static Color highlightFg(BuildContext context, String difficulty) {
    final dark = _isDark(context);
    switch (difficulty.toLowerCase()) {
      case "easy":
        return dark ? const Color(0xFFBBF7D0) : const Color(0xFF166534);
      case "medium":
        return dark ? const Color(0xFFFDE68A) : const Color(0xFF92400E);
      case "hard":
        return dark ? const Color(0xFFFECACA) : const Color(0xFF991B1B);
      default:
        return dark ? const Color(0xFFBFDBFE) : const Color(0xFF1E3A8A);
    }
  }
}
