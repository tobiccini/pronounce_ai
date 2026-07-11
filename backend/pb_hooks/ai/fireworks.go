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

type FireworksProvider struct {
	APIKey string
	Model  string
	// // jj
	// FallbackModel string
}

type FireworksRequest struct {
	Model       string             `json:"model"`
	MaxTokens   int                `json:"max_tokens"`
	Messages    []FireworksMessage `json:"messages"`
	Temperature float64            `json:"temperature,omitempty"`
}

type FireworksMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type FireworksResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func (f *FireworksProvider) Analyze(script string) (*PronunciationResponse, error) {

	response, err := f.analyzeWithModel(script, f.Model)

	if err == nil {
		return response, nil
	}

	fmt.Printf("Primary model failed (%s): %v\n", f.Model, err)

	// Don't retry if we're already using Flash
	if f.Model == "accounts/fireworks/models/nemotron-3-ultra-nvfp4" {
		return nil, err
	}
	// jj
	fallback := "accounts/fireworks/models/nemotron-3-ultra-nvfp4"

	fmt.Printf("Retrying with fallback model: %s\n", fallback)

	return f.analyzeWithModel(script, fallback)
}

func (f *FireworksProvider) analyzeWithModel(
	script string,
	model string,
) (*PronunciationResponse, error) {

	start := time.Now()

	prompt := BuildPronunciationPrompt(script)

	req := FireworksRequest{
		Model:       model,
		MaxTokens:   4096,
		Temperature: 0,
		Messages: []FireworksMessage{
			{
				Role:    "system",
				Content: "You are an expert pronunciation coach for professional radio and television broadcasters. Return ONLY valid JSON.",
			},
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest(
		http.MethodPost,
		"https://api.fireworks.ai/inference/v1/chat/completions",
		bytes.NewBuffer(body),
	)

	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Authorization", "Bearer "+f.APIKey)
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "application/json")

	client := &http.Client{}

	resp, err := client.Do(httpReq)
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
			"fireworks returned status %d: %s",
			resp.StatusCode,
			string(responseBody),
		)
	}

	var fwResponse FireworksResponse

	if err := json.Unmarshal(responseBody, &fwResponse); err != nil {
		return nil, err
	}

	if len(fwResponse.Choices) == 0 {
		return nil, fmt.Errorf("no choices returned")
	}

	// text := fwResponse.Choices[0].Message.Content

	// text = strings.TrimSpace(text)
	// text = strings.TrimPrefix(text, "```json")
	// text = strings.TrimPrefix(text, "```")
	// text = strings.TrimSuffix(text, "```")
	// text = strings.TrimSpace(text)

	// var aiResponse AIPronunciationResponse

	// if err := json.Unmarshal([]byte(text), &aiResponse); err != nil {
	// 	return nil, fmt.Errorf("failed parsing Fireworks JSON: %w", err)
	// }

	text := CleanJSON(fwResponse.Choices[0].Message.Content)

	var aiResponse AIPronunciationResponse

	if err := json.Unmarshal([]byte(text), &aiResponse); err != nil {

		fmt.Println("========== RAW FIREWORKS RESPONSE ==========")
		fmt.Println(text)
		fmt.Println("===========================================")

		return nil, fmt.Errorf("failed parsing Fireworks JSON: %w", err)
	}

	PopulateIndexes(script, aiResponse.Words)

	for i := range aiResponse.Words {

		aiResponse.Words[i].AudioAvailable = false

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

	return &PronunciationResponse{
		Success:  true,
		Provider: "Fireworks",
		// Version:          "1.0",
		// Model:            model,

		Version:          model,
		Model:            model,
		ProcessingTimeMs: time.Since(start).Milliseconds(),
		Count:            len(aiResponse.Words),
		Words:            aiResponse.Words,
	}, nil
}

func CleanJSON(text string) string {

	text = strings.TrimSpace(text)

	text = strings.TrimPrefix(text, "```json")
	text = strings.TrimPrefix(text, "```")
	text = strings.TrimSuffix(text, "```")

	text = strings.Trim(text, "\"'")
	text = strings.TrimSpace(text)

	start := strings.Index(text, "{")
	end := strings.LastIndex(text, "}")

	if start >= 0 && end > start {
		return text[start : end+1]
	}

	return text
}

// func CleanJSON(text string) string {

// 	text = strings.TrimSpace(text)

// 	text = strings.TrimPrefix(text, "```json")
// 	text = strings.TrimPrefix(text, "```")
// 	text = strings.TrimSuffix(text, "```")

// 	text = strings.TrimSpace(text)

// 	start := strings.Index(text, "{")
// 	end := strings.LastIndex(text, "}")

// 	if start >= 0 && end > start {
// 		return text[start : end+1]
// 	}

// 	return text
// }
