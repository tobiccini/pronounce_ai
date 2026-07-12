// package api

// import (
// 	"fmt"
// 	"net/http"
// 	"os"

// 	"myap/pb_hooks/ai"

// 	"github.com/pocketbase/pocketbase/core"
// )

// type PronounceRequest struct {
// 	Script string `json:"script"`
// }

// func RegisterRoutes(se *core.ServeEvent) {

// 	se.Router.POST("/api/pronounce", func(e *core.RequestEvent) error {

// 		var req PronounceRequest

// 		if err := e.BindBody(&req); err != nil {
// 			return e.BadRequestError(
// 				"Invalid request body",
// 				err,
// 			)
// 		}

// 		if req.Script == "" {
// 			return e.BadRequestError(
// 				"script cannot be empty",
// 				nil,
// 			)
// 		}
// 		// 1
// 		//apiKey := os.Getenv("GEMINI_API_KEY")
// 		apiKey := os.Getenv("FIREWORKS_API_KEY")

// 		if apiKey == "" {
// 			return e.InternalServerError(
// 				"FIREWORKS_API_KEY not configured",
// 				nil,
// 			)
// 		}

// 		// fmt.Println(os.Getenv("FIREWORKS_API_KEY"))
// 		fmt.Printf("starting process for script")

// 		var provider ai.AIProvider = &ai.FireworksProvider{
// 			APIKey: apiKey,

// 			// Model: "accounts/fireworks/models/deepseek-v4-pro",
// 			Model: os.Getenv("FIREWORKS_MODEL"),
// 		}

// 		cached, found, err := ai.GetCachedAnalysis(
// 			e.App,
// 			req.Script,
// 		)

// 		if err != nil {

// 			return e.InternalServerError(
// 				"Failed checking cache",
// 				err,
// 			)
// 		}

// 		if found {

// 			e.App.Logger().Info("========== CACHE HIT ==========")

// 			if err := ai.ApplyDictionaryOverrides(
// 				e.App,
// 				cached,
// 			); err != nil {

// 				e.App.Logger().Error(
// 					"failed to apply dictionary",
// 					"error",
// 					err,
// 				)
// 			}

// 			return e.JSON(
// 				http.StatusOK,
// 				cached,
// 			)
// 		}

// 		// ------------------------------------------------------------
// 		// Step 2: No cache -> Call Gemini
// 		// ------------------------------------------------------------

// 		results, err := provider.Analyze(req.Script)

// 		if err != nil {
// 			return e.InternalServerError(
// 				"Failed to analyze script",
// 				err,
// 			)
// 		}

// 		results.FromCache = false

// 		if err := ai.UpdateDictionary(
// 			e.App,
// 			results,
// 		); err != nil {

// 			e.App.Logger().Error(
// 				"failed updating dictionary",
// 				"error",
// 				err,
// 			)
// 		}

// 		if err := ai.ApplyDictionaryOverrides(
// 			e.App,
// 			results,
// 		); err != nil {

// 			e.App.Logger().Error(
// 				"failed to apply dictionary",
// 				"error",
// 				err,
// 			)
// 		}

// 		// ------------------------------------------------------------
// 		// Step 3: Save
// 		// ------------------------------------------------------------

// 		if err := ai.SaveAnalysis(
// 			e.App,
// 			req.Script,
// 			results,
// 		); err != nil {

// 			e.App.Logger().Error(
// 				"failed to save analysis",
// 				"error",
// 				err,
// 			)
// 		}

// 		// if err := ai.UpdateDictionary(
// 		// 	e.App,
// 		// 	results,
// 		// ); err != nil {

// 		// 	e.App.Logger().Error(
// 		// 		"failed to update dictionary",
// 		// 		"error",
// 		// 		err,
// 		// 	)
// 		// }

// 		return e.JSON(
// 			http.StatusOK,
// 			results,
// 		)

// 	})

// }




// package api

// import (
// 	"fmt"
// 	"net/http"
// 	"os"

// 	"myap/pb_hooks/ai"

// 	"github.com/pocketbase/pocketbase/core"
// )

// type PronounceRequest struct {
// 	Script string `json:"script"`
// }

// func RegisterRoutes(se *core.ServeEvent) {

// 	// Friendly landing page for judges/visitors hitting the root URL
// 	se.Router.GET("/", func(e *core.RequestEvent) error {
// 		html := `<!DOCTYPE html>
// <html>
// <head><title>PronounceAI</title></head>
// <body style="font-family: sans-serif; background:#0d1117; color:#c9d1d9; padding:2rem;">
//   <h1>🎙️ PronounceAI is live</h1>
//   <p>AI-powered name pronunciation assistant for broadcasters.</p>
//   <p>Health check: <a href="/api/health" style="color:#58a6ff;">/api/health</a></p>
//   <p>Main endpoint: <code>POST /api/pronounce</code> (send JSON: <code>{"script": "..."}</code>)</p>
// </body>
// </html>`
// 		return e.HTML(http.StatusOK, html)
// 	})

// 	// Friendly response if someone visits /api/pronounce directly in a browser (GET)
// 	se.Router.GET("/api/pronounce", func(e *core.RequestEvent) error {
// 		return e.JSON(http.StatusOK, map[string]any{
// 			"status":  "ok",
// 			"message": "This endpoint is live. Send a POST request with JSON body { \"script\": \"...\" } to analyze pronunciations.",
// 		})
// 	})

// 	se.Router.POST("/api/pronounce", func(e *core.RequestEvent) error {

// 		var req PronounceRequest

// 		if err := e.BindBody(&req); err != nil {
// 			return e.BadRequestError(
// 				"Invalid request body",
// 				err,
// 			)
// 		}

// 		if req.Script == "" {
// 			return e.BadRequestError(
// 				"script cannot be empty",
// 				nil,
// 			)
// 		}
// 		// 1
// 		//apiKey := os.Getenv("GEMINI_API_KEY")
// 		apiKey := os.Getenv("FIREWORKS_API_KEY")

// 		if apiKey == "" {
// 			return e.InternalServerError(
// 				"FIREWORKS_API_KEY not configured",
// 				nil,
// 			)
// 		}

// 		// fmt.Println(os.Getenv("FIREWORKS_API_KEY"))
// 		fmt.Printf("starting process for script")

// 		var provider ai.AIProvider = &ai.FireworksProvider{
// 			APIKey: apiKey,

// 			// Model: "accounts/fireworks/models/deepseek-v4-pro",
// 			Model: os.Getenv("FIREWORKS_MODEL"),
// 		}

// 		cached, found, err := ai.GetCachedAnalysis(
// 			e.App,
// 			req.Script,
// 		)

// 		if err != nil {

// 			return e.InternalServerError(
// 				"Failed checking cache",
// 				err,
// 			)
// 		}

// 		if found {

// 			e.App.Logger().Info("========== CACHE HIT ==========")

// 			if err := ai.ApplyDictionaryOverrides(
// 				e.App,
// 				cached,
// 			); err != nil {

// 				e.App.Logger().Error(
// 					"failed to apply dictionary",
// 					"error",
// 					err,
// 				)
// 			}

// 			return e.JSON(
// 				http.StatusOK,
// 				cached,
// 			)
// 		}

// 		// ------------------------------------------------------------
// 		// Step 2: No cache -> Call Gemini
// 		// ------------------------------------------------------------

// 		results, err := provider.Analyze(req.Script)

// 		if err != nil {
// 			return e.InternalServerError(
// 				"Failed to analyze script",
// 				err,
// 			)
// 		}

// 		results.FromCache = false

// 		if err := ai.UpdateDictionary(
// 			e.App,
// 			results,
// 		); err != nil {

// 			e.App.Logger().Error(
// 				"failed updating dictionary",
// 				"error",
// 				err,
// 			)
// 		}

// 		if err := ai.ApplyDictionaryOverrides(
// 			e.App,
// 			results,
// 		); err != nil {

// 			e.App.Logger().Error(
// 				"failed to apply dictionary",
// 				"error",
// 				err,
// 			)
// 		}

// 		// ------------------------------------------------------------
// 		// Step 3: Save
// 		// ------------------------------------------------------------

// 		if err := ai.SaveAnalysis(
// 			e.App,
// 			req.Script,
// 			results,
// 		); err != nil {

// 			e.App.Logger().Error(
// 				"failed to save analysis",
// 				"error",
// 				err,
// 			)
// 		}

// 		// if err := ai.UpdateDictionary(
// 		// 	e.App,
// 		// 	results,
// 		// ); err != nil {

// 		// 	e.App.Logger().Error(
// 		// 		"failed to update dictionary",
// 		// 		"error",
// 		// 		err,
// 		// 	)
// 		// }

// 		return e.JSON(
// 			http.StatusOK,
// 			results,
// 		)

// 	})

// }



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

	// Interactive demo page for judges/visitors hitting the root URL
	se.Router.GET("/", func(e *core.RequestEvent) error {
		html := `<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>PronounceAI</title>
<style>
  body {
    font-family: 'Segoe UI', sans-serif;
    background: #0d1117;
    color: #c9d1d9;
    padding: 2rem;
    max-width: 720px;
    margin: 0 auto;
  }
  h1 { color: #58a6ff; margin-bottom: 0.2rem; }
  p.sub { color: #8b949e; margin-top: 0; }
  textarea {
    width: 100%;
    min-height: 120px;
    background: #161b22;
    color: #c9d1d9;
    border: 1px solid #30363d;
    border-radius: 8px;
    padding: 0.75rem;
    font-size: 1rem;
    box-sizing: border-box;
    resize: vertical;
  }
  button {
    margin-top: 0.75rem;
    background: #238636;
    color: white;
    border: none;
    padding: 0.6rem 1.2rem;
    border-radius: 6px;
    font-size: 1rem;
    cursor: pointer;
  }
  button:disabled { background: #30363d; cursor: not-allowed; }
  #result { margin-top: 1.5rem; }
  .meta { color: #8b949e; font-size: 0.85rem; margin-bottom: 1rem; }
  .card {
    background: #161b22;
    border: 1px solid #30363d;
    border-radius: 8px;
    padding: 1rem;
    margin-bottom: 1rem;
  }
  .card h3 { margin: 0 0 0.5rem 0; color: #58a6ff; }
  .badge {
    background: #30363d;
    color: #c9d1d9;
    font-size: 0.75rem;
    padding: 0.15rem 0.5rem;
    border-radius: 4px;
    margin-left: 0.5rem;
  }
  .card ul { margin: 0.5rem 0 0 1.2rem; padding: 0; }
  .error { color: #f85149; }
  code { background: #161b22; padding: 0.1rem 0.4rem; border-radius: 4px; }
  footer { margin-top: 2rem; color: #8b949e; font-size: 0.8rem; }
  a { color: #58a6ff; }
</style>
</head>
<body>
  <h1>PronounceAI</h1>
  <p class="sub">Paste a news script below to see how unfamiliar names should be pronounced.</p>

  <textarea id="script" placeholder="e.g. President Volodymyr Zelenskyy met with NATO Secretary General Mark Rutte in The Hague."></textarea>
  <br>
  <button id="btn" onclick="analyze()">Analyze Pronunciation</button>

  <div id="result"></div>

  <footer>
    Health check: <a href="/api/health">/api/health</a> &middot;
    API: <code>POST /api/pronounce</code>
  </footer>

<script>
function escapeHtml(str) {
  if (str === undefined || str === null) return '';
  return String(str)
    .replace(/&/g, '&amp;')
    .replace(/</g, '&lt;')
    .replace(/>/g, '&gt;');
}

async function analyze() {
  var input = document.getElementById('script').value.trim();
  var resultEl = document.getElementById('result');
  var btn = document.getElementById('btn');

  if (!input) {
    resultEl.innerHTML = '<p class="error">Please enter a script first.</p>';
    return;
  }

  btn.disabled = true;
  btn.textContent = 'Analyzing...';
  resultEl.innerHTML = '<p class="meta">Analyzing script, please wait...</p>';

  try {
    var res = await fetch('/api/pronounce', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ script: input })
    });

    var data = await res.json();
    render(data);
  } catch (err) {
    resultEl.innerHTML = '<p class="error">Request failed: ' + escapeHtml(err.message) + '</p>';
  }

  btn.disabled = false;
  btn.textContent = 'Analyze Pronunciation';
}

function render(data) {
  var resultEl = document.getElementById('result');

  if (!data || data.success === false) {
    var msg = (data && data.message) ? data.message : 'Something went wrong.';
    resultEl.innerHTML = '<p class="error">' + escapeHtml(msg) + '</p>';
    return;
  }

  if (!data.words || data.words.length === 0) {
    resultEl.innerHTML = '<p class="meta">No unfamiliar names detected in this script.</p>';
    return;
  }

  var html = '<div class="meta">Provider: ' + escapeHtml(data.provider) +
    ' &middot; ' + data.count + ' word(s) found' +
    ' &middot; ' + (data.fromCache ? 'from cache' : 'fresh analysis') +
    ' &middot; ' + data.processingTimeMs + 'ms</div>';

  data.words.forEach(function (w) {
    html += '<div class="card">';
    html += '<h3>' + escapeHtml(w.word) + '<span class="badge">' + escapeHtml(w.difficulty) + '</span></h3>';
    html += '<p><strong>Easy:</strong> ' + escapeHtml(w.easyPronunciation) + '</p>';
    html += '<p><strong>IPA (English):</strong> ' + escapeHtml(w.ipaEnglish) + '</p>';
    html += '<p><strong>Language:</strong> ' + escapeHtml(w.language) + ' &middot; <strong>Meaning:</strong> ' + escapeHtml(w.meaning) + '</p>';

    if (w.presenterTips && w.presenterTips.length > 0) {
      html += '<ul>';
      w.presenterTips.forEach(function (tip) {
        html += '<li>' + escapeHtml(tip) + '</li>';
      });
      html += '</ul>';
    }

    html += '</div>';
  });

  resultEl.innerHTML = html;
}
</script>
</body>
</html>`
		return e.HTML(http.StatusOK, html)
	})

	// Friendly response if someone visits /api/pronounce directly in a browser (GET)
	se.Router.GET("/api/pronounce", func(e *core.RequestEvent) error {
		return e.JSON(http.StatusOK, map[string]any{
			"status":  "ok",
			"message": "This endpoint is live. Send a POST request with JSON body { \"script\": \"...\" } to analyze pronunciations, or use the form on the homepage.",
		})
	})

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
