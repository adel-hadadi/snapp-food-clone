package main

import (
	"fmt"
	"os"
	"snapp-food/cmd/app"
	"snapp-food/data/database"
	"snapp-food/internal/delivery/http"
	"snapp-food/pkg/logs"

	"github.com/joho/godotenv"
)

const ErrLoadConfigFile = "error on loading config file: %w"

func main() {
	if err := godotenv.Load(); err != nil {
		panic(fmt.Errorf(ErrLoadConfigFile, err))
	}

	logger := logs.NewLogger()
	defer logger.Sync()

	db := database.New()

	app := app.New(db)

	http.New(app.Handlers, app.Services.Token).
		Run(os.Getenv("PORT"))
}
