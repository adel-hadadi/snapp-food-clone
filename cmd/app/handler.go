package app

import (
	"snapp-food/internal/delivery/http/handler"
	"snapp-food/internal/delivery/http/handler/authhandler"
	"snapp-food/internal/delivery/http/handler/sellerstorehandler"
)

type Handlers struct {
	Auth            authhandler.Handler
	Profile         handler.ProfileHandler
	StoreType       handler.StoreTypeHandler
	Store           handler.StoreHandler
	Product         handler.ProductHandler
	StoreCategory   handler.StoreCategoryHandler
	ProductCategory handler.ProductCategoryHandler
	Order           handler.OrderHandler
	Province        handler.ProvinceHandler
	City            handler.CityHandler
	SellerStore     sellerstorehandler.Handler
}

func (a *Application) setupHandlers() {
	a.Handlers = Handlers{
		Auth:            authhandler.New(a.Services.Auth),
		Profile:         handler.NewProfileHandler(a.Services.User, a.Services.UserAddress),
		StoreType:       handler.NewStoreTypeHandler(a.Services.StoreType),
		Store:           handler.NewStoreHandler(a.Services.Store),
		Product:         handler.NewProductHandler(a.Services.Product),
		StoreCategory:   handler.NewStoreCategoryHandler(a.Services.StoreCategory),
		ProductCategory: handler.NewProductCategoryHandler(a.Services.ProductCategory),
		Order:           handler.NewOrderHandler(a.Services.Order),
		Province:        handler.NewProvinceHandler(a.Services.Province),
		City:            handler.NewCityHandler(a.Services.City),
		SellerStore:     sellerstorehandler.New(a.Services.Store),
	}
}
