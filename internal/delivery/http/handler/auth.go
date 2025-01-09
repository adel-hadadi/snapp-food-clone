package handler

import (
	"context"
	"net/http"
	"snapp-food/internal/adapters/validation"
	"snapp-food/internal/dto"
	authservice "snapp-food/internal/service/auth"
	tokenservice "snapp-food/internal/service/token"
	"snapp-food/pkg/httpres"
	"snapp-food/pkg/server/httpreq"
)

type AuthHandler struct {
	validator validation.AuthValidation
	authSvc   authService
}

type authService interface {
	LoginRegister(ctx context.Context, req authservice.LoginRegisterReq) (tokenservice.TokenRes, error)
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

	if errs, ok := h.validator.ValidateLoginRegister(req); !ok {
		httpres.ValidationErr(w, errs, http.StatusBadRequest)
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

	httpres.Success(w, dto.LoginRegisterRes{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, http.StatusOK)
}
