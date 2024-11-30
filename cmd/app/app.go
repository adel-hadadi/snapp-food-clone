package app

type Application struct {
	Services Services
}

type Services struct {
}

func New() Application {
	return Application{}
}
