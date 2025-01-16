package app

import (
	"github.com/jmoiron/sqlx"
)

type Application struct {
	Handlers     Handlers
	Repositories Repositories
	Services     Services
}

func New(db *sqlx.DB) *Application {
	app := new(Application)

	app.setupRepositories(db)

	app.setupServices()

	app.setupHandlers()

	return app
}
