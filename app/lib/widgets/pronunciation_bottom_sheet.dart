import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

import '../models/pronunciation.dart';
import '../theme/app_theme.dart';

class PronunciationBottomSheet extends StatelessWidget {
  final Pronunciation pronunciation;

  const PronunciationBottomSheet({super.key, required this.pronunciation});

  @override
  Widget build(BuildContext context) {
    return DraggableScrollableSheet(
      expand: false,
      initialChildSize: .82,
      maxChildSize: .95,
      minChildSize: .55,
      builder: (_, controller) {
        final scheme = Theme.of(context).colorScheme;

        return Container(
          decoration: BoxDecoration(
            color: scheme.surfaceContainerLow,
            borderRadius: const BorderRadius.vertical(
              top: Radius.circular(30),
            ),
          ),
          clipBehavior: Clip.antiAlias,
          child: ListView(
            controller: controller,
            // padding: const EdgeInsets.all(24),
            padding: const EdgeInsets.fromLTRB(24, 20, 24, 40),

            children: [
              Center(
                child: Container(
                  // width: 60,
                  // height: 6,
                  width: 72,
                  height: 7,
                  decoration: BoxDecoration(
                    color: scheme.outlineVariant,
                    borderRadius: BorderRadius.circular(20),
                  ),
                ),
              ),

              const SizedBox(height: 24),

              SelectableText(
                pronunciation.word,

                // style: const TextStyle(
                //   fontSize: 28,
                //   fontWeight: FontWeight.bold,
                // ),
                style: const TextStyle(
                  fontSize: 30,
                  fontWeight: FontWeight.w800,
                  letterSpacing: -.5,
                ),
              ),

              const SizedBox(height: 8),

              Text(
                pronunciation.replacement,
                style: TextStyle(fontSize: 16, color: scheme.onSurfaceVariant),
              ),

              const SizedBox(height: 10),

              Wrap(
                spacing: 10,
                runSpacing: 10,
                children: [
                  Chip(
                    label: Text(pronunciation.category),
                    visualDensity: VisualDensity.compact,
                  ),
                  Chip(
                    label: Text(pronunciation.language),
                    visualDensity: VisualDensity.compact,
                  ),
                  Chip(
                    surfaceTintColor: Colors.transparent,
                    backgroundColor: AppColors.highlightBg(
                      context,
                      pronunciation.difficulty,
                    ),
                    label: Text(
                      pronunciation.difficulty,
                      style: TextStyle(
                        color: AppColors.highlightFg(
                          context,
                          pronunciation.difficulty,
                        ),
                      ),
                    ),
                    visualDensity: VisualDensity.compact,
                  ),
                ],
              ),

              const SizedBox(height: 24),

              _Section(
                title: "Easy Pronunciation",
                child: SelectableText(
                  pronunciation.easyPronunciation,
                  style: const TextStyle(
                    fontSize: 22,
                    fontWeight: FontWeight.w600,
                  ),
                ),
              ),

              _Section(
                title: "English IPA",
                child: SelectableText(
                  pronunciation.ipaEnglish,
                  style: const TextStyle(
                    fontSize: 18,
                    fontStyle: FontStyle.italic,
                  ),
                ),
              ),

              _Section(
                title: "Native IPA",
                child: SelectableText(
                  pronunciation.ipaNative,
                  style: const TextStyle(
                    fontSize: 18,
                    fontStyle: FontStyle.italic,
                  ),
                ),
              ),

              _Section(
                title: "Meaning",
                child: Text(
                  pronunciation.meaning,
                  style: const TextStyle(fontSize: 17, height: 1.6),
                ),
              ),

              _Section(
                title: "Suggested Replacement",
                child: Container(
                  width: double.infinity,
                  padding: const EdgeInsets.all(14),
                  decoration: BoxDecoration(
                    color: scheme.primaryContainer,
                    borderRadius: BorderRadius.circular(14),
                  ),
                  child: Text(
                    pronunciation.replacement,
                    style: TextStyle(
                      fontSize: 17,
                      color: scheme.onPrimaryContainer,
                    ),
                  ),
                ),
              ),

              _Section(
                title: "Presenter Tips",
                child: Column(
                  children: pronunciation.presenterTips
                      .map(
                        (tip) => Padding(
                          padding: const EdgeInsets.only(bottom: 10),
                          child: Row(
                            crossAxisAlignment: CrossAxisAlignment.start,
                            children: [
                              const Padding(
                                padding: EdgeInsets.only(top: 2),
                                child: Icon(
                                  Icons.lightbulb,
                                  size: 18,
                                  color: Colors.orange,
                                ),
                              ),
                              const SizedBox(width: 10),
                              Expanded(
                                child: Text(
                                  tip,
                                  style: const TextStyle(
                                    fontSize: 16,
                                    height: 1.5,
                                  ),
                                ),
                              ),
                            ],
                          ),
                        ),
                      )
                      .toList(),
                ),
              ),

              const SizedBox(height: 10),

              Container(
                padding: const EdgeInsets.all(18),
                decoration: BoxDecoration(
                  color: scheme.surfaceContainerHighest,
                  borderRadius: BorderRadius.circular(16),
                ),
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Row(
                      children: [
                        const Icon(Icons.verified),

                        const SizedBox(width: 8),

                        const Text(
                          "AI Confidence",
                          style: TextStyle(fontWeight: FontWeight.bold),
                        ),

                        const Spacer(),

                        Text(
                          "${(pronunciation.confidence * 100).toStringAsFixed(0)}%",
                          style: const TextStyle(fontWeight: FontWeight.bold),
                        ),
                      ],
                    ),

                    const SizedBox(height: 12),

                    ClipRRect(
                      borderRadius: BorderRadius.circular(20),
                      child: LinearProgressIndicator(
                        minHeight: 10,
                        value: pronunciation.confidence,
                      ),
                    ),
                  ],
                ),
              ),

              const SizedBox(height: 30),

              SizedBox(
                height: 54,
                width: double.infinity,
                child: FilledButton.icon(
                  onPressed: pronunciation.audioAvailable
                      ? () {
                          // Coming soon
                        }
                      : null,
                  icon: const Icon(Icons.volume_up),
                  label: Text(
                    pronunciation.audioAvailable
                        ? "Play Pronunciation"
                        : "Audio Not Available",
                  ),

                  style: FilledButton.styleFrom(
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(16),
                    ),
                  ),
                ),
              ),

              const SizedBox(height: 14),

              OutlinedButton.icon(
                onPressed: () {
                  Clipboard.setData(
                    ClipboardData(text: pronunciation.easyPronunciation),
                  );

                  ScaffoldMessenger.of(context).showSnackBar(
                    const SnackBar(content: Text("Pronunciation copied")),
                  );
                },

                icon: const Icon(Icons.copy),

                label: const Text("Copy Pronunciation"),
              ),

              const SizedBox(height: 30),
            ],
          ),
        );
      },
    );
  }

}

class _Section extends StatelessWidget {
  final String title;
  final Widget child;

  const _Section({required this.title, required this.child});

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.only(bottom: 22),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(
            title,
            style: TextStyle(
              color: Theme.of(context).colorScheme.onSurfaceVariant,
              fontWeight: FontWeight.bold,
              fontSize: 15,
            ),
          ),
          const SizedBox(height: 8),
          child,
        ],
      ),
    );
  }
}
