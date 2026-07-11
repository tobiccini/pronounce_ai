class Pronunciation {
  final String word;

  final String category;

  final String language;

  final String easyPronunciation;

  final String ipaEnglish;

  final String ipaNative;

  final String meaning;

  final String replacement;

  final List<String> presenterTips;

  final double confidence;

  final bool audioAvailable;

  /// NEW
  final String difficulty;

  /// Character index where this word starts in the script.
  final int startIndex;

  /// Character index where this word ends in the script.
  final int endIndex;

  const Pronunciation({
    required this.word,
    required this.category,
    required this.language,
    required this.easyPronunciation,
    required this.ipaEnglish,
    required this.ipaNative,
    required this.meaning,
    required this.replacement,
    required this.presenterTips,
    required this.confidence,
    required this.audioAvailable,

    required this.difficulty,
    required this.startIndex,
    required this.endIndex,
  });

  factory Pronunciation.fromJson(Map<String, dynamic> json) {
    return Pronunciation(
      word: json["word"] ?? "",

      category: json["category"] ?? "",

      language: json["language"] ?? "",

      easyPronunciation: json["easyPronunciation"] ?? "",

      ipaEnglish: json["ipaEnglish"] ?? "",

      ipaNative: json["ipaNative"] ?? "",

      meaning: json["meaning"] ?? "",

      replacement: json["replacement"] ?? "",

      presenterTips: List<String>.from(json["presenterTips"] ?? []),

      confidence: (json["confidence"] ?? 0).toDouble(),

      audioAvailable: json["audioAvailable"] ?? false,

      difficulty: json["difficulty"] ?? "Easy",

      startIndex: json["startIndex"] ?? 0,

      endIndex: json["endIndex"] ?? 0,
    );
  }

  /// Convenience getters used by the UI

  bool get isEasy => difficulty.toLowerCase() == "easy";

  bool get isMedium => difficulty.toLowerCase() == "medium";

  bool get isHard => difficulty.toLowerCase() == "hard";
}
