package handler

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	storeservice "snapp-food/internal/service/store"
	"snapp-food/pkg/httpres"
	"snapp-food/pkg/server/httpreq"

	"github.com/go-chi/chi/v5"
)

type StoreHandler struct {
	storeSvc storeService
}

type storeService interface {
	Create(ctx context.Context, req storeservice.CreateReq) error
	Find(ctx context.Context, slug string) (storeservice.StoreRes, error)
	List(ctx context.Context) ([]storeservice.StoreRes, error)
	ExistsByPhone(ctx context.Context, phone string) (bool, error)
	Nearest(ctx context.Context, userID int) ([]storeservice.StoreRes, error)
	ListByProductCategory(ctx context.Context, userID int, slug string) ([]storeservice.StoreRes, error)
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
	CityID           int     `json:"city_id"`
}

func (h StoreHandler) Create(w http.ResponseWriter, r *http.Request) {
	req, err := httpreq.Bind[CreateStoreReq](r)
	if err != nil {
		httpres.WithErr(w, err)
		return
	}

	if strings.HasPrefix(req.Phone, "0") {
		req.Phone = req.Phone[1:]
	}

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
		CityID:           req.CityID,
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

	StoreType StoreTypeRes `json:"store_type"`
}

func (h StoreHandler) Find(w http.ResponseWriter, r *http.Request) {
	storeSlug := chi.URLParam(r, "slug")

	store, err := h.storeSvc.Find(r.Context(), storeSlug)
	if err != nil {
		httpres.WithErr(w, err)
		return
	}

	httpres.Success(w, store, http.StatusOK)
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

func (h StoreHandler) ListNearest(w http.ResponseWriter, r *http.Request) {
	userID := httpreq.AuthID(r)

	stores, err := h.storeSvc.Nearest(r.Context(), userID)
	if err != nil {
		httpres.WithErr(w, err)
		return
	}

	res := make([]StoreRes, 0, len(stores))
	for s := range stores {
		res = append(res, h.DTOToRes(stores[s]))
	}

	httpres.Success(w, res, http.StatusOK)
}

func (h StoreHandler) ListByProductCategory(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	userID := httpreq.AuthID(r)

	stores, err := h.storeSvc.ListByProductCategory(r.Context(), userID, slug)
	if err != nil {
		httpres.WithErr(w, err)
		return
	}

	httpres.Success(w, stores, http.StatusOK)
}

func (h StoreHandler) Dashboard(w http.ResponseWriter, r *http.Request) {
	log.Println("dashboard")

	w.Write([]byte("hello store manager"))
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

		StoreType: StoreTypeRes{
			ID:    store.StoreType.ID,
			Name:  store.StoreType.Name,
			Image: store.StoreType.Image,
			URL:   store.StoreType.URL,
		},
	}
}
