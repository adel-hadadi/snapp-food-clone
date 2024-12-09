package app

import (
	"snapp-food/internal/adapters"
	authservice "snapp-food/internal/service/auth"
	otpservice "snapp-food/internal/service/otp"
	storetypeservice "snapp-food/internal/service/storetype"
	tokenservice "snapp-food/internal/service/token"
	userservice "snapp-food/internal/service/user"
	useraddressservice "snapp-food/internal/service/useraddress"
)

type Services struct {
	OTPService  otpservice.Service
	Token       tokenservice.Service
	Auth        authservice.Service
	User        userservice.Service
	StoreType   storetypeservice.Service
	UserAddress useraddressservice.Service
}

func (a *Application) setupServices() {
	notification := adapters.NewNotificationSMS()

	otpSvc := otpservice.New(notification, a.Repositories.OTPRepo)
	tokenSvc := tokenservice.New()

	a.Services = Services{
		OTPService:  otpSvc,
		Token:       tokenSvc,
		Auth:        authservice.New(otpSvc, a.Repositories.UserRepo, tokenSvc),
		User:        userservice.New(a.Repositories.UserRepo),
		StoreType:   storetypeservice.New(a.Repositories.StoreTypeRepo),
		UserAddress: useraddressservice.New(a.Repositories.UserAddressRepo),
	}
}
