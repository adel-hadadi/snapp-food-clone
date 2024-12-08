package handler

import (
	"context"
	"net/http"
	authservice "snapp-food/internal/service/auth"
	tokenservice "snapp-food/internal/service/token"
	"snapp-food/pkg/httpres"
	"snapp-food/pkg/server/httpreq"
	"snapp-food/pkg/validate"
)

type AuthHandler struct {
	validator validate.Validator
	authSvc   authService
}

type authService interface {
	LoginRegister(ctx context.Context, req authservice.LoginRegisterReq) (tokenservice.TokenRes, error)
}

func NewAuthHandler(v validate.Validator, authSvc authService) AuthHandler {
	return AuthHandler{
		validator: v,
		authSvc:   authSvc,
	}
}

type LoginRegisterReq struct {
	Phone string `json:"phone"`
	Code  int    `json:"code"`
}

type LoginRegisterRes struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (h AuthHandler) LoginRegister(w http.ResponseWriter, r *http.Request) {
	req, err := httpreq.Bind[LoginRegisterReq](r)
	if err != nil {
		httpres.WithErr(w, err)
		return
	}

	token, err := h.authSvc.LoginRegister(r.Context(), authservice.LoginRegisterReq{
		Phone: req.Phone,
		Code:  req.Code,
	})
	if err != nil {
		httpres.WithErr(w, err)
		return
	}

	httpres.Success(w, LoginRegisterRes{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, http.StatusOK)
}

type RegisterReq struct {
	Name  string `json:"name" validate:"required"`
	Phone string `json:"phone" validate:"required"`
}

func (h AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	req, err := httpreq.Bind[RegisterReq](r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.validator.Struct(req); err != nil {
		httpres.ValidationErr(w, err, http.StatusBadRequest)
		return
	}

	httpres.Success(w, nil, http.StatusOK)
}
