import 'package:flutter/material.dart';
import 'package:pronounce_ai/repository/pronunciation_repository.dart';
import 'package:pronounce_ai/widgets/difficulty_summary.dart';
import 'package:pronounce_ai/widgets/highlighted_scripts.dart';

import '../models/pronunciation_response.dart';

class HomePage extends StatefulWidget {
  final PronunciationRepository repository;

  const HomePage({super.key, required this.repository});

  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  double _opacity = 0;
  final TextEditingController _controller = TextEditingController();

  PronunciationResponse? _response;
  String _analyzedScript = "";

  bool _loading = false;

  String? _error;

  Future<void> _analyze() async {
    final script = _controller.text.trim();

    if (script.isEmpty) return;

    setState(() {
      _loading = true;
      _error = null;
    });

    try {
      final result = await widget.repository.analyze(script);

      setState(() {
        _response = result;
        _analyzedScript = script;
        _opacity = 0;
      });

      Future.delayed(const Duration(milliseconds: 40), () {
        if (!mounted) return;

        setState(() {
          _opacity = 1;
        });
      });
    } catch (e) {
      setState(() {
        _error = e.toString();
      });
    } finally {
      setState(() {
        _loading = false;
      });
    }

    // _controller.clear();
  }

  @override
  void dispose() {
    _controller.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text("Pronunciation Assistant"),
        centerTitle: true,
        actions: [
          IconButton(
            onPressed: _response == null ? null : _analyze,

            icon: const Icon(Icons.refresh),

            tooltip: "Analyze Again",
          ),
        ],
      ),

      body: SafeArea(
        child: Column(
          children: [
            /// Input Area
            Padding(
              padding: const EdgeInsets.all(16),

              child: Column(
                children: [
                  // TextField(
                  //   controller: _controller,
                  //   maxLines: 8,
                  //   decoration: InputDecoration(
                  //     hintText: "Paste a news script here...",
                  //     border: OutlineInputBorder(
                  //       borderRadius: BorderRadius.circular(12),
                  //     ),
                  //   ),
                  // ),
                  TextField(
                    controller: _controller,
                    onChanged: (_) {
                      if (_response != null) {
                        setState(() {
                          _response = null;
                        });
                      }
                    },
                    maxLines: 8,
                    decoration: InputDecoration(
                      hintText: "Paste a news script here...",
                      border: OutlineInputBorder(
                        borderRadius: BorderRadius.circular(12),
                      ),
                    ),
                  ),

                  const SizedBox(height: 16),

                  SizedBox(
                    width: double.infinity,
                    height: 50,

                    child: FilledButton(
                      onPressed: _loading ? null : _analyze,

                      child: _loading
                          ? const SizedBox(
                              height: 22,
                              width: 22,
                              child: CircularProgressIndicator(strokeWidth: 2),
                            )
                          : const Text("Analyze Script"),
                    ),
                  ),
                ],
              ),
            ),

            if (_error != null)
              Padding(
                padding: const EdgeInsets.symmetric(horizontal: 16),
                child: Text(_error!, style: const TextStyle(color: Colors.red)),
              ),

            const SizedBox(height: 8),

            Expanded(
              child: _response == null
                  ? Center(
                      child: Column(
                        mainAxisAlignment: MainAxisAlignment.center,
                        children: [
                          Icon(
                            Icons.record_voice_over,
                            size: 90,
                            color: Colors.grey.shade400,
                          ),

                          const SizedBox(height: 24),

                          const Text(
                            "Paste a news script",
                            style: TextStyle(
                              fontSize: 24,
                              fontWeight: FontWeight.bold,
                            ),
                          ),

                          const SizedBox(height: 8),

                          Text(
                            "We'll detect difficult names,\nplaces and organizations.",
                            textAlign: TextAlign.center,
                            style: TextStyle(color: Colors.grey.shade600),
                          ),
                        ],
                      ),
                    )
                  : AnimatedOpacity(
                      opacity: _opacity,
                      duration: const Duration(milliseconds: 450),

                      curve: Curves.easeOut,
                      child: RefreshIndicator(
                        onRefresh: () async {
                          if (_controller.text.trim().isEmpty) return;

                          await _analyze();
                        },
                        child: SingleChildScrollView(
                          padding: const EdgeInsets.all(16),

                          child: Column(
                            children: [
                              DifficultySummary(response: _response!),

                              const SizedBox(height: 16),

                              HighlightedScript(
                                script: _analyzedScript,
                                pronunciations: _response!.words,
                              ),

                              const SizedBox(height: 24),

                              Center(
                                child: Container(
                                  padding: const EdgeInsets.symmetric(
                                    horizontal: 18,
                                    vertical: 10,
                                  ),

                                  decoration: BoxDecoration(
                                    color: Colors.grey.shade100,

                                    borderRadius: BorderRadius.circular(30),
                                  ),

                                  child: const Row(
                                    mainAxisSize: MainAxisSize.min,

                                    children: [
                                      Icon(Icons.memory, size: 18),

                                      SizedBox(width: 8),

                                      Text("Powered by AMD + Fireworks"),
                                    ],
                                  ),
                                ),
                              ),

                              const SizedBox(height: 30),
                            ],
                          ),
                        ),
                      ),
                    ),
            ),
          ],
        ),
      ),
    );
  }
}
