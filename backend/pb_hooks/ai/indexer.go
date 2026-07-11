package ai

import (
	"strings"
	"unicode/utf8"
)

// PopulateIndexes computes Unicode-safe character indexes for every
// detected word or phrase in the original script.
//
// If a word appears multiple times, each occurrence is matched in order.
// If a word cannot be found, StartIndex and EndIndex are set to -1.
func PopulateIndexes(script string, words []PronunciationResult) {

	lowerScript := strings.ToLower(script)

	// Tracks where to continue searching for each unique word.
	searchPositions := make(map[string]int)

	for i := range words {

		word := strings.TrimSpace(words[i].Word)

		if word == "" {
			words[i].StartIndex = -1
			words[i].EndIndex = -1
			continue
		}

		lowerWord := strings.ToLower(word)

		startByte := findWholeWord(
			lowerScript,
			lowerWord,
			searchPositions[lowerWord],
		)

		if startByte == -1 {

			// Fall back to searching from the beginning.
			startByte = findWholeWord(
				lowerScript,
				lowerWord,
				0,
			)
		}

		if startByte == -1 {
			words[i].StartIndex = -1
			words[i].EndIndex = -1
			continue
		}

		endByte := startByte + len(word)

		// Convert byte offsets to Unicode character (rune) offsets.
		words[i].StartIndex = utf8.RuneCountInString(script[:startByte])
		words[i].EndIndex = words[i].StartIndex + utf8.RuneCountInString(word)

		// Continue searching after this occurrence.
		searchPositions[lowerWord] = endByte
	}
}

// findWholeWord searches for a whole word or phrase starting at startByte.
// It avoids matching inside another word.
func findWholeWord(text, target string, startByte int) int {

	search := startByte

	for {

		index := strings.Index(text[search:], target)

		if index == -1 {
			return -1
		}

		index += search

		beforeOK := index == 0 || !isWordChar(rune(text[index-1]))

		after := index + len(target)

		afterOK := after >= len(text) || !isWordChar(rune(text[after]))

		if beforeOK && afterOK {
			return index
		}

		search = index + 1

		if search >= len(text) {
			return -1
		}
	}
}

// isWordChar returns true for ASCII letters, digits and underscore.
func isWordChar(r rune) bool {
	return (r >= 'a' && r <= 'z') ||
		(r >= 'A' && r <= 'Z') ||
		(r >= '0' && r <= '9') ||
		r == '_'
}
