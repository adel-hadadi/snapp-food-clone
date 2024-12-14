package storeservice

import (
	"context"
	"time"

	"snapp-food/pkg/apperr"
)

type StoreRes struct {
	ID               int
	Name             string
	Slug             string
	ManagerFirstName string
	ManagerLastName  string
	Phone            string
	Address          string
	Latitude         float64
	Longitude        float64
	Logo             string
	StoreTypeID      int
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (s Service) Find(ctx context.Context, slug string) (StoreRes, error) {
	const findStoreBySlugSysMsg = "store service find by slug"

	store, err := s.repo.FindBySlug(ctx, slug)
	if err != nil {
		if apperr.IsSQLNoRows(err) {
			return StoreRes{}, apperr.New(apperr.NotFound)
		}

		return StoreRes{}, apperr.New(apperr.Unexpected).WithErr(err).WithSysMsg(findStoreBySlugSysMsg)
	}

	return StoreRes{
		ID:               store.ID,
		Name:             store.Name,
		Slug:             store.Slug,
		ManagerFirstName: store.ManagerFirstName,
		ManagerLastName:  store.ManagerLastName,
		Phone:            store.Phone,
		Address:          store.Address,
		Latitude:         store.Latitude,
		Longitude:        store.Longitude,
		Logo:             store.Logo,
		StoreTypeID:      store.StoreTypeID,
		CreatedAt:        store.CreatedAt,
		UpdatedAt:        store.UpdatedAt,
	}, nil
}

func (s Service) List(ctx context.Context) ([]StoreRes, error) {
	const getListOfStoresSysMsg = "store service get list of stores"

	stores, err := s.repo.Get(ctx)
	if err != nil {
		return nil, apperr.New(apperr.Unexpected).WithErr(err).WithSysMsg(getListOfStoresSysMsg)
	}

	res := make([]StoreRes, 0, len(stores))
	for i := range stores {
		res = append(res, StoreRes{
			ID:               stores[i].ID,
			Name:             stores[i].Name,
			Slug:             stores[i].Slug,
			ManagerFirstName: stores[i].ManagerFirstName,
			ManagerLastName:  stores[i].ManagerLastName,
			Phone:            stores[i].Phone,
			Address:          stores[i].Address,
			Latitude:         stores[i].Latitude,
			Longitude:        stores[i].Longitude,
			Logo:             stores[i].Logo,
			StoreTypeID:      stores[i].StoreTypeID,
			CreatedAt:        stores[i].CreatedAt,
			UpdatedAt:        stores[i].UpdatedAt,
		})
	}

	return res, nil
}
