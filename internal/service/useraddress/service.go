package useraddressservice

import (
	"context"

	"snapp-food/internal/entity"
)

type Service struct {
	repo userAddressRepository
}

type userAddressRepository interface {
	GetByUserID(ctx context.Context, userID int) ([]entity.UserAddress, error)
	Create(ctx context.Context, userID int, address entity.UserAddress) error
}

func New(repo userAddressRepository) Service {
	return Service{
		repo: repo,
	}
}
