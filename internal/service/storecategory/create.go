package storecategoryservice

import (
	"context"

	"snapp-food/internal/entity"
	"snapp-food/pkg/apperr"
)

func (s Service) Create(ctx context.Context, storeID int, name string) error {
	if err := s.repo.Create(ctx, entity.StoreCategory{
		Name:    name,
		StoreID: storeID,
	}); err != nil {
		return apperr.New(apperr.Unexpected).WithErr(err)
	}

	return nil
}
