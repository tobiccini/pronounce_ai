package ai

import (
	"crypto/sha256"
	"encoding/hex"

	"database/sql"
	"errors"
	"time"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	//"github.com/pocketbase/pocketbase/tools/types"
)

func HashScript(script string) string {
	hash := sha256.Sum256([]byte(script))
	return hex.EncodeToString(hash[:])
}

// ============================================================================
// Cache
// ============================================================================

func GetCachedAnalysis(
	app core.App,
	script string,
) (*PronunciationResponse, bool, error) {

	start := time.Now()

	hash := HashScript(script)

	analysis, err := app.FindFirstRecordByFilter(
		"analyses",
		"scriptHash = {:hash}",
		dbx.Params{
			"hash": hash,
		},
	)

	// Normal cache miss.
	if errors.Is(err, sql.ErrNoRows) {
		return nil, false, nil
	}

	// Actual database problem.
	if err != nil {
		return nil, false, err
	}
	pronunciations, err := app.FindRecordsByFilter(
		"pronunciations",
		"analysis = {:id}",
		"",
		0,
		0,
		dbx.Params{
			"id": analysis.Id,
		},
	)

	if err != nil {
		return nil, false, err
	}

	response := &PronunciationResponse{
		Success: true,

		Provider: analysis.GetString("provider"),

		Version: analysis.GetString("model"),

		// Real cache retrieval time
		ProcessingTimeMs: time.Since(start).Milliseconds(),

		Count: len(pronunciations),

		FromCache: true,
	}

	for _, record := range pronunciations {

		response.Words = append(
			response.Words,
			PronunciationResult{
				Word:              record.GetString("word"),
				Category:          record.GetString("category"),
				Language:          record.GetString("language"),
				EasyPronunciation: record.GetString("easyPronunciation"),
				IPAEnglish:        record.GetString("ipaEnglish"),
				IPANative:         record.GetString("ipaNative"),
				Meaning:           record.GetString("meaning"),
				Replacement:       record.GetString("replacement"),
				PresenterTips:     record.GetStringSlice("presenterTips"),
				Confidence:        record.GetFloat("confidence"),
				AudioAvailable:    record.GetBool("audioAvailable"),
				Difficulty:        record.GetString("difficulty"),
				StartIndex:        record.GetInt("startIndex"),
				EndIndex:          record.GetInt("endIndex"),
			},
		)
	}

	return response, true, nil
}

// ============================================================================
// Save
// ============================================================================

func SaveAnalysis(
	app core.App,
	script string,
	response *PronunciationResponse,
) error {

	analysisCollection, err := app.FindCollectionByNameOrId("analyses")
	if err != nil {
		return err
	}

	pronCollection, err := app.FindCollectionByNameOrId("pronunciations")
	if err != nil {
		return err
	}

	analysis := core.NewRecord(analysisCollection)

	analysis.Set("script", script)
	analysis.Set("scriptHash", HashScript(script))
	analysis.Set("provider", response.Provider)
	// analysis.Set("model", response.Version)
	analysis.Set("model", response.Model)
	analysis.Set("processingTimeMs", response.ProcessingTimeMs)

	if err := app.Save(analysis); err != nil {
		return err
	}

	for _, word := range response.Words {

		record := core.NewRecord(pronCollection)

		record.Set("analysis", analysis.Id)

		record.Set("word", word.Word)
		record.Set("category", word.Category)
		record.Set("language", word.Language)
		record.Set("easyPronunciation", word.EasyPronunciation)
		record.Set("ipaEnglish", word.IPAEnglish)
		record.Set("ipaNative", word.IPANative)

		record.Set("meaning", word.Meaning)
		record.Set("notes", "")
		record.Set("replacement", word.Replacement)
		record.Set("presenterTips", word.PresenterTips)

		record.Set("confidence", word.Confidence)
		record.Set("audioAvailable", word.AudioAvailable)
		record.Set("difficulty", word.Difficulty)

		record.Set("startIndex", word.StartIndex)
		record.Set("endIndex", word.EndIndex)

		if err := app.Save(record); err != nil {
			return err
		}
	}

	return nil
}
