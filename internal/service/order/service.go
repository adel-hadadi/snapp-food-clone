package orderservice

import (
	"context"
	"snapp-food/internal/entity"
)

type Service struct {
	repo            orderRepository
	productRepo     productRepository
	storeRepo       storeRepository
	userAddressRepo userAddressRepository
}

type orderRepository interface {
	Create(ctx context.Context, userID, userAddressID, storeID, totalAmount int, items []entity.OrderItem) error
	GetByUserID(ctx context.Context, userID int) ([]entity.Order, error)
	GetByID(ctx context.Context, orderID int) (entity.Order, error)
	UpdateStatus(ctx context.Context, orderID int, status int8) error
	RemovePending(ctx context.Context) error
}

type productRepository interface {
	Find(ctx context.Context, id int) (entity.Product, error)
}

type userAddressRepository interface {
	GetUserDefaultAddress(ctx context.Context, userID int) (entity.UserAddress, error)
}

type storeRepository interface {
	FindBySlug(ctx context.Context, slug string) (entity.Store, error)
}

func New(orderRepo orderRepository, productRepo productRepository, storeRepo storeRepository, userAddressRepo userAddressRepository) Service {
	return Service{
		repo:            orderRepo,
		productRepo:     productRepo,
		storeRepo:       storeRepo,
		userAddressRepo: userAddressRepo,
	}
}
