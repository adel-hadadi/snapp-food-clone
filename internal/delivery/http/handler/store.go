package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	storeservice "snapp-food/internal/service/store"
	"snapp-food/pkg/httpres"
	"snapp-food/pkg/server/httpreq"
)

type StoreHandler struct {
	storeSvc storeService
}

type storeService interface {
	Create(ctx context.Context, req storeservice.CreateReq) error
	Find(ctx context.Context, slug string) (storeservice.StoreRes, error)
	List(ctx context.Context) ([]storeservice.StoreRes, error)
}

func NewStoreHandler(storeSvc storeService) StoreHandler {
	return StoreHandler{storeSvc: storeSvc}
}

type CreateStoreReq struct {
	Name             string  `json:"name"`
	Latitude         float64 `json:"latitude"`
	Longitude        float64 `json:"longitude"`
	Address          string  `json:"address"`
	StoreTypeID      int     `json:"store_type_id"`
	Logo             string  `json:"logo"`
	Phone            string  `json:"phone"`
	ManagerFirstName string  `json:"manager_first_name"`
	ManagerLastName  string  `json:"manager_last_name"`
}

func (h StoreHandler) Create(w http.ResponseWriter, r *http.Request) {
	req, err := httpreq.Bind[CreateStoreReq](r)
	if err != nil {
		httpres.WithErr(w, err)
		return
	}

	// TODO: validate data

	if err := h.storeSvc.Create(r.Context(), storeservice.CreateReq{
		Name:             req.Name,
		Phone:            req.Phone,
		ManagerFirstName: req.ManagerFirstName,
		ManagerLastName:  req.ManagerLastName,
		Address:          req.Address,
		StoreTypeID:      req.StoreTypeID,
		Latitude:         req.Latitude,
		Longitude:        req.Longitude,
		Logo:             req.Logo,
	}); err != nil {
		httpres.WithErr(w, err)
		return
	}

	httpres.Success(w, nil, http.StatusCreated)
}

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
}

func (h StoreHandler) Find(w http.ResponseWriter, r *http.Request) {
	storeSlug := chi.URLParam(r, "slug")

	store, err := h.storeSvc.Find(r.Context(), storeSlug)
	if err != nil {
		httpres.WithErr(w, err)
		return
	}

	httpres.Success(w, h.DTOToRes(store), http.StatusOK)
}

func (h StoreHandler) List(w http.ResponseWriter, r *http.Request) {
	stores, err := h.storeSvc.List(r.Context())
	if err != nil {
		httpres.WithErr(w, err)
		return
	}

	res := make([]StoreRes, 0, len(stores))

	for _, store := range stores {
		res = append(res, h.DTOToRes(store))
	}

	httpres.Success(w, res, http.StatusOK)
}

func (h StoreHandler) DTOToRes(store storeservice.StoreRes) StoreRes {
	return StoreRes{
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
}
