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

	// Original terminal-inspired demo page for judges/visitors hitting the root URL
	se.Router.GET("/", func(e *core.RequestEvent) error {
		html := `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>PronounceAI</title>
<link rel="preconnect" href="https://fonts.googleapis.com">
<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
<link href="https://fonts.googleapis.com/css2?family=JetBrains+Mono:wght@400;500;600;700;800&display=swap" rel="stylesheet">
<style>
  * { box-sizing: border-box; }

  :root {
    --bg: #050708;
    --panel: #0c0f12;
    --line: #1c2226;
    --text: #d7dde1;
    --dim: #5c6670;
    --accent: #00ff9c;
    --accent2: #00c8ff;
    --warn: #ffb84d;
    --danger: #ff5470;
  }

  body {
    margin: 0;
    background:
      linear-gradient(var(--bg), var(--bg)),
      repeating-linear-gradient(0deg, rgba(255,255,255,0.015) 0px, rgba(255,255,255,0.015) 1px, transparent 1px, transparent 3px);
    color: var(--text);
    font-family: 'JetBrains Mono', monospace;
    padding: 2.5rem 1.25rem 5rem;
    min-height: 100vh;
  }

  .wrap { max-width: 760px; margin: 0 auto; }

  .top-tag {
    display: inline-flex;
    align-items: center;
    gap: 0.5rem;
    color: var(--accent);
    font-size: 0.72rem;
    letter-spacing: 0.15em;
    text-transform: uppercase;
    border: 1px solid rgba(0,255,156,0.3);
    padding: 0.25rem 0.7rem;
    border-radius: 999px;
    margin-bottom: 1.2rem;
  }

  .top-tag .pulse {
    width: 6px; height: 6px; border-radius: 50%;
    background: var(--accent);
    box-shadow: 0 0 8px var(--accent);
    animation: pulse 1.6s ease-in-out infinite;
  }

  @keyframes pulse { 0%,100% { opacity: 1; } 50% { opacity: 0.3; } }

  h1 {
    font-size: 2rem;
    margin: 0 0 0.4rem;
    font-weight: 800;
    letter-spacing: -0.01em;
  }

  h1 .cursor {
    display: inline-block;
    width: 0.5em;
    background: var(--accent);
    margin-left: 4px;
    animation: blink 1s step-end infinite;
  }

  @keyframes blink { 50% { opacity: 0; } }

  p.sub {
    color: var(--dim);
    margin: 0 0 2rem;
    font-size: 0.9rem;
    line-height: 1.6;
    max-width: 52ch;
  }

  .term {
    background: var(--panel);
    border: 1px solid var(--line);
    border-radius: 10px;
    overflow: hidden;
  }

  .term-bar {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    padding: 0.6rem 0.9rem;
    border-bottom: 1px solid var(--line);
    background: #0a0d0f;
  }

  .term-bar span {
    width: 10px; height: 10px; border-radius: 50%;
    background: #2a3237;
  }

  .term-bar .label {
    margin-left: 0.5rem;
    color: var(--dim);
    font-size: 0.75rem;
    letter-spacing: 0.03em;
  }

  .term-body { padding: 1rem 1.1rem; }

  .prompt-line {
    color: var(--accent);
    font-size: 0.85rem;
    margin-bottom: 0.4rem;
  }

  textarea {
    width: 100%;
    min-height: 100px;
    background: transparent;
    color: var(--text);
    border: none;
    outline: none;
    font-family: inherit;
    font-size: 0.95rem;
    line-height: 1.6;
    resize: vertical;
  }

  textarea::placeholder { color: #3a4248; }

  .term-footer {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0.8rem 1.1rem;
    border-top: 1px solid var(--line);
  }

  .hint { color: var(--dim); font-size: 0.72rem; }

  button#btn {
    background: var(--accent);
    color: #04140c;
    border: none;
    padding: 0.6rem 1.3rem;
    border-radius: 6px;
    font-family: inherit;
    font-weight: 700;
    font-size: 0.82rem;
    letter-spacing: 0.02em;
    text-transform: uppercase;
    cursor: pointer;
    transition: box-shadow 0.15s ease, transform 0.15s ease;
  }

  button#btn:hover { box-shadow: 0 0 18px rgba(0,255,156,0.45); transform: translateY(-1px); }
  button#btn:disabled {
    background: #1c2226;
    color: var(--dim);
    cursor: not-allowed;
    box-shadow: none;
    transform: none;
  }

  #result { margin-top: 2rem; }

  .status-line {
    color: var(--dim);
    font-size: 0.85rem;
    display: flex;
    align-items: center;
    gap: 0.6rem;
  }

  .spinner {
    width: 12px; height: 12px;
    border: 2px solid var(--line);
    border-top-color: var(--accent);
    border-radius: 50%;
    animation: spin 0.7s linear infinite;
  }

  @keyframes spin { to { transform: rotate(360deg); } }

  .meta {
    color: var(--dim);
    font-size: 0.75rem;
    margin-bottom: 1.2rem;
    padding-bottom: 0.8rem;
    border-bottom: 1px dashed var(--line);
  }

  .meta b { color: var(--text); font-weight: 600; }

  .card {
    background: var(--panel);
    border: 1px solid var(--line);
    border-left: 3px solid var(--accent);
    border-radius: 8px;
    padding: 1.3rem 1.4rem;
    margin-bottom: 1.1rem;
  }

  .card.medium { border-left-color: var(--warn); }
  .card.hard { border-left-color: var(--danger); }

  .card-head {
    display: flex;
    align-items: flex-start;
    justify-content: space-between;
    flex-wrap: wrap;
    gap: 0.6rem;
    margin-bottom: 0.9rem;
  }

  .word-title {
    font-size: 1.25rem;
    font-weight: 700;
    color: #fff;
  }

  .tags { display: flex; gap: 0.4rem; flex-wrap: wrap; }

  .tag {
    font-size: 0.68rem;
    color: var(--dim);
    border: 1px solid var(--line);
    padding: 0.15rem 0.5rem;
    border-radius: 4px;
    letter-spacing: 0.03em;
  }

  .tag.difficulty-easy { color: var(--accent); border-color: rgba(0,255,156,0.35); }
  .tag.difficulty-medium { color: var(--warn); border-color: rgba(255,184,77,0.35); }
  .tag.difficulty-hard { color: var(--danger); border-color: rgba(255,84,112,0.35); }

  .field { margin-top: 0.85rem; }

  .field .k {
    color: var(--accent2);
    font-size: 0.7rem;
    letter-spacing: 0.05em;
  }

  .field .k:before { content: "// "; color: #2a3237; }

  .easy-pron {
    font-size: 1.05rem;
    font-weight: 600;
    color: #fff;
    margin-top: 0.2rem;
  }

  .ipa-row { display: flex; gap: 2rem; flex-wrap: wrap; }
  .ipa-row .field { flex: 1; min-width: 150px; }

  .ipa-value {
    font-size: 0.9rem;
    color: #b8c2ca;
    margin-top: 0.2rem;
  }

  .meaning-value {
    font-size: 0.88rem;
    color: #b8c2ca;
    margin-top: 0.2rem;
    line-height: 1.55;
  }

  .replacement-box {
    margin-top: 0.3rem;
    background: rgba(0,255,156,0.06);
    border: 1px solid rgba(0,255,156,0.25);
    color: #b6ffe0;
    border-radius: 6px;
    padding: 0.6rem 0.8rem;
    font-size: 0.85rem;
    line-height: 1.5;
  }

  .replacement-box:before { content: "+ "; color: var(--accent); font-weight: 700; }

  .tips { margin: 0.3rem 0 0; padding: 0; list-style: none; }

  .tips li {
    display: flex;
    gap: 0.5rem;
    font-size: 0.85rem;
    color: #b8c2ca;
    line-height: 1.55;
    margin-top: 0.4rem;
  }

  .tips li:before { content: ">"; color: var(--accent2); flex-shrink: 0; }

  .error {
    color: var(--danger);
    background: rgba(255,84,112,0.06);
    border: 1px solid rgba(255,84,112,0.3);
    padding: 0.8rem 1rem;
    border-radius: 8px;
    font-size: 0.85rem;
  }

  .empty {
    color: var(--dim);
    font-size: 0.85rem;
    padding: 1rem 0;
  }

  footer {
    margin-top: 3rem;
    color: var(--dim);
    font-size: 0.75rem;
    border-top: 1px solid var(--line);
    padding-top: 1.2rem;
  }

  footer a { color: var(--accent2); text-decoration: none; }
  footer a:hover { text-decoration: underline; }

  code {
    background: #0a0d0f;
    color: var(--accent);
    padding: 0.12rem 0.4rem;
    border-radius: 4px;
    font-size: 0.8rem;
    border: 1px solid var(--line);
  }
</style>
</head>
<body>
<div class="wrap">

  <div class="top-tag"><span class="pulse"></span> system online</div>
  <h1>PronounceAI<span class="cursor">&nbsp;</span></h1>
  <p class="sub">Paste a news script. We flag unfamiliar names and return presenter-ready pronunciation guidance.</p>

  <div class="term">
    <div class="term-bar">
      <span></span><span></span><span></span>
      <span class="label">script.txt</span>
    </div>
    <div class="term-body">
      <div class="prompt-line">$ paste-script --analyze</div>
      <textarea id="script" placeholder="President Volodymyr Zelenskyy met with NATO Secretary General Mark Rutte in The Hague."></textarea>
    </div>
    <div class="term-footer">
      <span class="hint">ctrl not required, just click below</span>
      <button id="btn" onclick="analyze()">Run analysis</button>
    </div>
  </div>

  <div id="result"></div>

  <footer>
    health: <a href="/api/health">/api/health</a> &nbsp;&middot;&nbsp;
    endpoint: <code>POST /api/pronounce</code>
  </footer>

</div>

<script>
function escapeHtml(str) {
  if (str === undefined || str === null) return '';
  return String(str)
    .replace(/&/g, '&amp;')
    .replace(/</g, '&lt;')
    .replace(/>/g, '&gt;');
}

function difficultyClass(diff) {
  var d = (diff || '').toLowerCase();
  if (d === 'easy') return 'easy';
  if (d === 'hard') return 'hard';
  return 'medium';
}

async function analyze() {
  var input = document.getElementById('script').value.trim();
  var resultEl = document.getElementById('result');
  var btn = document.getElementById('btn');

  if (!input) {
    resultEl.innerHTML = '<div class="error">error: script input is empty</div>';
    return;
  }

  btn.disabled = true;
  btn.textContent = 'Running...';
  resultEl.innerHTML = '<div class="status-line"><span class="spinner"></span> analyzing script...</div>';

  try {
    var res = await fetch('/api/pronounce', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ script: input })
    });

    var data = await res.json();
    render(data);
  } catch (err) {
    resultEl.innerHTML = '<div class="error">request failed: ' + escapeHtml(err.message) + '</div>';
  }

  btn.disabled = false;
  btn.textContent = 'Run analysis';
}

function render(data) {
  var resultEl = document.getElementById('result');

  if (!data || data.success === false) {
    var msg = (data && data.message) ? data.message : 'something went wrong.';
    resultEl.innerHTML = '<div class="error">' + escapeHtml(msg) + '</div>';
    return;
  }

  if (!data.words || data.words.length === 0) {
    resultEl.innerHTML = '<div class="empty">no unfamiliar names detected in this script.</div>';
    return;
  }

  var html = '<div class="meta">provider <b>' + escapeHtml(data.provider) + '</b> &nbsp;&middot;&nbsp; ' +
    '<b>' + data.count + '</b> word(s) found &nbsp;&middot;&nbsp; ' +
    (data.fromCache ? 'from cache' : 'fresh analysis') + ' &nbsp;&middot;&nbsp; ' +
    data.processingTimeMs + 'ms</div>';

  data.words.forEach(function (w) {
    var diffClass = difficultyClass(w.difficulty);

    html += '<div class="card ' + diffClass + '">';

    html += '<div class="card-head">';
    html += '<div class="word-title">' + escapeHtml(w.word) + '</div>';
    html += '<div class="tags">';
    if (w.category) html += '<span class="tag">' + escapeHtml(w.category) + '</span>';
    if (w.language) html += '<span class="tag">' + escapeHtml(w.language) + '</span>';
    if (w.difficulty) html += '<span class="tag difficulty-' + diffClass + '">' + escapeHtml(w.difficulty) + '</span>';
    html += '</div></div>';

    html += '<div class="field">';
    html += '<div class="k">easy_pronunciation</div>';
    html += '<div class="easy-pron">' + escapeHtml(w.easyPronunciation) + '</div>';
    html += '</div>';

    html += '<div class="field ipa-row">';
    html += '<div class="field"><div class="k">ipa_english</div><div class="ipa-value">' + escapeHtml(w.ipaEnglish) + '</div></div>';
    if (w.ipaNative) {
      html += '<div class="field"><div class="k">ipa_native</div><div class="ipa-value">' + escapeHtml(w.ipaNative) + '</div></div>';
    }
    html += '</div>';

    if (w.meaning) {
      html += '<div class="field">';
      html += '<div class="k">meaning</div>';
      html += '<div class="meaning-value">' + escapeHtml(w.meaning) + '</div>';
      html += '</div>';
    }

    if (w.replacement) {
      html += '<div class="field">';
      html += '<div class="k">suggested_replacement</div>';
      html += '<div class="replacement-box">' + escapeHtml(w.replacement) + '</div>';
      html += '</div>';
    }

    if (w.presenterTips && w.presenterTips.length > 0) {
      html += '<div class="field">';
      html += '<div class="k">presenter_tips</div>';
      html += '<ul class="tips">';
      w.presenterTips.forEach(function (tip) {
        html += '<li>' + escapeHtml(tip) + '</li>';
      });
      html += '</ul></div>';
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

// 	// Original terminal-inspired demo page for judges/visitors hitting the root URL
// 	se.Router.GET("/", func(e *core.RequestEvent) error {
// 		html := `<!DOCTYPE html>
// <html lang="en">
// <head>
// <meta charset="UTF-8">
// <meta name="viewport" content="width=device-width, initial-scale=1.0">
// <title>PronounceAI</title>
// <link rel="preconnect" href="https://fonts.googleapis.com">
// <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
// <link href="https://fonts.googleapis.com/css2?family=JetBrains+Mono:wght@400;500;600;700;800&display=swap" rel="stylesheet">
// <style>
//   * { box-sizing: border-box; }

//   :root {
//     --bg: #050708;
//     --panel: #0c0f12;
//     --line: #1c2226;
//     --text: #d7dde1;
//     --dim: #5c6670;
//     --accent: #00ff9c;
//     --accent2: #00c8ff;
//     --warn: #ffb84d;
//     --danger: #ff5470;
//   }

//   body {
//     margin: 0;
//     background:
//       linear-gradient(var(--bg), var(--bg)),
//       repeating-linear-gradient(0deg, rgba(255,255,255,0.015) 0px, rgba(255,255,255,0.015) 1px, transparent 1px, transparent 3px);
//     color: var(--text);
//     font-family: 'JetBrains Mono', monospace;
//     padding: 2.5rem 1.25rem 5rem;
//     min-height: 100vh;
//   }

//   .wrap { max-width: 760px; margin: 0 auto; }

//   .top-tag {
//     display: inline-flex;
//     align-items: center;
//     gap: 0.5rem;
//     color: var(--accent);
//     font-size: 0.72rem;
//     letter-spacing: 0.15em;
//     text-transform: uppercase;
//     border: 1px solid rgba(0,255,156,0.3);
//     padding: 0.25rem 0.7rem;
//     border-radius: 999px;
//     margin-bottom: 1.2rem;
//   }

//   .top-tag .pulse {
//     width: 6px; height: 6px; border-radius: 50%;
//     background: var(--accent);
//     box-shadow: 0 0 8px var(--accent);
//     animation: pulse 1.6s ease-in-out infinite;
//   }

//   @keyframes pulse { 0%,100% { opacity: 1; } 50% { opacity: 0.3; } }

//   h1 {
//     font-size: 2rem;
//     margin: 0 0 0.4rem;
//     font-weight: 800;
//     letter-spacing: -0.01em;
//   }

//   h1 .cursor {
//     display: inline-block;
//     width: 0.5em;
//     background: var(--accent);
//     margin-left: 4px;
//     animation: blink 1s step-end infinite;
//   }

//   @keyframes blink { 50% { opacity: 0; } }

//   p.sub {
//     color: var(--dim);
//     margin: 0 0 2rem;
//     font-size: 0.9rem;
//     line-height: 1.6;
//     max-width: 52ch;
//   }

//   .term {
//     background: var(--panel);
//     border: 1px solid var(--line);
//     border-radius: 10px;
//     overflow: hidden;
//   }

//   .term-bar {
//     display: flex;
//     align-items: center;
//     gap: 0.5rem;
//     padding: 0.6rem 0.9rem;
//     border-bottom: 1px solid var(--line);
//     background: #0a0d0f;
//   }

//   .term-bar span {
//     width: 10px; height: 10px; border-radius: 50%;
//     background: #2a3237;
//   }

//   .term-bar .label {
//     margin-left: 0.5rem;
//     color: var(--dim);
//     font-size: 0.75rem;
//     letter-spacing: 0.03em;
//   }

//   .term-body { padding: 1rem 1.1rem; }

//   .prompt-line {
//     color: var(--accent);
//     font-size: 0.85rem;
//     margin-bottom: 0.4rem;
//   }

//   textarea {
//     width: 100%;
//     min-height: 100px;
//     background: transparent;
//     color: var(--text);
//     border: none;
//     outline: none;
//     font-family: inherit;
//     font-size: 0.95rem;
//     line-height: 1.6;
//     resize: vertical;
//   }

//   textarea::placeholder { color: #3a4248; }

//   .term-footer {
//     display: flex;
//     align-items: center;
//     justify-content: space-between;
//     padding: 0.8rem 1.1rem;
//     border-top: 1px solid var(--line);
//   }

//   .hint { color: var(--dim); font-size: 0.72rem; }

//   button#btn {
//     background: var(--accent);
//     color: #04140c;
//     border: none;
//     padding: 0.6rem 1.3rem;
//     border-radius: 6px;
//     font-family: inherit;
//     font-weight: 700;
//     font-size: 0.82rem;
//     letter-spacing: 0.02em;
//     text-transform: uppercase;
//     cursor: pointer;
//     transition: box-shadow 0.15s ease, transform 0.15s ease;
//   }

//   button#btn:hover { box-shadow: 0 0 18px rgba(0,255,156,0.45); transform: translateY(-1px); }
//   button#btn:disabled {
//     background: #1c2226;
//     color: var(--dim);
//     cursor: not-allowed;
//     box-shadow: none;
//     transform: none;
//   }

//   #result { margin-top: 2rem; }

//   .status-line {
//     color: var(--dim);
//     font-size: 0.85rem;
//     display: flex;
//     align-items: center;
//     gap: 0.6rem;
//   }

//   .spinner {
//     width: 12px; height: 12px;
//     border: 2px solid var(--line);
//     border-top-color: var(--accent);
//     border-radius: 50%;
//     animation: spin 0.7s linear infinite;
//   }

//   @keyframes spin { to { transform: rotate(360deg); } }

//   .meta {
//     color: var(--dim);
//     font-size: 0.75rem;
//     margin-bottom: 1.2rem;
//     padding-bottom: 0.8rem;
//     border-bottom: 1px dashed var(--line);
//   }

//   .meta b { color: var(--text); font-weight: 600; }

//   .card {
//     background: var(--panel);
//     border: 1px solid var(--line);
//     border-left: 3px solid var(--accent);
//     border-radius: 8px;
//     padding: 1.3rem 1.4rem;
//     margin-bottom: 1.1rem;
//   }

//   .card.medium { border-left-color: var(--warn); }
//   .card.hard { border-left-color: var(--danger); }

//   .card-head {
//     display: flex;
//     align-items: flex-start;
//     justify-content: space-between;
//     flex-wrap: wrap;
//     gap: 0.6rem;
//     margin-bottom: 0.9rem;
//   }

//   .word-title {
//     font-size: 1.25rem;
//     font-weight: 700;
//     color: #fff;
//   }

//   .tags { display: flex; gap: 0.4rem; flex-wrap: wrap; }

//   .tag {
//     font-size: 0.68rem;
//     color: var(--dim);
//     border: 1px solid var(--line);
//     padding: 0.15rem 0.5rem;
//     border-radius: 4px;
//     letter-spacing: 0.03em;
//   }

//   .tag.difficulty-easy { color: var(--accent); border-color: rgba(0,255,156,0.35); }
//   .tag.difficulty-medium { color: var(--warn); border-color: rgba(255,184,77,0.35); }
//   .tag.difficulty-hard { color: var(--danger); border-color: rgba(255,84,112,0.35); }

//   .field { margin-top: 0.85rem; }

//   .field .k {
//     color: var(--accent2);
//     font-size: 0.7rem;
//     letter-spacing: 0.05em;
//   }

//   .field .k:before { content: "// "; color: #2a3237; }

//   .easy-pron {
//     font-size: 1.05rem;
//     font-weight: 600;
//     color: #fff;
//     margin-top: 0.2rem;
//   }

//   .ipa-row { display: flex; gap: 2rem; flex-wrap: wrap; }
//   .ipa-row .field { flex: 1; min-width: 150px; }

//   .ipa-value {
//     font-size: 0.9rem;
//     color: #b8c2ca;
//     margin-top: 0.2rem;
//   }

//   .meaning-value {
//     font-size: 0.88rem;
//     color: #b8c2ca;
//     margin-top: 0.2rem;
//     line-height: 1.55;
//   }

//   .replacement-box {
//     margin-top: 0.3rem;
//     background: rgba(0,255,156,0.06);
//     border: 1px solid rgba(0,255,156,0.25);
//     color: #b6ffe0;
//     border-radius: 6px;
//     padding: 0.6rem 0.8rem;
//     font-size: 0.85rem;
//     line-height: 1.5;
//   }

//   .replacement-box:before { content: "+ "; color: var(--accent); font-weight: 700; }

//   .tips { margin: 0.3rem 0 0; padding: 0; list-style: none; }

//   .tips li {
//     display: flex;
//     gap: 0.5rem;
//     font-size: 0.85rem;
//     color: #b8c2ca;
//     line-height: 1.55;
//     margin-top: 0.4rem;
//   }

//   .tips li:before { content: ">"; color: var(--accent2); flex-shrink: 0; }

//   .error {
//     color: var(--danger);
//     background: rgba(255,84,112,0.06);
//     border: 1px solid rgba(255,84,112,0.3);
//     padding: 0.8rem 1rem;
//     border-radius: 8px;
//     font-size: 0.85rem;
//   }

//   .empty {
//     color: var(--dim);
//     font-size: 0.85rem;
//     padding: 1rem 0;
//   }

//   footer {
//     margin-top: 3rem;
//     color: var(--dim);
//     font-size: 0.75rem;
//     border-top: 1px solid var(--line);
//     padding-top: 1.2rem;
//   }

//   footer a { color: var(--accent2); text-decoration: none; }
//   footer a:hover { text-decoration: underline; }

//   code {
//     background: #0a0d0f;
//     color: var(--accent);
//     padding: 0.12rem 0.4rem;
//     border-radius: 4px;
//     font-size: 0.8rem;
//     border: 1px solid var(--line);
//   }
// </style>
// </head>
// <body>
// <div class="wrap">

//   <div class="top-tag"><span class="pulse"></span> system online</div>
//   <h1>PronounceAI<span class="cursor">&nbsp;</span></h1>
//   <p class="sub">Paste a news script. We flag unfamiliar names and return presenter-ready pronunciation guidance.</p>

//   <div class="term">
//     <div class="term-bar">
//       <span></span><span></span><span></span>
//       <span class="label">script.txt</span>
//     </div>
//     <div class="term-body">
//       <div class="prompt-line">$ paste-script --analyze</div>
//       <textarea id="script" placeholder="President Volodymyr Zelenskyy met with NATO Secretary General Mark Rutte in The Hague."></textarea>
//     </div>
//     <div class="term-footer">
//       <span class="hint">ctrl not required, just click below</span>
//       <button id="btn" onclick="analyze()">Run analysis</button>
//     </div>
//   </div>

//   <div id="result"></div>

//   <footer>
//     health: <a href="/api/health">/api/health</a> &nbsp;&middot;&nbsp;
//     endpoint: <code>POST /api/pronounce</code>
//   </footer>

// </div>

// <script>
// function escapeHtml(str) {
//   if (str === undefined || str === null) return '';
//   return String(str)
//     .replace(/&/g, '&amp;')
//     .replace(/</g, '&lt;')
//     .replace(/>/g, '&gt;');
// }

// function difficultyClass(diff) {
//   var d = (diff || '').toLowerCase();
//   if (d === 'easy') return 'easy';
//   if (d === 'hard') return 'hard';
//   return 'medium';
// }

// async function analyze() {
//   var input = document.getElementById('script').value.trim();
//   var resultEl = document.getElementById('result');
//   var btn = document.getElementById('btn');

//   if (!input) {
//     resultEl.innerHTML = '<div class="error">error: script input is empty</div>';
//     return;
//   }

//   btn.disabled = true;
//   btn.textContent = 'Running...';
//   resultEl.innerHTML = '<div class="status-line"><span class="spinner"></span> analyzing script...</div>';

//   try {
//     var res = await fetch('/api/pronounce', {
//       method: 'POST',
//       headers: { 'Content-Type': 'application/json' },
//       body: JSON.stringify({ script: input })
//     });

//     var data = await res.json();
//     render(data);
//   } catch (err) {
//     resultEl.innerHTML = '<div class="error">request failed: ' + escapeHtml(err.message) + '</div>';
//   }

//   btn.disabled = false;
//   btn.textContent = 'Run analysis';
// }

// function render(data) {
//   var resultEl = document.getElementById('result');

//   if (!data || data.success === false) {
//     var msg = (data && data.message) ? data.message : 'something went wrong.';
//     resultEl.innerHTML = '<div class="error">' + escapeHtml(msg) + '</div>';
//     return;
//   }

//   if (!data.words || data.words.length === 0) {
//     resultEl.innerHTML = '<div class="empty">no unfamiliar names detected in this script.</div>';
//     return;
//   }

//   var html = '<div class="meta">provider <b>' + escapeHtml(data.provider) + '</b> &nbsp;&middot;&nbsp; ' +
//     '<b>' + data.count + '</b> word(s) found &nbsp;&middot;&nbsp; ' +
//     (data.fromCache ? 'from cache' : 'fresh analysis') + ' &nbsp;&middot;&nbsp; ' +
//     data.processingTimeMs + 'ms</div>';

//   data.words.forEach(function (w) {
//     var diffClass = difficultyClass(w.difficulty);

//     html += '<div class="card ' + diffClass + '">';

//     html += '<div class="card-head">';
//     html += '<div class="word-title">' + escapeHtml(w.word) + '</div>';
//     html += '<div class="tags">';
//     if (w.category) html += '<span class="tag">' + escapeHtml(w.category) + '</span>';
//     if (w.language) html += '<span class="tag">' + escapeHtml(w.language) + '</span>';
//     if (w.difficulty) html += '<span class="tag difficulty-' + diffClass + '">' + escapeHtml(w.difficulty) + '</span>';
//     html += '</div></div>';

//     html += '<div class="field">';
//     html += '<div class="k">easy_pronunciation</div>';
//     html += '<div class="easy-pron">' + escapeHtml(w.easyPronunciation) + '</div>';
//     html += '</div>';

//     html += '<div class="field ipa-row">';
//     html += '<div class="field"><div class="k">ipa_english</div><div class="ipa-value">' + escapeHtml(w.ipaEnglish) + '</div></div>';
//     if (w.ipaNative) {
//       html += '<div class="field"><div class="k">ipa_native</div><div class="ipa-value">' + escapeHtml(w.ipaNative) + '</div></div>';
//     }
//     html += '</div>';

//     if (w.meaning) {
//       html += '<div class="field">';
//       html += '<div class="k">meaning</div>';
//       html += '<div class="meaning-value">' + escapeHtml(w.meaning) + '</div>';
//       html += '</div>';
//     }

//     if (w.replacement) {
//       html += '<div class="field">';
//       html += '<div class="k">suggested_replacement</div>';
//       html += '<div class="replacement-box">' + escapeHtml(w.replacement) + '</div>';
//       html += '</div>';
//     }

//     if (w.presenterTips && w.presenterTips.length > 0) {
//       html += '<div class="field">';
//       html += '<div class="k">presenter_tips</div>';
//       html += '<ul class="tips">';
//       w.presenterTips.forEach(function (tip) {
//         html += '<li>' + escapeHtml(tip) + '</li>';
//       });
//       html += '</ul></div>';
//     }

//     html += '</div>';
//   });

//   resultEl.innerHTML = html;
// }
// </script>
// </body>
// </html>`
// 		return e.HTML(http.StatusOK, html)
// 	})

// 	// Friendly response if someone visits /api/pronounce directly in a browser (GET)
// 	se.Router.GET("/api/pronounce", func(e *core.RequestEvent) error {
// 		return e.JSON(http.StatusOK, map[string]any{
// 			"status":  "ok",
// 			"message": "This endpoint is live. Send a POST request with JSON body { \"script\": \"...\" } to analyze pronunciations, or use the form on the homepage.",
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
