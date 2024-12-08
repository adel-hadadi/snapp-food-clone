package app

import (
	"snapp-food/pkg/validate"

	"github.com/jmoiron/sqlx"
)

type Application struct {
	Handlers     Handlers
	Repositories Repositories
	Services     Services
}

func New(db *sqlx.DB, validator validate.Validator) *Application {
	app := new(Application)

	app.setupRepositories(db)

	app.setupServices()

	app.setupHandlers(validator)

	return app
}
