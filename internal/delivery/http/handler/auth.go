package handler

import (
	"context"
	"net/http"
	"snapp-food/internal/adapters/validation"
	"snapp-food/internal/dto"
	authservice "snapp-food/internal/service/auth"
	"snapp-food/pkg/httpres"
	"snapp-food/pkg/server/httpreq"
)

type AuthHandler struct {
	validator validation.AuthValidation
	authSvc   authService
}

type authService interface {
	LoginRegister(ctx context.Context, req authservice.LoginRegisterReq) (authservice.LoginRegisterRes, error)
	GetAccessToken(ctx context.Context, refreshToken string) (string, error)
}

func NewAuthHandler(v validation.AuthValidation, authSvc authService) AuthHandler {
	return AuthHandler{
		validator: v,
		authSvc:   authSvc,
	}
}

func (h AuthHandler) LoginRegister(w http.ResponseWriter, r *http.Request) {
	req, err := httpreq.Bind[dto.AuthLoginRegisterReq](r)
	if err != nil {
		httpres.WithErr(w, err)
		return
	}

	// TODO: fix validation here

	userTokenRes, err := h.authSvc.LoginRegister(r.Context(), authservice.LoginRegisterReq{
		Phone: req.Phone,
		Code:  req.Code,
	})
	if err != nil {
		httpres.WithErr(w, err)
		return
	}

	httpres.Success(w, userTokenRes, http.StatusOK)
}

type RefreshReq struct {
	RefreshToken string `json:"refresh_token"`
}

func (h AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	req, err := httpreq.Bind[RefreshReq](r)
	if err != nil {
		httpres.WithErr(w, err)
		return
	}

	token, err := h.authSvc.GetAccessToken(r.Context(), req.RefreshToken)
	if err != nil {
		httpres.WithErr(w, err)
		return
	}

	httpres.Success(w, map[string]string{"access_token": token}, http.StatusOK)
}
