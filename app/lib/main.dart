import 'package:flutter/material.dart';
import 'package:pronounce_ai/repository/pronunciation_repository.dart';

import 'screens/home_page.dart';
import 'services/pocketbase_pronunciation_service.dart';
import 'theme/app_theme.dart';

void main() {
  final service = PocketBasePronunciationService(
    baseUrl: "http://127.0.0.1:8090", // for web browsers
    // baseUrl: "http://YOUR-IP:8090", // for physical devices
    //baseUrl: "http://10.0.2.2:8090", // for Android emulator
  );

  final repository = PronunciationRepository(service);

  runApp(PronounceAI(repository));
}

class PronounceAI extends StatefulWidget {
  final PronunciationRepository repository;

  const PronounceAI(this.repository, {super.key});

  @override
  State<PronounceAI> createState() => _PronounceAIState();
}

class _PronounceAIState extends State<PronounceAI> {
  // Defaults to following the device setting; the app bar toggle lets the
  // user force light or dark mode for this session.
  ThemeMode _themeMode = ThemeMode.system;

  void _toggleTheme() {
    setState(() {
      final isDark =
          _themeMode == ThemeMode.dark ||
          (_themeMode == ThemeMode.system &&
              WidgetsBinding.instance.platformDispatcher.platformBrightness ==
                  Brightness.dark);

      _themeMode = isDark ? ThemeMode.light : ThemeMode.dark;
    });
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,

      title: "AMD Pronunciation",

      theme: AppTheme.light,
      darkTheme: AppTheme.dark,
      themeMode: _themeMode,

      home: HomePage(
        repository: widget.repository,
        themeMode: _themeMode,
        onToggleTheme: _toggleTheme,
      ),
    );
  }
}
