// package ai

// import (
// 	"database/sql"
// 	"errors"
// 	"strings"

// 	"github.com/pocketbase/dbx"
// 	"github.com/pocketbase/pocketbase/core"
// )

// func UpdateDictionary(
// 	app core.App,
// 	response *PronunciationResponse,
// ) error {

// 	collection, err := app.FindCollectionByNameOrId("dictionary")
// 	if err != nil {
// 		return err
// 	}

// 	for _, word := range response.Words {

// 		name := strings.TrimSpace(word.Word)

// 		// Skip empty words
// 		if name == "" {
// 			continue
// 		}

// 		// Check whether it already exists
// 		_, err := app.FindFirstRecordByFilter(
// 			collection,
// 			"word = {:word}",
// 			dbx.Params{
// 				"word": name,
// 			},
// 		)

// 		// Already exists
// 		if err == nil {
// 			continue
// 		}

// 		// Unexpected database error
// 		if !errors.Is(err, sql.ErrNoRows) {
// 			return err
// 		}

// 		// Doesn't exist → create it
// 		record := core.NewRecord(collection)

// 		record.Set("word", word.Word)
// 		record.Set("category", word.Category)
// 		record.Set("language", word.Language)
// 		record.Set("easyPronunciation", word.EasyPronunciation)
// 		record.Set("ipaEnglish", word.IPAEnglish)
// 		record.Set("ipaNative", word.IPANative)

// 		// Store the AI meaning as notes so editors can improve it later
// 		record.Set("notes", word.Meaning)

// 		if err := app.Save(record); err != nil {
// 			return err
// 		}

// 		app.Logger().Info(
// 			"DICTIONARY ENTRY ADDED",
// 			"word", word.Word,
// 		)
// 	}

// 	return nil
// }


package ai

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

func UpdateDictionary(
	app core.App,
	response *PronunciationResponse,
) error {

	collection, err := app.FindCollectionByNameOrId("dictionary")
	if err != nil {
		return err
	}

	for _, word := range response.Words {

		name := strings.TrimSpace(word.Word)

		// Skip empty words
		if name == "" {
			continue
		}

		// Check whether it already exists
		_, err := app.FindFirstRecordByFilter(
			collection,
			"word = {:word}",
			dbx.Params{
				"word": name,
			},
		)

		// Already exists
		if err == nil {
			continue
		}

		// Unexpected database error
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}

		// Doesn't exist → create it
		record := core.NewRecord(collection)

		record.Set("word", word.Word)
		record.Set("category", word.Category)
		record.Set("language", word.Language)
		record.Set("easyPronunciation", word.EasyPronunciation)
		record.Set("ipaEnglish", word.IPAEnglish)
		record.Set("ipaNative", word.IPANative)

		// Store the AI meaning as notes so editors can improve it later
		record.Set("notes", word.Meaning)

		if err := app.Save(record); err != nil {
			return err
		}

		app.Logger().Info(
			"DICTIONARY ENTRY ADDED",
			"word", word.Word,
		)
	}

	return nil
}