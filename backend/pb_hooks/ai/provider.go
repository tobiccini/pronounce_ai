package ai

// PronunciationResult represents a single pronunciation entry
// returned by any AI provider (Gemini today, AMD later).
type PronunciationResult struct {
	// Original detected word or phrase.
	Word string `json:"word"`

	// Person, Place, Organization, Acronym, Medical, etc.
	Category string `json:"category"`

	// Native language of the word.
	Language string `json:"language"`

	// Human-friendly pronunciation.
	EasyPronunciation string `json:"easyPronunciation"`

	// IPA optimized for English speakers.
	IPAEnglish string `json:"ipaEnglish"`

	// Native IPA pronunciation.
	IPANative string `json:"ipaNative"`

	// Brief description or meaning.
	Meaning string `json:"meaning"`

	// Suggested replacement for live broadcasting.
	Replacement string `json:"replacement"`

	// Helpful pronunciation advice.
	PresenterTips []string `json:"presenterTips"`

	// Estimated confidence (0–1).
	Confidence float64 `json:"confidence"`

	// Indicates whether audio pronunciation is available.
	AudioAvailable bool `json:"audioAvailable"`

	// Easy | Medium | Hard
	Difficulty string `json:"difficulty"`

	// Character offsets in the original script.
	StartIndex int `json:"startIndex"`

	EndIndex int `json:"endIndex"`
}

// Response returned to Flutter.
type PronunciationResponse struct {
	Success bool `json:"success"`

	Provider string `json:"provider"`

	Version string `json:"version"`

	Model string `json:"model"`

	ProcessingTimeMs int64 `json:"processingTimeMs"`

	Count int `json:"count"`

	// NEW
	FromCache bool `json:"fromCache"`

	Words []PronunciationResult `json:"words"`
}

// Every AI backend must satisfy this interface.
type AIProvider interface {
	Analyze(script string) (*PronunciationResponse, error)
}
