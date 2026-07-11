import 'package:flutter/material.dart';
import 'package:pronounce_ai/repository/pronunciation_repository.dart';

import 'screens/home_page.dart';
import 'services/pocketbase_pronunciation_service.dart';

void main() {
  final service = PocketBasePronunciationService(
    //baseUrl: "http://127.0.0.1:8090", // for web browsers
    // baseUrl: "http://YOUR-IP:8090", // for physical devices
    baseUrl: "http://10.0.2.2:8090", // for Android emulator
  );

  final repository = PronunciationRepository(service);

  runApp(PronounceAI(repository));
}

class PronounceAI extends StatelessWidget {
  final PronunciationRepository repository;

  const PronounceAI(this.repository, {super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,

      title: "AMD Pronunciation",

      home: HomePage(repository: repository),
    );
  }
}
