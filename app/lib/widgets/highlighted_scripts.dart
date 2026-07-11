import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

import '../models/pronunciation.dart';
import 'pronunciation_bottom_sheet.dart';

class HighlightedScript extends StatefulWidget {
  final String script;
  final List<Pronunciation> pronunciations;

  const HighlightedScript({
    super.key,
    required this.script,
    required this.pronunciations,
  });

  @override
  State<HighlightedScript> createState() => _HighlightedScriptState();
}

class _HighlightedScriptState extends State<HighlightedScript> {
  // int? selectedIndex;

  int? selectedIndex;

  final ScrollController _scrollController = ScrollController();

  @override
  Widget build(BuildContext context) {
    if (widget.script.isEmpty) {
      return const SizedBox();
    }

    if (widget.pronunciations.isEmpty) {
      return SelectableText(
        widget.script,
        style: const TextStyle(fontSize: 18, height: 1.8),
      );
    }

    final words = List<Pronunciation>.from(widget.pronunciations)
      ..sort((a, b) => a.startIndex.compareTo(b.startIndex));

    List<InlineSpan> spans = [];

    int current = 0;

    for (int i = 0; i < words.length; i++) {
      final item = words[i];

      if (item.startIndex > current) {
        spans.add(
          TextSpan(
            text: widget.script.substring(current, item.startIndex),
            style: const TextStyle(color: Colors.black87),
          ),
        );
      }

      spans.add(
        WidgetSpan(
          alignment: PlaceholderAlignment.middle,

          child: GestureDetector(
            onTap: () async {
              HapticFeedback.lightImpact();

              setState(() {
                selectedIndex = i;
              });

              await Future.delayed(const Duration(milliseconds: 180));

              if (!mounted) return;

              await showModalBottomSheet(
                context: context,
                isScrollControlled: true,
                useSafeArea: true,
                backgroundColor: Colors.transparent,
                builder: (_) => PronunciationBottomSheet(pronunciation: item),
              );

              if (!mounted) return;

              setState(() {
                selectedIndex = null;
              });
            },

            child: AnimatedScale(
              duration: const Duration(milliseconds: 220),
              curve: Curves.easeOutBack,
              scale: selectedIndex == i ? 1.22 : 1,

              child: AnimatedContainer(
                duration: const Duration(milliseconds: 220),
                curve: Curves.easeOut,

                padding: const EdgeInsets.symmetric(horizontal: 6, vertical: 4),

                decoration: BoxDecoration(
                  color: _backgroundColor(item.difficulty),

                  borderRadius: BorderRadius.circular(8),

                  border: Border.all(
                    color: _textColor(item.difficulty).withValues(alpha: .25),
                  ),

                  boxShadow: selectedIndex == i
                      ? [
                          BoxShadow(
                            color: _backgroundColor(item.difficulty),
                            blurRadius: 20,
                            spreadRadius: 3,
                            offset: const Offset(0, 6),
                          ),
                        ]
                      : [
                          BoxShadow(
                            color: Colors.black.withValues(alpha: .05),
                            blurRadius: 4,
                          ),
                        ],
                ),

                child: Text(
                  widget.script.substring(item.startIndex, item.endIndex),

                  style: TextStyle(
                    color: _textColor(item.difficulty),
                    fontWeight: FontWeight.w700,
                    fontSize: 18,
                    letterSpacing: .2,
                  ),
                ),
              ),
            ),
          ),
        ),
      );

      current = item.endIndex;
    }

    if (current < widget.script.length) {
      spans.add(
        TextSpan(
          text: widget.script.substring(current),
          style: const TextStyle(color: Colors.black87),
        ),
      );
    }

    return Card(
      elevation: 4,
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(18)),
      child: SingleChildScrollView(
        controller: _scrollController,
        padding: const EdgeInsets.all(22),
        child: RichText(
          text: TextSpan(
            style: const TextStyle(
              fontSize: 18,
              height: 2,
              color: Colors.black87,
            ),
            children: spans,
          ),
        ),
      ),
    );
  }

  Color _backgroundColor(String difficulty) {
    switch (difficulty.toLowerCase()) {
      case "easy":
        return const Color(0xffDCFCE7);

      case "medium":
        return const Color(0xffFEF3C7);

      case "hard":
        return const Color(0xffFEE2E2);

      default:
        return Colors.blue.shade100;
    }
  }

  Color _textColor(String difficulty) {
    switch (difficulty.toLowerCase()) {
      case "easy":
        return const Color(0xff166534);

      case "medium":
        return const Color(0xff92400E);

      case "hard":
        return const Color(0xff991B1B);

      default:
        return Colors.blue.shade900;
    }
  }
}
