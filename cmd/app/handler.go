package app

import (
	"snapp-food/internal/delivery/http/handler"
	"snapp-food/pkg/validate"
)

type Handlers struct {
	Auth      handler.AuthHandler
	OTP       handler.OTPHandler
	Profile   handler.ProfileHandler
	StoreType handler.StoreTypeHandler
	Store     handler.StoreHandler
}

func (a *Application) setupHandlers(v validate.Validator) {
	a.Handlers = Handlers{
		Auth:      handler.NewAuthHandler(v, a.Services.Auth),
		OTP:       handler.NewOTPHandler(v, a.Services.OTPService),
		Profile:   handler.NewProfileHandler(a.Services.User, a.Services.UserAddress),
		StoreType: handler.NewStoreTypeHandler(a.Services.StoreType),
		Store:     handler.NewStoreHandler(a.Services.Store),
	}
}
