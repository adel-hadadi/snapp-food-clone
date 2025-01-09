package handler

import (
	"net/http"
	cityservice "snapp-food/internal/service/city"
	"snapp-food/pkg/httpres"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type CityHandler struct {
	citySvc cityservice.Service
}

func NewCityHandler(citySvc cityservice.Service) CityHandler {
	return CityHandler{citySvc: citySvc}
}

func (h CityHandler) ListByProvince(w http.ResponseWriter, r *http.Request) {
	provinceIDRaw := chi.URLParam(r, "provinceID")
	provinceID, _ := strconv.Atoi(provinceIDRaw)

	cities, err := h.citySvc.ListByProvinceID(r.Context(), provinceID)
	if err != nil {
		httpres.WithErr(w, err)
		return
	}

	httpres.Success(w, cities, http.StatusOK)
}
