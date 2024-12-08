package handler

import (
	"context"
	"net/http"
	"snapp-food/internal/delivery/http/middleware"
	userservice "snapp-food/internal/service/user"
	"snapp-food/pkg/httpres"
	"snapp-food/pkg/server/httpreq"
)

type ProfileHandler struct {
	userSvc userService
}

type userService interface {
	Get(ctx context.Context, userID int) (userservice.UserRes, error)
	Update(ctx context.Context, userID int, req userservice.UpdateUserReq) error
}

func NewProfileHandler(userSvc userService) ProfileHandler {
	return ProfileHandler{
		userSvc: userSvc,
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

	h.userSvc.Update(r.Context(), userID, userservice.UpdateUserReq{
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		NationalID: req.NationalID,
	})

	httpres.Success(w, nil, http.StatusOK)
}
