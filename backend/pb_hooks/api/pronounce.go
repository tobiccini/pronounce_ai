package api

import (
	"fmt"
	"net/http"
	"os"

	"myap/pb_hooks/ai"

	"github.com/pocketbase/pocketbase/core"
)

type PronounceRequest struct {
	Script string `json:"script"`
}

func RegisterRoutes(se *core.ServeEvent) {

	se.Router.POST("/api/pronounce", func(e *core.RequestEvent) error {

		var req PronounceRequest

		if err := e.BindBody(&req); err != nil {
			return e.BadRequestError(
				"Invalid request body",
				err,
			)
		}

		if req.Script == "" {
			return e.BadRequestError(
				"script cannot be empty",
				nil,
			)
		}
		// 1
		//apiKey := os.Getenv("GEMINI_API_KEY")
		apiKey := os.Getenv("FIREWORKS_API_KEY")

		if apiKey == "" {
			return e.InternalServerError(
				"FIREWORKS_API_KEY not configured",
				nil,
			)
		}

		// fmt.Println(os.Getenv("FIREWORKS_API_KEY"))
		fmt.Printf("starting process for script")

		var provider ai.AIProvider = &ai.FireworksProvider{
			APIKey: apiKey,

			// Model: "accounts/fireworks/models/deepseek-v4-pro",
			Model: os.Getenv("FIREWORKS_MODEL"),
		}

		cached, found, err := ai.GetCachedAnalysis(
			e.App,
			req.Script,
		)

		if err != nil {

			return e.InternalServerError(
				"Failed checking cache",
				err,
			)
		}

		if found {

			e.App.Logger().Info("========== CACHE HIT ==========")

			if err := ai.ApplyDictionaryOverrides(
				e.App,
				cached,
			); err != nil {

				e.App.Logger().Error(
					"failed to apply dictionary",
					"error",
					err,
				)
			}

			return e.JSON(
				http.StatusOK,
				cached,
			)
		}

		// ------------------------------------------------------------
		// Step 2: No cache -> Call Gemini
		// ------------------------------------------------------------

		results, err := provider.Analyze(req.Script)

		if err != nil {
			return e.InternalServerError(
				"Failed to analyze script",
				err,
			)
		}

		results.FromCache = false

		if err := ai.UpdateDictionary(
			e.App,
			results,
		); err != nil {

			e.App.Logger().Error(
				"failed updating dictionary",
				"error",
				err,
			)
		}

		if err := ai.ApplyDictionaryOverrides(
			e.App,
			results,
		); err != nil {

			e.App.Logger().Error(
				"failed to apply dictionary",
				"error",
				err,
			)
		}

		// ------------------------------------------------------------
		// Step 3: Save
		// ------------------------------------------------------------

		if err := ai.SaveAnalysis(
			e.App,
			req.Script,
			results,
		); err != nil {

			e.App.Logger().Error(
				"failed to save analysis",
				"error",
				err,
			)
		}

		// if err := ai.UpdateDictionary(
		// 	e.App,
		// 	results,
		// ); err != nil {

		// 	e.App.Logger().Error(
		// 		"failed to update dictionary",
		// 		"error",
		// 		err,
		// 	)
		// }

		return e.JSON(
			http.StatusOK,
			results,
		)

	})

}
