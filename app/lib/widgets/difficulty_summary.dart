import 'package:flutter/material.dart';
import 'package:pronounce_ai/models/pronunciation_response.dart';

class DifficultySummary extends StatelessWidget {
  final PronunciationResponse response;

  const DifficultySummary({super.key, required this.response});

  @override
  Widget build(BuildContext context) {
    // final easy = pronunciations.where((e) => e.isEasy).length;
    // final medium = pronunciations.where((e) => e.isMedium).length;
    // final hard = pronunciations.where((e) => e.isHard).length;

    final easy = response.words.where((e) => e.isEasy).length;
    final medium = response.words.where((e) => e.isMedium).length;
    final hard = response.words.where((e) => e.isHard).length;

    return Card(
      elevation: 3,
      child: Padding(
        padding: const EdgeInsets.all(18),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Row(
              children: [
                Expanded(
                  child: Text(
                    "Found ${response.words.length} pronunciation ${response.words.length == 1 ? "challenge" : "challenges"}",
                    style: const TextStyle(
                      fontSize: 18,
                      fontWeight: FontWeight.bold,
                    ),
                  ),
                ),

                const SizedBox(height: 14),

                Row(
                  children: [
                    Chip(
                      avatar: Icon(
                        response.fromCache
                            ? Icons.offline_bolt
                            : Icons.auto_awesome,
                        size: 18,
                      ),

                      label: Text(response.fromCache ? "Cache" : "AI"),
                    ),

                    const SizedBox(width: 8),

                    Chip(
                      avatar: const Icon(Icons.timer, size: 18),

                      label: Text("${response.processingTimeMs} ms"),
                    ),
                  ],
                ),
              ],
            ),

            const SizedBox(height: 20),

            Row(
              children: [
                Expanded(
                  child: _badge(
                    color: Colors.green,
                    title: "Easy",
                    count: easy,
                  ),
                ),

                const SizedBox(width: 12),

                Expanded(
                  child: _badge(
                    color: Colors.orange,
                    title: "Medium",
                    count: medium,
                  ),
                ),

                const SizedBox(width: 12),

                Expanded(
                  child: _badge(color: Colors.red, title: "Hard", count: hard),
                ),
              ],
            ),

            const SizedBox(height: 16),

            Align(
              alignment: Alignment.centerRight,
              child: Text(
                "${response.provider} • v${response.version}",
                style: TextStyle(color: Colors.grey.shade600, fontSize: 13),
              ),
            ),
          ],
        ),
      ),
    );
  }

  Widget _badge({
    required Color color,
    required String title,
    required int count,
  }) {
    return Container(
      padding: const EdgeInsets.symmetric(vertical: 14),
      decoration: BoxDecoration(
        color: color.withValues(alpha: .12),
        borderRadius: BorderRadius.circular(14),
      ),
      child: Column(
        children: [
          Icon(Icons.circle, color: color, size: 16),

          const SizedBox(height: 8),

          Text(
            title,
            style: TextStyle(color: color, fontWeight: FontWeight.bold),
          ),

          const SizedBox(height: 6),

          Text(
            "$count",
            style: const TextStyle(fontSize: 24, fontWeight: FontWeight.bold),
          ),
        ],
      ),
    );
  }
}
