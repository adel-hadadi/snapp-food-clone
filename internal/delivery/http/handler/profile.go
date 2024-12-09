package handler

import (
	"context"
	"net/http"

	"snapp-food/internal/delivery/http/middleware"
	userservice "snapp-food/internal/service/user"
	useraddressservice "snapp-food/internal/service/useraddress"
	"snapp-food/pkg/httpres"
	"snapp-food/pkg/server/httpreq"
)

type ProfileHandler struct {
	userSvc        userService
	userAddressSvc userAddressService
}

type userService interface {
	Get(ctx context.Context, userID int) (userservice.UserRes, error)
	Update(ctx context.Context, userID int, req userservice.UpdateUserReq) error
}

type userAddressService interface {
	Get(ctx context.Context, userID int) ([]useraddressservice.UserAddressRes, error)
	Create(ctx context.Context, userID int, req useraddressservice.CreateReq) error
}

func NewProfileHandler(userSvc userService, userAddressSvc userAddressService) ProfileHandler {
	return ProfileHandler{
		userSvc:        userSvc,
		userAddressSvc: userAddressSvc,
	}
}

type PersonalInfoRes struct {
	ID         int     `json:"id"`
	FirstName  *string `json:"first_name"`
	LastName   *string `json:"last_name"`
	Phone      string  `json:"phone"`
	NationalID *string `json:"national_id"`
}

func (h ProfileHandler) PersonalInfo(w http.ResponseWriter, r *http.Request) {
	userIDRaw := r.Context().Value(middleware.UserCtxKey)
	userID := userIDRaw.(float64)

	user, err := h.userSvc.Get(r.Context(), int(userID))
	if err != nil {
		httpres.WithErr(w, err)
		return
	}

	httpres.Success(w, PersonalInfoRes{
		ID:         user.ID,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Phone:      user.Phone,
		NationalID: user.NationalID,
	}, http.StatusOK)
}

type UpdateProfileReq struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	NationalID string `json:"national_id"`
}

func (h ProfileHandler) Update(w http.ResponseWriter, r *http.Request) {
	req, err := httpreq.Bind[UpdateProfileReq](r)
	if err != nil {
		httpres.WithErr(w, err)
		return
	}

	userID := httpreq.AuthID(r)

	if err := h.userSvc.Update(r.Context(), userID, userservice.UpdateUserReq{
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		NationalID: req.NationalID,
	}); err != nil {
		httpres.WithErr(w, err)
		return
	}

	httpres.Success(w, nil, http.StatusOK)
}

type UserAddresses struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Address   string  `json:"address"`
	Detail    string  `json:"detail"`
	Province  string  `json:"province"`
	City      string  `json:"city"`
}

func (h ProfileHandler) GetAddresses(w http.ResponseWriter, r *http.Request) {
	userID := httpreq.AuthID(r)

	addresses, err := h.userAddressSvc.Get(r.Context(), userID)
	if err != nil {
		httpres.WithErr(w, err)
		return
	}

	response := make([]UserAddresses, 0, len(addresses))
	for i := range addresses {
		response = append(response, UserAddresses{
			ID:        addresses[i].ID,
			Name:      addresses[i].Name,
			Latitude:  addresses[i].Latitude,
			Longitude: addresses[i].Longitude,
			Address:   addresses[i].Address,
			Detail:    addresses[i].Detail,
			Province:  addresses[i].Province,
			City:      addresses[i].City,
		})
	}

	httpres.Success(w, response, http.StatusOK)
}

type CreateAddressReq struct {
	Name       string  `json:"name"`
	Address    string  `json:"address"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Detail     string  `json:"detail"`
	ProvinceID int     `json:"province_id"`
	CityID     int     `json:"city_id"`
}

func (h ProfileHandler) CreateAddress(w http.ResponseWriter, r *http.Request) {
	req, err := httpreq.Bind[CreateAddressReq](r)
	if err != nil {
		httpres.WithErr(w, err)
		return
	}

	userID := httpreq.AuthID(r)

	err = h.userAddressSvc.Create(r.Context(), userID, useraddressservice.CreateReq{
		Name:       req.Name,
		Address:    req.Address,
		Detail:     req.Detail,
		Latitude:   req.Latitude,
		Longitude:  req.Longitude,
		ProvinceID: req.ProvinceID,
		CityID:     req.CityID,
	})
	if err != nil {
		httpres.WithErr(w, err)
		return
	}

	httpres.Success(w, nil, http.StatusCreated)
}
