package productserviec

import (
	"context"

	"snapp-food/pkg/apperr"
)

type ProductRes struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
	Slug        string  `json:"slug"`
	Rate        float32 `json:"rate"`
	Description string  `json:"description"`
	CategoryID  *int    `json:"category_id"`
	StoreID     int     `json:"store_id"`
	Status      int8    `json:"status"`

	Store StoreRes `json:"store,omitempty"`
}

type StoreRes struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (s Service) List(ctx context.Context, storeID int) ([]ProductRes, error) {
	const getListOfProductsSysMsg = "product service get list by store id"

	products, err := s.repo.GetByStoreID(ctx, storeID)
	if err != nil {
		return nil, apperr.New(apperr.Unexpected).
			WithErr(err).
			WithSysMsg(getListOfProductsSysMsg)
	}

	productRes := make([]ProductRes, 0, len(products))

	for i := range products {
		productRes = append(productRes, ProductRes{
			ID:         products[i].ID,
			Name:       products[i].Name,
			Price:      float64(products[i].Price / 10),
			Image:      products[i].Image,
			Slug:       products[i].Slug,
			Rate:       products[i].Rate,
			CategoryID: products[i].CategoryID,
			StoreID:    products[i].StoreID,
			Status:     products[i].Status,
			Store: StoreRes{
				ID:   products[i].Store.ID,
				Name: products[i].Store.Name,
				Slug: products[i].Store.Slug,
			},
		})
	}

	return productRes, nil
}

type Filters struct {
	Sort []string
}

func (s Service) FilteredList(ctx context.Context, userID int, req Filters) ([]ProductRes, error) {
	const getListOfProductsByUserIDSysMSG = "product service filteredList"

	products, err := s.repo.GetByUserID(ctx, userID, req.Sort)
	if err != nil {
		return nil, apperr.New(apperr.Unexpected).WithErr(err).WithSysMsg(getListOfProductsByUserIDSysMSG)
	}

	productRes := make([]ProductRes, 0, len(products))
	for _, product := range products {
		productRes = append(productRes, ProductRes{
			ID:         product.ID,
			Name:       product.Name,
			Price:      float64(product.Price / 10),
			Image:      product.Image,
			Slug:       product.Slug,
			Rate:       product.Rate,
			CategoryID: product.CategoryID,
			StoreID:    product.StoreID,
			Status:     product.Status,
			Store: StoreRes{
				ID:   product.Store.ID,
				Name: product.Store.Name,
				Slug: product.Store.Slug,
			},
		})
	}

	return productRes, nil
}
