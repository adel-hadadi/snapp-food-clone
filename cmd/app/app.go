package app

import "snapp-food/internal/delivery/http/handler"

type Application struct {
	Handlers Handlers
}

type Handlers struct {
	Auth handler.AuthHandler
}

func New() Application {
	return Application{
		Handlers: Handlers{
			Auth: handler.NewAuthHandler(),
		},
	}
}
