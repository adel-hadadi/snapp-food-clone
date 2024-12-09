package handler

import (
	"context"
	"net/http"

	storetypeservice "snapp-food/internal/service/storetype"
	"snapp-food/pkg/httpres"
)

type StoreTypeHandler struct {
	storeTypeSvc storeTypeService
}

type storeTypeService interface {
	Get(ctx context.Context) ([]storetypeservice.StoreType, error)
}

func NewStoreTypeHandler(storeTypeSvc storeTypeService) StoreTypeHandler {
	return StoreTypeHandler{
		storeTypeSvc: storeTypeSvc,
	}
}

type StoreTypeRes struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
	URL   string `json:"url"`
}

func (h StoreTypeHandler) Get(w http.ResponseWriter, r *http.Request) {
	types, err := h.storeTypeSvc.Get(r.Context())
	if err != nil {
		httpres.WithErr(w, err)
		return
	}

	response := make([]StoreTypeRes, 0, len(types))
	for t := range types {
		response = append(response, StoreTypeRes{
			ID:    types[t].ID,
			Name:  types[t].Name,
			Image: types[t].Image,
			URL:   types[t].URL,
		})
	}
	httpres.Success(w, response, http.StatusOK)
}
