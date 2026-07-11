import 'dart:convert';

import 'package:http/http.dart' as http;

import '../models/pronunciation_response.dart';
import 'pronunciation_services.dart';

class PocketBasePronunciationService implements PronunciationService {
  final String baseUrl;

  PocketBasePronunciationService({required this.baseUrl});

  @override
  Future<PronunciationResponse> analyzeScript(String script) async {
    final uri = Uri.parse("$baseUrl/api/pronounce");

    final response = await http.post(
      uri,
      headers: {"Content-Type": "application/json"},
      body: jsonEncode({"script": script}),
    );

    if (response.statusCode != 200) {
      throw Exception("Failed to analyze script.");
    }

    final json = jsonDecode(response.body);

    return PronunciationResponse.fromJson(json);
  }
}
