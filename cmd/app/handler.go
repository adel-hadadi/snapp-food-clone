package app

import (
	"snapp-food/internal/delivery/http/handler"
)

type Handlers struct {
	Auth            handler.AuthHandler
	OTP             handler.OTPHandler
	Profile         handler.ProfileHandler
	StoreType       handler.StoreTypeHandler
	Store           handler.StoreHandler
	StoreManager    handler.StoreManagerHandler
	Product         handler.ProductHandler
	StoreCategory   handler.StoreCategoryHandler
	ProductCategory handler.ProductCategoryHandler
	Order           handler.OrderHandler
	Province        handler.ProvinceHandler
	City            handler.CityHandler
}

func (a *Application) setupHandlers() {
	a.Handlers = Handlers{
		Auth:            handler.NewAuthHandler(a.Validations.Auth, a.Services.Auth),
		OTP:             handler.NewOTPHandler(a.Services.OTPService),
		Profile:         handler.NewProfileHandler(a.Services.User, a.Services.UserAddress),
		StoreType:       handler.NewStoreTypeHandler(a.Services.StoreType),
		Store:           handler.NewStoreHandler(a.Services.Store),
		StoreManager:    handler.NewStoreManagerHandler(a.Services.Store, a.Services.OTPService, a.Services.StoreManager),
		Product:         handler.NewProductHandler(a.Services.Product),
		StoreCategory:   handler.NewStoreCategoryHandler(a.Services.StoreCategory),
		ProductCategory: handler.NewProductCategoryHandler(a.Services.ProductCategory),
		Order:           handler.NewOrderHandler(a.Services.Order),
		Province:        handler.NewProvinceHandler(a.Services.Province),
		City:            handler.NewCityHandler(a.Services.City),
	}
}
