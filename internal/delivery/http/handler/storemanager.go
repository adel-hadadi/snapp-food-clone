package handler

import (
	"context"
	"net/http"

	otpservice "snapp-food/internal/service/otp"
	tokenservice "snapp-food/internal/service/token"
	"snapp-food/pkg/apperr"
	"snapp-food/pkg/httpres"
	"snapp-food/pkg/server/httpreq"
)

type StoreManagerHandler struct {
	storeSvc        storeService
	storeManagerSvc storeManagerService
	otpSvc          otpService
}

type storeManagerService interface {
	Login(ctx context.Context, phone string, code int) (tokenservice.TokenRes, error)
}

func NewStoreManagerHandler(storeSvc storeService, otpSvc otpService, storeManagerSvc storeManagerService) StoreManagerHandler {
	return StoreManagerHandler{
		storeSvc:        storeSvc,
		otpSvc:          otpSvc,
		storeManagerSvc: storeManagerSvc,
	}
}

type SendOTPReq struct {
	Phone string `json:"phone"`
}

const storeManagerOTPPrefix = "store-manager"

func (h StoreManagerHandler) SendOTP(w http.ResponseWriter, r *http.Request) {
	req, err := httpreq.Bind[SendOTPReq](r)
	if err != nil {
		httpres.WithErr(w, err)
		return
	}

	exists, err := h.storeSvc.ExistsByPhone(r.Context(), req.Phone)
	if err != nil {
		httpres.WithErr(w, err)
		return
	}

	if !exists {
		httpres.WithErr(w, apperr.New(apperr.NotFound))
		return
	}

	if err := h.otpSvc.Send(r.Context(), otpservice.OTPSendReq{
		Phone:  req.Phone,
		Prefix: storeManagerOTPPrefix,
	}); err != nil {
		httpres.WithErr(w, err)
		return
	}

	httpres.Success(w, nil, http.StatusOK)
}

type LoginReq struct {
	Phone string `json:"phone"`
	Code  int    `json:"code"`
}

func (h StoreManagerHandler) Login(w http.ResponseWriter, r *http.Request) {
	req, err := httpreq.Bind[LoginReq](r)
	if err != nil {
		httpres.ValidationErr(w, err, http.StatusBadRequest)
		return
	}

	token, err := h.storeManagerSvc.Login(r.Context(), req.Phone, req.Code)
	if err != nil {
		httpres.WithErr(w, err)
		return
	}

	httpres.Success(w, token, http.StatusOK)
}
