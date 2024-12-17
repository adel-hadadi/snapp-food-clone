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
}

func NewStoreCategoryHandler(service storeCategoryService) StoreCategoryHandler {
	return StoreCategoryHandler{storeCategorySvc: service}
}

type CreateStoreCategoryReq struct {
	Name string `json:"name"`
}

func (h StoreCategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	req, err := httpreq.Bind[CreateStoreCategoryReq](r)
	if err != nil {
		httpres.ValidationErr(w, err, http.StatusBadRequest)
		return
	}

	storeID := r.Context().Value(middleware.StoreCtxKey).(float64)

	if err := h.storeCategorySvc.Create(r.Context(), int(storeID), req.Name); err != nil {
		httpres.WithErr(w, err)
		return
	}

	httpres.Success(w, nil, http.StatusCreated)
}
