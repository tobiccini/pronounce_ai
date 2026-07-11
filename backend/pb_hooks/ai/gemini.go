package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// GeminiProvider implements the AIProvider interface.
type GeminiProvider struct {
	APIKey string
}

// Analyze sends a news script to Gemini and converts the response
// into the application's standard response format.
func (g *GeminiProvider) Analyze(script string) (*PronunciationResponse, error) {

	start := time.Now()

	// Build the prompt.
	prompt := BuildPronunciationPrompt(script)

	request := GeminiRequest{
		Contents: []GeminiContent{
			{
				Parts: []GeminiPart{
					{
						Text: prompt,
					},
				},
			},
		},
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(
		"https://generativelanguage.googleapis.com/v1beta/models/gemini-2.5-flash:generateContent?key=%s",
		g.APIKey,
	)

	resp, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(body),
	)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"gemini returned status %d: %s",
			resp.StatusCode,
			string(responseBody),
		)
	}

	var geminiResponse GeminiResponse

	if err := json.Unmarshal(responseBody, &geminiResponse); err != nil {
		return nil, err
	}

	if len(geminiResponse.Candidates) == 0 {
		return nil, fmt.Errorf("gemini returned no candidates")
	}

	if len(geminiResponse.Candidates[0].Content.Parts) == 0 {
		return nil, fmt.Errorf("gemini returned no content")
	}

	text := geminiResponse.Candidates[0].Content.Parts[0].Text

	// Gemini occasionally wraps JSON inside Markdown code fences.
	text = strings.TrimSpace(text)
	text = strings.TrimPrefix(text, "```json")
	text = strings.TrimPrefix(text, "```")
	text = strings.TrimSuffix(text, "```")
	text = strings.TrimSpace(text)

	var aiResponse AIPronunciationResponse

	if err := json.Unmarshal([]byte(text), &aiResponse); err != nil {
		return nil, fmt.Errorf(
			"failed to parse Gemini JSON: %w",
			err,
		)
	}

	// PopulateIndexes(
	// 	script,
	// 	aiResponse.Words,
	// )

	PopulateIndexes(
		script,
		aiResponse.Words,
	)

	// Populate backend-controlled fields.
	for i := range aiResponse.Words {

		// Audio generation will be implemented later.
		aiResponse.Words[i].AudioAvailable = false

		// Ensure confidence stays within range.
		if aiResponse.Words[i].Confidence < 0 {
			aiResponse.Words[i].Confidence = 0
		}

		if aiResponse.Words[i].Confidence > 1 {
			aiResponse.Words[i].Confidence = 1
		}

		if aiResponse.Words[i].Difficulty == "" {
			aiResponse.Words[i].Difficulty = "Medium"
		}
	}

	result := &PronunciationResponse{
		Success: true,

		Provider: "Gemini",

		Version: "1.0",

		ProcessingTimeMs: time.Since(start).Milliseconds(),

		Count: len(aiResponse.Words),

		Words: aiResponse.Words,
	}

	return result, nil
}
