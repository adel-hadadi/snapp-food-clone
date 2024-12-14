package storeservice

import (
	"context"

	"snapp-food/internal/entity"
)

type Service struct {
	repo storeRepository
}

type storeRepository interface {
	Find(ctx context.Context, id int) (entity.Store, error)
	FindBySlug(ctx context.Context, slug string) (entity.Store, error)
	Create(ctx context.Context, store entity.Store) error
	Get(ctx context.Context) ([]entity.Store, error)
}

func New(repo storeRepository) Service {
	return Service{repo: repo}
}
