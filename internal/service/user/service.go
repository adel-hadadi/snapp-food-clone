package userservice

import (
	"context"
	"snapp-food/internal/entity"
)

type Service struct {
	userRepo userRepository
}

type userRepository interface {
	Get(ctx context.Context, userID int) (entity.User, error)
	Update(ctx context.Context, userID int, firstName, lastName, nationalID string) error
}

func New(userRepo userRepository) Service {
	return Service{
		userRepo: userRepo,
	}
}
