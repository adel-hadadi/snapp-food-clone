package main

import (
	"fmt"
	"os"
	"snapp-food/cmd/app"
	"snapp-food/internal/delivery/http"

	"github.com/joho/godotenv"
)

const ErrLoadConfigFile = "error on loading config file: %w"

func main() {
	if err := godotenv.Load(); err != nil {
		panic(fmt.Errorf(ErrLoadConfigFile, err))
	}

	app := app.New()

	http.New(app.Handlers).Run(os.Getenv("PORT"))
}
