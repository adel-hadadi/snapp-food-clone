package storeservice

import (
	"context"

	"snapp-food/internal/dto"
	"snapp-food/internal/entity"
	"snapp-food/pkg/apperr"

	"github.com/gosimple/slug"
)

func (s Service) Create(ctx context.Context, req dto.StoreCreateReq) error {
	const createStoreSysMsg = "store service create store"

	err := s.repo.Create(ctx, entity.Store{
		Name:        req.Name,
		Slug:        slug.MakeLang(req.Name, "fa"),
		Latitude:    req.Latitude,
		Longitude:   req.Longitude,
		Logo:        req.Logo,
		StoreTypeID: req.StoreTypeID,
		CityID:      &req.CityID,
		ManagerID:   req.ManagerID,
	})
	if err != nil {
		return apperr.New(apperr.Unexpected).WithErr(err).WithSysMsg(createStoreSysMsg)
	}

	return nil
}
