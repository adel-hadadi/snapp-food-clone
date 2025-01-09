package app

import (
	"github.com/jmoiron/sqlx"
	"github.com/rezakhademix/govalidator/v2"
)

type Application struct {
	Handlers     Handlers
	Repositories Repositories
	Services     Services
	Validations  Validations
}

func New(db *sqlx.DB, validator govalidator.Validator) *Application {
	app := new(Application)

	app.setupRepositories(db)

	app.setupServices()

	app.setupValidations(validator)

	app.setupHandlers()

	return app
}
