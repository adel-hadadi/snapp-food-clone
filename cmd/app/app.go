package app

import (
	"snapp-food/internal/delivery/http/handler"
	"snapp-food/pkg/validate"
)

type Application struct {
	Handlers Handlers
}

type Handlers struct {
	Auth handler.AuthHandler
}

func New(validator validate.Validator) Application {
	return Application{
		Handlers: Handlers{
			Auth: handler.NewAuthHandler(validator),
		},
	}
}
