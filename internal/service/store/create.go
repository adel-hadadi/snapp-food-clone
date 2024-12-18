package storeservice

import (
	"context"

	"github.com/gosimple/slug"
	"snapp-food/internal/entity"
	"snapp-food/pkg/apperr"
)

type CreateReq struct {
	Name             string
	Phone            string
	ManagerFirstName string
	ManagerLastName  string
	Address          string
	StoreTypeID      int
	Latitude         float64
	Longitude        float64
	Logo             string // TODO: change logo to file
	CityID           int
}

func (s Service) Create(ctx context.Context, req CreateReq) error {
	const createStoreSysMsg = "store service create store"

	err := s.repo.Create(ctx, entity.Store{
		Name:             req.Name,
		Slug:             slug.MakeLang(req.Name, "fa"),
		ManagerFirstName: req.ManagerFirstName,
		ManagerLastName:  req.ManagerLastName,
		Phone:            req.Phone,
		Latitude:         req.Latitude,
		Longitude:        req.Longitude,
		Logo:             req.Logo,
		StoreTypeID:      req.StoreTypeID,
		CityID:           &req.CityID,
	})
	if err != nil {
		return apperr.New(apperr.Unexpected).WithErr(err).WithSysMsg(createStoreSysMsg)
	}

	return nil
}
