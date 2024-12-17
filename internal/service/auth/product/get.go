package productserviec

import (
	"context"

	"snapp-food/pkg/apperr"
)

type ProductRes struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Image      string  `json:"image"`
	Slug       string  `json:"slug"`
	Rate       float32 `json:"rate"`
	CategoryID int     `json:"category_id"`
	StoreID    int     `json:"store_id"`
	Status     int8    `json:"status"`
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
		})
	}

	return productRes, nil
}
