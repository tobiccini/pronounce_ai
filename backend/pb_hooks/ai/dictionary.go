package ai

import (
	"strings"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

func ApplyDictionaryOverrides(
	app core.App,
	response *PronunciationResponse,
) error {
	app.Logger().Info("========== ENTERED DICTIONARY ==========")

	collection, err := app.FindCollectionByNameOrId("dictionary")
	if err != nil {
		return err
	}

	for i := range response.Words {

		word := strings.TrimSpace(response.Words[i].Word)

		record, err := app.FindFirstRecordByFilter(
			collection,
			"word = {:word}",
			dbx.Params{
				"word": word,
			},
		)

		if err != nil {

			app.Logger().Error(
				"DICTIONARY LOOKUP FAILED",
				"word", word,
				"error", err,
			)

			continue
		}
		app.Logger().Info(
			"DICTIONARY MATCH",
			"word",
			word,
		)

		if v := record.GetString("easyPronunciation"); v != "" {
			response.Words[i].EasyPronunciation = v
		}

		if v := record.GetString("ipaEnglish"); v != "" {
			response.Words[i].IPAEnglish = v
		}

		if v := record.GetString("ipaNative"); v != "" {
			response.Words[i].IPANative = v
		}

		if v := record.GetString("category"); v != "" {
			response.Words[i].Category = v
		}

		if v := record.GetString("language"); v != "" {
			response.Words[i].Language = v
		}

		if v := record.GetString("notes"); v != "" {
			response.Words[i].Meaning = v
		}
	}

	return nil
}
