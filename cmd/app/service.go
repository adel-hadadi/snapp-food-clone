package app

import (
	"snapp-food/internal/adapters"
	authservice "snapp-food/internal/service/auth"
	productserviec "snapp-food/internal/service/auth/product"
	cityservice "snapp-food/internal/service/city"
	orderservice "snapp-food/internal/service/order"
	otpservice "snapp-food/internal/service/otp"
	productcategoryservice "snapp-food/internal/service/productcategory"
	provinceservice "snapp-food/internal/service/province"
	storeservice "snapp-food/internal/service/store"
	storecategoryservice "snapp-food/internal/service/storecategory"
	storemanagerservice "snapp-food/internal/service/storemanager"
	storetypeservice "snapp-food/internal/service/storetype"
	tokenservice "snapp-food/internal/service/token"
	userservice "snapp-food/internal/service/user"
	useraddressservice "snapp-food/internal/service/useraddress"
)

type Services struct {
	OTPService      otpservice.Service
	Token           tokenservice.Service
	Auth            authservice.Service
	User            userservice.Service
	StoreType       storetypeservice.Service
	UserAddress     useraddressservice.Service
	Store           storeservice.Service
	StoreManager    storemanagerservice.Service
	Product         productserviec.Service
	StoreCategory   storecategoryservice.Service
	ProductCategory productcategoryservice.Service
	Order           orderservice.Service
	Province        provinceservice.Service
	City            cityservice.Service
}

func (a *Application) setupServices() {
	notification := adapters.NewNotificationSMS()

	otpSvc := otpservice.New(notification, a.Repositories.OTPRepo)
	tokenSvc := tokenservice.New()

	a.Services = Services{
		OTPService:      otpSvc,
		Token:           tokenSvc,
		Auth:            authservice.New(otpSvc, a.Repositories.UserRepo, tokenSvc),
		User:            userservice.New(a.Repositories.UserRepo, a.Repositories.UserAddressRepo),
		StoreType:       storetypeservice.New(a.Repositories.StoreTypeRepo),
		UserAddress:     useraddressservice.New(a.Repositories.UserAddressRepo),
		Store:           storeservice.New(a.Repositories.StoreRepo, a.Repositories.StoreCategoryRepo),
		StoreManager:    storemanagerservice.New(otpSvc, tokenSvc, a.Repositories.StoreRepo),
		Product:         productserviec.New(a.Repositories.ProductRepo),
		StoreCategory:   storecategoryservice.New(a.Repositories.StoreCategoryRepo),
		ProductCategory: productcategoryservice.New(a.Repositories.ProductCategoryRepo),
		Order:           orderservice.New(a.Repositories.OrderRepo, a.Repositories.ProductRepo, a.Repositories.StoreRepo, a.Repositories.UserAddressRepo),
		Province:        provinceservice.New(a.Repositories.ProvinceRepo),
		City:            cityservice.New(a.Repositories.CityRepo),
	}
}
