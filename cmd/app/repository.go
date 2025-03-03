package app

import (
	"snapp-food/internal/repository"

	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	OTPRepo             repository.OTPRepository
	UserRepo            repository.UserRepository
	StoreTypeRepo       repository.StoreTypeRepository
	UserAddressRepo     repository.UserAddressRepository
	StoreRepo           repository.StoreRepository
	ProductRepo         repository.ProductRepository
	StoreCategoryRepo   repository.StoreCategoryRepository
	ProductCategoryRepo repository.ProductCategoryRepository
	OrderRepo           repository.OrderRepository
	ProvinceRepo        repository.ProvinceRepository
	CityRepo            repository.CityRepository
	TokenRepo           repository.TokenRepository
}

func (a *Application) setupRepositories(db *sqlx.DB) {
	a.Repositories = Repositories{
		OTPRepo:             repository.NewOTPRepository(db),
		UserRepo:            repository.NewUserRepository(db),
		StoreTypeRepo:       repository.NewStoreTypeRepository(db),
		UserAddressRepo:     repository.NewUserAddressRepository(db),
		StoreRepo:           repository.NewStoreRepository(db),
		ProductRepo:         repository.NewProductRepository(db),
		StoreCategoryRepo:   repository.NewStoreCategoryRepository(db),
		ProductCategoryRepo: repository.NewProductCategoryRepository(db),
		OrderRepo:           repository.NewOrderRepository(db),
		ProvinceRepo:        repository.NewProvinceRepository(db),
		CityRepo:            repository.NewCityRepository(db),
		TokenRepo:           repository.NewTokenRepository(db),
	}
}
