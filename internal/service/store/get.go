package storeservice

import (
	"context"
	"snapp-food/pkg/convert"
	"time"

	productserviec "snapp-food/internal/service/auth/product"
	"snapp-food/pkg/apperr"
)

type StoreRes struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	Slug             string    `json:"slug"`
	ManagerFirstName string    `json:"manager_first_name"`
	ManagerLastName  string    `json:"manager_last_name"`
	Phone            string    `json:"phone"`
	Address          string    `json:"address"`
	Latitude         float64   `json:"latitude"`
	Longitude        float64   `json:"longitude"`
	Logo             string    `json:"logo"`
	StoreTypeID      int       `json:"store_type_id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`

	Categories []CategoryRes `json:"categories"`
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
		})
	}

	return storeRes, nil
}
