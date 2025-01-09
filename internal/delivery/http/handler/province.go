package handler

import (
	"net/http"
	provinceservice "snapp-food/internal/service/province"
	"snapp-food/pkg/httpres"
)

type ProvinceHandler struct {
	provinceSvc provinceservice.Service
}

func NewProvinceHandler(provinceSvc provinceservice.Service) ProvinceHandler {
	return ProvinceHandler{provinceSvc: provinceSvc}
}

func (h ProvinceHandler) List(w http.ResponseWriter, r *http.Request) {
	provinces, err := h.provinceSvc.List(r.Context())
	if err != nil {
		httpres.WithErr(w, err)
		return
	}

	httpres.Success(w, provinces, http.StatusOK)
}
