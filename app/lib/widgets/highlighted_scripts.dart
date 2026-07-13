import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

import '../models/pronunciation.dart';
import '../theme/app_theme.dart';
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

    final onSurface = Theme.of(context).colorScheme.onSurface;

    if (widget.pronunciations.isEmpty) {
      return SelectableText(
        widget.script,
        style: TextStyle(fontSize: 18, height: 1.8, color: onSurface),
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
            style: TextStyle(color: onSurface),
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
                  color: AppColors.highlightBg(context, item.difficulty),

                  borderRadius: BorderRadius.circular(8),

                  border: Border.all(
                    color: AppColors.highlightFg(
                      context,
                      item.difficulty,
                    ).withValues(alpha: .25),
                  ),

                  boxShadow: selectedIndex == i
                      ? [
                          BoxShadow(
                            color: AppColors.highlightBg(
                              context,
                              item.difficulty,
                            ),
                            blurRadius: 20,
                            spreadRadius: 3,
                            offset: const Offset(0, 6),
                          ),
                        ]
                      : [
                          BoxShadow(
                            color: Theme.of(
                              context,
                            ).colorScheme.shadow.withValues(alpha: .08),
                            blurRadius: 4,
                          ),
                        ],
                ),

                child: Text(
                  widget.script.substring(item.startIndex, item.endIndex),

                  style: TextStyle(
                    color: AppColors.highlightFg(context, item.difficulty),
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
          style: TextStyle(color: onSurface),
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
            style: TextStyle(fontSize: 18, height: 2, color: onSurface),
            children: spans,
          ),
        ),
      ),
    );
  }
}
