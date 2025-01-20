package storeservice

import (
	"context"
	"snapp-food/pkg/convert"
	"time"

	"snapp-food/internal/dto"
	productserviec "snapp-food/internal/service/auth/product"
	storetypeservice "snapp-food/internal/service/storetype"
	"snapp-food/pkg/apperr"
)

type StoreRes struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Rate        float32   `json:"rate"`
	Address     string    `json:"address"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	Logo        string    `json:"logo"`
	StoreTypeID int       `json:"store_type_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	Categories []CategoryRes              `json:"categories"`
	StoreType  storetypeservice.StoreType `json:"store_type"`
}

type CategoryRes struct {
	ID       int                         `json:"id"`
	Name     string                      `json:"name"`
	Products []productserviec.ProductRes `json:"products"`
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

	storeRes := StoreRes{
		ID:          store.ID,
		Name:        store.Name,
		Slug:        store.Slug,
		Rate:        store.Rate,
		Address:     store.Address,
		Latitude:    store.Latitude,
		Longitude:   store.Longitude,
		Logo:        store.Logo,
		StoreTypeID: store.StoreTypeID,
		CreatedAt:   store.CreatedAt,
		UpdatedAt:   store.UpdatedAt,

		StoreType: storetypeservice.StoreType{
			ID:    store.StoreType.ID,
			Name:  store.StoreType.Name,
			Image: store.StoreType.Image,
			URL:   store.StoreType.URL,
		},
	}

	const getStoreCategoriesSysMSG = "store service get store's categories"

	storeCategories, err := s.storeCategoryRepo.GetByStoreID(ctx, store.ID)
	if err != nil {
		return StoreRes{}, apperr.New(apperr.Unexpected).WithErr(err).WithSysMsg(getStoreCategoriesSysMSG)
	}

	categories := make([]CategoryRes, 0, len(storeCategories))
	for _, category := range storeCategories {
		c, _ := convert.ToStruct[CategoryRes](category)

		categories = append(categories, c)
	}

	storeRes.Categories = categories

	return storeRes, nil
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
			ID:          stores[i].ID,
			Name:        stores[i].Name,
			Slug:        stores[i].Slug,
			Address:     stores[i].Address,
			Latitude:    stores[i].Latitude,
			Longitude:   stores[i].Longitude,
			Logo:        stores[i].Logo,
			StoreTypeID: stores[i].StoreTypeID,
			CreatedAt:   stores[i].CreatedAt,
			UpdatedAt:   stores[i].UpdatedAt,

			StoreType: storetypeservice.StoreType{
				ID:    stores[i].StoreType.ID,
				Name:  stores[i].StoreType.Name,
				Image: stores[i].StoreType.Image,
				URL:   stores[i].StoreType.URL,
			},
		})
	}

	return res, nil
}

func (s Service) Nearest(ctx context.Context, userID int) ([]StoreRes, error) {
	const findNearestStores = "store service nearest method"

	stores, err := s.repo.Nearest(ctx, userID)
	if err != nil {
		return nil, apperr.New(apperr.Unexpected).WithErr(err).WithSysMsg(findNearestStores)
	}

	storeRes := make([]StoreRes, 0, len(stores))

	for _, store := range stores {
		storeRes = append(storeRes, StoreRes{
			ID:      store.ID,
			Name:    store.Name,
			Slug:    store.Slug,
			Address: store.Address,
			Logo:    store.Logo,
			StoreType: storetypeservice.StoreType{
				ID:    store.StoreType.ID,
				Name:  store.StoreType.Name,
				Image: store.StoreType.Image,
				URL:   store.StoreType.URL,
			},
		})
	}

	return storeRes, nil
}

func (s Service) ListByProductCategory(ctx context.Context, userID int, productCategorySlug string) ([]StoreRes, error) {
	const getListOfStoresSysMsg = "store service get list of stores"

	stores, err := s.repo.GetByProductCategory(ctx, userID, productCategorySlug)
	if err != nil {
		return nil, apperr.New(apperr.Unexpected).WithErr(err).WithSysMsg(getListOfStoresSysMsg)
	}

	res := make([]StoreRes, 0, len(stores))
	for i := range stores {
		res = append(res, StoreRes{
			ID:          stores[i].ID,
			Name:        stores[i].Name,
			Slug:        stores[i].Slug,
			Address:     stores[i].Address,
			Latitude:    stores[i].Latitude,
			Longitude:   stores[i].Longitude,
			Logo:        stores[i].Logo,
			StoreTypeID: stores[i].StoreTypeID,
			CreatedAt:   stores[i].CreatedAt,
			UpdatedAt:   stores[i].UpdatedAt,

			StoreType: storetypeservice.StoreType{
				ID:    stores[i].StoreType.ID,
				Name:  stores[i].StoreType.Name,
				Image: stores[i].StoreType.Image,
				URL:   stores[i].StoreType.URL,
			},
		})
	}

	return res, nil
}

func (s Service) ListByManagerID(ctx context.Context, managerID int) ([]dto.StoreRes, error) {
	stores, err := s.repo.GetByManagerID(ctx, managerID)
	if err != nil {
		return nil, apperr.New(apperr.Unexpected).WithErr(err).WithSysMsg("retrive-list-of-stores")
	}

	storeRes := make([]dto.StoreRes, 0, len(stores))
	for _, store := range stores {
		storeRes = append(storeRes, dto.StoreRes{
			ID:   store.ID,
			Name: store.Name,
			Logo: store.Logo,
			Slug: store.Slug,
		})
	}

	return storeRes, nil
}
