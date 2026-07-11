package ai

// ============================================================================
// Gemini API Request Models
// ============================================================================

// GeminiRequest is the request body sent to the Gemini API.
type GeminiRequest struct {
	Contents []GeminiContent `json:"contents"`
}

type GeminiContent struct {
	Parts []GeminiPart `json:"parts"`
}

type GeminiPart struct {
	Text string `json:"text"`
}

// ============================================================================
// Gemini API Response Models
// ============================================================================

// GeminiResponse represents the response returned by Gemini.
type GeminiResponse struct {
	Candidates []GeminiCandidate `json:"candidates"`
}

type GeminiCandidate struct {
	Content GeminiContent `json:"content"`
}

// ============================================================================
// Internal JSON Parsing Models
// ============================================================================

// Gemini returns a JSON object:
//
// {
//   "words": [
//      ...
//   ]
// }
//
// We decode that into this struct before converting it into
// our application response.
type AIPronunciationResponse struct {
	Words []PronunciationResult `json:"words"`
}
