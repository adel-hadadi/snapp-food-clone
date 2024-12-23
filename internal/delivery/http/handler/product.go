package handler

import (
	"context"
	"net/http"

	"snapp-food/internal/delivery/http/middleware"
	productservice "snapp-food/internal/service/auth/product"
	"snapp-food/pkg/httpres"
	"snapp-food/pkg/server/httpreq"
)

type ProductHandler struct {
	productSvc productService
}

type productService interface {
	Create(ctx context.Context, storeID int, req productservice.CreateProductReq) error
	List(ctx context.Context, storeID int) ([]productservice.ProductRes, error)
	FilteredList(ctx context.Context, authID int, req productservice.Filters) ([]productservice.ProductRes, error)
}

func NewProductHandler(service productService) ProductHandler {
	return ProductHandler{
		productSvc: service,
	}
}

type CreateProductReq struct {
	Name       string  `json:"name"`
	Image      string  `json:"image"`
	Price      float64 `json:"price"`
	CategoryID int     `json:"category_id"`
}

func (h ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	req, err := httpreq.Bind[CreateProductReq](r)
	if err != nil {
		httpres.ValidationErr(w, err, http.StatusBadRequest)
		return
	}

	storeID := r.Context().Value(middleware.StoreCtxKey).(float64)

	if err := h.productSvc.Create(r.Context(), int(storeID), productservice.CreateProductReq{
		Name:       req.Name,
		Image:      req.Image,
		CategoryID: req.CategoryID,
		Price:      req.Price,
	}); err != nil {
		httpres.WithErr(w, err)
		return
	}

	httpres.Success(w, nil, http.StatusCreated)
}

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

func (h ProductHandler) List(w http.ResponseWriter, r *http.Request) {
	storeID := r.Context().Value(middleware.StoreCtxKey).(float64)

	products, err := h.productSvc.List(r.Context(), int(storeID))
	if err != nil {
		httpres.WithErr(w, err)
		return
	}

	res := make([]ProductRes, 0, len(products))

	for i := range products {
		res = append(res, ProductRes{
			ID:         products[i].ID,
			Name:       products[i].Name,
			Price:      products[i].Price,
			Image:      products[i].Image,
			Slug:       products[i].Slug,
			Rate:       products[i].Rate,
			CategoryID: products[i].CategoryID,
			StoreID:    products[i].StoreID,
			Status:     products[i].Status,
		})
	}

	httpres.Success(w, res, http.StatusOK)
}

func (h ProductHandler) FilteredList(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	products, err := h.productSvc.FilteredList(
		r.Context(),
		httpreq.AuthID(r),
		productservice.Filters{
			Sort: queryParams["sort"],
		})

	if err != nil {
		httpres.WithErr(w, err)
		return
	}

	httpres.Success(w, products, http.StatusOK)
}
