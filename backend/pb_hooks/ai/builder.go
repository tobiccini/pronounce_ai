package ai

import (
	"fmt"
	"strings"
)

func BuildPronunciationPrompt(script string) string {

	// Normalize whitespace.
	script = strings.TrimSpace(script)

	// Normalize line endings.
	script = strings.ReplaceAll(script, "\r\n", "\n")

	// Remove excessive blank lines.
	for strings.Contains(script, "\n\n\n") {
		script = strings.ReplaceAll(script, "\n\n\n", "\n\n")
	}

	return fmt.Sprintf(
		PronunciationPrompt,
		script,
	)
}
