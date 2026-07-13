package main

import (
	"log"

	"myap/pb_hooks/api"

	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/jsvm"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Falling back to system environment variables.")
	}

	app := pocketbase.New()

	jsvm.MustRegister(app, jsvm.Config{
		MigrationsDir: "pb_migrations",
	})

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {

		api.RegisterRoutes(se)

		return se.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}


// package main

// import (
// 	"log"
// 	"os"

// 	"myap/pb_hooks/api"

// 	"github.com/joho/godotenv"
// 	"github.com/pocketbase/pocketbase"
// 	"github.com/pocketbase/pocketbase/core"
// 	"github.com/pocketbase/pocketbase/plugins/migratecmd"
// )


// func main() {

// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Println("No .env file found. Falling back to system environment variables.")
// 	}

// 	app := pocketbase.New()

// 	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
// 		// Automigrate keeps pb_migrations/ in sync with schema changes made in the dashboard
// 		// during local dev; safe to leave true, but data only comes from committed migration files
// 		Automigrate: os.Getenv("ENVIRONMENT") != "production",
// 	})

// 	app.OnServe().BindFunc(func(se *core.ServeEvent) error {

// 		api.RegisterRoutes(se)

// 		return se.Next()
// 	})

// 	if err := app.Start(); err != nil {
// 		log.Fatal(err)
// 	}
// }

// func main() {

// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Println("No .env file found. Falling back to system environment variables.")
// 	}

// 	app := pocketbase.New()

// 	app.OnServe().BindFunc(func(se *core.ServeEvent) error {

// 		api.RegisterRoutes(se)

// 		return se.Next()
// 	})

// 	if err := app.Start(); err != nil {
// 		log.Fatal(err)
// 	}
// }
