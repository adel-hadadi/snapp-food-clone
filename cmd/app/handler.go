package app

import (
	"snapp-food/internal/delivery/http/handler"
	"snapp-food/pkg/validate"
)

type Handlers struct {
	Auth    handler.AuthHandler
	OTP     handler.OTPHandler
	Profile handler.ProfileHandler
}

func (a *Application) setupHandlers(v validate.Validator) {
	a.Handlers = Handlers{
		Auth:    handler.NewAuthHandler(v, a.Services.Auth),
		OTP:     handler.NewOTPHandler(v, a.Services.OTPService),
		Profile: handler.NewProfileHandler(),
	}
}
