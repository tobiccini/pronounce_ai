// package main

// import (
// 	"log"
// 	"myap/pb_hooks/api"
// 	"os"

// 	"github.com/pocketbase/pocketbase"
// 	"github.com/pocketbase/pocketbase/apis"
// 	"github.com/pocketbase/pocketbase/core"
// )

// func main() {
//     app := pocketbase.New()

//     app.OnServe().BindFunc(func(se *core.ServeEvent) error {
//         // serves static files from the provided public dir (if exists)
//         se.Router.GET("/{path...}", apis.Static(os.DirFS("./pb_public"), false))

//         api.RegisterPronunciationRoutes(se)

//         return se.Next()
//     })

//     if err := app.Start(); err != nil {
//         log.Fatal(err)
//     }
// }



package main

import (
	"log"

	"myap/pb_hooks/api"

	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Falling back to system environment variables.")
	}

	app := pocketbase.New()

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {

		api.RegisterRoutes(se)

		return se.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}









// {
//   "script": "President Volodymyr Zelenskyy met with NATO Secretary General Mark Rutte in The Hague."
// }


// About the config package

// When I suggested:

// package config

// var GeminiAPIKey = os.Getenv("GEMINI_API_KEY")

// I meant creating a new package like this:

// project/
// │
// ├── config/
// │   └── config.go
// │
// ├── pb_hooks/
// │
// ├── main.go
// │
// └── go.mod

// with:

// package config

// import "os"

// var GeminiAPIKey = os.Getenv("GEMINI_API_KEY")

// Then in pronounce.go you'd write:

// provider := &ai.GeminiProvider{
//     APIKey: config.GeminiAPIKey,
// }

// instead of:

// provider := &ai.GeminiProvider{
//     APIKey: os.Getenv("GEMINI_API_KEY"),