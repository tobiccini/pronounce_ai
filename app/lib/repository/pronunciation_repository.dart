import 'package:pronounce_ai/models/pronunciation_response.dart';
import 'package:pronounce_ai/services/pronunciation_services.dart';

class PronunciationRepository {
  final PronunciationService service;

  PronunciationRepository(this.service);

  Future<PronunciationResponse> analyze(String script) {
    return service.analyzeScript(script);
  }
}
