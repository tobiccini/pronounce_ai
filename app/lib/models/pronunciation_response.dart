import 'pronunciation.dart';

class PronunciationResponse {
  final String script;

  final bool fromCache;

  final List<Pronunciation> words;

  final String provider;

  final String version;

  final int processingTimeMs;

  const PronunciationResponse({
    required this.script,
    required this.fromCache,
    required this.words,

    required this.provider,
    required this.version,
    required this.processingTimeMs,
  });

  factory PronunciationResponse.fromJson(Map<String, dynamic> json) {
    return PronunciationResponse(
      script: json["script"] ?? "",

      words: (json["words"] as List)
          .map((e) => Pronunciation.fromJson(e))
          .toList(),

      fromCache: json["fromCache"] ?? false,

      provider: json["provider"] ?? "",

      version: json["version"] ?? "",

      processingTimeMs: json["processingTimeMs"] ?? 0,
    );
  }
}
