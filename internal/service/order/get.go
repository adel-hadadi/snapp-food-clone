package orderservice

import (
	"context"
	productserviec "snapp-food/internal/service/auth/product"
	storeservice "snapp-food/internal/service/store"
	"snapp-food/pkg/apperr"
	"time"
)

type OrderRes struct {
	ID            int                   `json:"id"`
	StoreID       int                   `json:"store_id"`
	UserID        int                   `json:"user_id"`
	UserAddressID int                   `json:"user_address_id"`
	Amount        int                   `json:"amount"`
	CreatedAt     time.Time             `json:"created_at"`
	Store         storeservice.StoreRes `json:"store"`
	Status        int8                  `json:"status"`
	StatusLabel   string                `json:"status_label"`
	StatusLabelFa string                `json:"status_label_fa"`
	Items         []OrderItemRes        `json:"items"`
}

type OrderItemRes struct {
	ID        int     `json:"id"`
	OrderID   int     `json:"order_id"`
	ProductID int     `json:"product_id"`
	Quantity  int8    `json:"quantity"`
	Price     float64 `json:"price"`

	Product productserviec.ProductRes `json:"product"`
}

func (s Service) List(ctx context.Context, userID int) ([]OrderRes, error) {
	const getListOfOrdersSysMSG = "order service get list of orders"

	orders, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, apperr.New(apperr.Unexpected).WithErr(err).WithSysMsg(getListOfOrdersSysMSG)
	}

	const convertOrdersSysMSG = "order service cnvert orders"

	orderRes := make([]OrderRes, 0, len(orders))
	for _, order := range orders {
		o := OrderRes{
			ID:            order.ID,
			StoreID:       order.StoreID,
			UserID:        order.UserID,
			UserAddressID: order.UserAddressID,
			Amount:        order.Amount,
			CreatedAt:     order.CreatedAt,
			Store: storeservice.StoreRes{
				ID:      order.Store.ID,
				Name:    order.Store.Name,
				Slug:    order.Store.Slug,
				Rate:    order.Store.Rate,
				Address: order.Store.Address,
				Logo:    order.Store.Logo,
			},
			Status:        order.Status,
			StatusLabel:   order.StatusLabel().Label,
			StatusLabelFa: order.StatusLabel().LabelFa,
			Items:         []OrderItemRes{},
		}

		for _, item := range order.Items {
			o.Items = append(o.Items, OrderItemRes{
				ID:        item.ID,
				OrderID:   item.OrderID,
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
				Price:     float64(item.Price),
				Product: productserviec.ProductRes{
					ID:   item.Product.ID,
					Name: item.Product.Name,
					Slug: item.Product.Slug,
				},
			})
		}

		orderRes = append(orderRes, o)
	}

	return orderRes, nil
}
