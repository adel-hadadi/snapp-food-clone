package handler

import (
	"context"
	"net/http"

	"snapp-food/internal/delivery/http/middleware"
	"snapp-food/pkg/httpres"
	"snapp-food/pkg/server/httpreq"
)

type StoreCategoryHandler struct {
	storeCategorySvc storeCategoryService
}

type storeCategoryService interface {
	Create(ctx context.Context, storeID int, name string) error
	// List(ctx context.Context) ([]storecategoryservice.StoreCategoryRes, error)
}

func NewStoreCategoryHandler(service storeCategoryService) StoreCategoryHandler {
	return StoreCategoryHandler{storeCategorySvc: service}
}

// func (h StoreCategoryHandler) List(w http.ResponseWriter, r *http.Request) {
// 	storeCategories, err := h.storeCategorySvc.List(r.Context())
// 	if err != nil {
// 		httpres.WithErr(w, err)
// 		return
// 	}
//
// 	httpres.Success(w, storeCategories, http.StatusOK)
// }

type CreateStoreCategoryReq struct {
	Name string `json:"name"`
}

func (h StoreCategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	req, err := httpreq.Bind[CreateStoreCategoryReq](r)
	if err != nil {
		httpres.ValidationErr(w, err, http.StatusBadRequest)
		return
	}

	storeID := r.Context().Value(middleware.SellerCtxKey).(float64)

	if err := h.storeCategorySvc.Create(r.Context(), int(storeID), req.Name); err != nil {
		httpres.WithErr(w, err)
		return
	}

	httpres.Success(w, nil, http.StatusCreated)
}
