package userservice

import (
	"context"
	"snapp-food/internal/entity"
)

type Service struct {
	userRepo        userRepository
	userAddressRepo userAddressRepository
}

type userRepository interface {
	Get(ctx context.Context, userID int) (entity.User, error)
	Update(ctx context.Context, userID int, firstName, lastName, nationalID string, defaultAddress int) error
}

type userAddressRepository interface {
	BelongsToUser(ctx context.Context, addressID, userID int) (bool, error)
}

func New(userRepo userRepository, userAddressRepo userAddressRepository) Service {
	return Service{
		userRepo:        userRepo,
		userAddressRepo: userAddressRepo,
	}
}
