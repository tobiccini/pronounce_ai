import 'package:flutter/material.dart';

import '../models/pronunciation.dart';

class PronunciationCard extends StatelessWidget {
  final Pronunciation pronunciation;

  const PronunciationCard({super.key, required this.pronunciation});

  @override
  Widget build(BuildContext context) {
    return Card(
      margin: const EdgeInsets.only(bottom: 16),

      child: Padding(
        padding: const EdgeInsets.all(16),

        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,

          children: [
            Row(
              children: [
                Expanded(
                  child: Text(
                    pronunciation.word,

                    style: const TextStyle(
                      fontSize: 22,

                      fontWeight: FontWeight.bold,
                    ),
                  ),
                ),

                Chip(label: Text(pronunciation.category)),
              ],
            ),

            const SizedBox(height: 10),

            Text("Language: ${pronunciation.language}"),

            const SizedBox(height: 8),

            Text(
              "Easy Pronunciation",
              style: TextStyle(fontWeight: FontWeight.bold),
            ),

            Text(pronunciation.easyPronunciation),

            const SizedBox(height: 12),

            Text(
              "IPA (English)",
              style: TextStyle(fontWeight: FontWeight.bold),
            ),

            Text(pronunciation.ipaEnglish),

            const SizedBox(height: 12),

            Text("IPA (Native)", style: TextStyle(fontWeight: FontWeight.bold)),

            Text(pronunciation.ipaNative),

            const SizedBox(height: 12),

            Text("Meaning", style: TextStyle(fontWeight: FontWeight.bold)),

            Text(pronunciation.meaning),

            const SizedBox(height: 12),

            Text(
              "Safer Replacement",
              style: TextStyle(fontWeight: FontWeight.bold),
            ),

            Text(pronunciation.replacement),

            const SizedBox(height: 12),

            Text(
              "Presenter Tips",
              style: TextStyle(fontWeight: FontWeight.bold),
            ),

            ...pronunciation.presenterTips.map(
              (tip) => Padding(
                padding: const EdgeInsets.only(top: 4),

                child: Row(
                  crossAxisAlignment: CrossAxisAlignment.start,

                  children: [
                    const Text("• "),

                    Expanded(child: Text(tip)),
                  ],
                ),
              ),
            ),

            const SizedBox(height: 12),

            LinearProgressIndicator(value: pronunciation.confidence),

            const SizedBox(height: 6),

            Text(
              "Confidence: ${(pronunciation.confidence * 100).toStringAsFixed(0)}%",
            ),

            const SizedBox(height: 12),

            Row(
              children: [
                Icon(
                  pronunciation.audioAvailable
                      ? Icons.volume_up
                      : Icons.volume_off,
                ),

                const SizedBox(width: 8),

                Text(
                  pronunciation.audioAvailable ? "Audio Available" : "No Audio",
                ),
              ],
            ),
          ],
        ),
      ),
    );
  }
}
