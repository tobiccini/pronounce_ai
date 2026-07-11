import 'package:pronounce_ai/models/pronunciation_response.dart';

abstract class PronunciationService {
  Future<PronunciationResponse> analyzeScript(String script);
}
