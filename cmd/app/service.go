package app

import (
	"snapp-food/internal/adapters"
	authservice "snapp-food/internal/service/auth"
	otpservice "snapp-food/internal/service/otp"
	tokenservice "snapp-food/internal/service/token"
)

type Services struct {
	OTPService otpservice.Service
	Token      tokenservice.Service
	Auth       authservice.Service
}

func (a *Application) setupServices() {
	notification := adapters.NewNotificationSMS()

	otpSvc := otpservice.New(notification, a.Repositories.OTPRepo)
	tokenSvc := tokenservice.New()

	a.Services = Services{
		OTPService: otpSvc,
		Token:      tokenSvc,
		Auth:       authservice.New(otpSvc, a.Repositories.UserRepo, tokenSvc),
	}
}
