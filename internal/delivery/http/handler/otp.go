package handler

import (
	"context"
	"net/http"

	otpservice "snapp-food/internal/service/otp"
	"snapp-food/pkg/httpres"
	"snapp-food/pkg/server/httpreq"
	"snapp-food/pkg/validate"
)

type OTPHandler struct {
	otpSvc    otpService
	validator validate.Validator
}

type otpService interface {
	Send(ctx context.Context, req otpservice.OTPSendReq) error
}

func NewOTPHandler(otpSvc otpService) OTPHandler {
	return OTPHandler{
		// TODO: add validation
		otpSvc: otpSvc,
	}
}

type OTPSendReq struct {
	Phone string `json:"phone" validate:"required"`
}

const userOTPPrefix = "user"

func (h OTPHandler) Send(w http.ResponseWriter, r *http.Request) {
	req, err := httpreq.Bind[OTPSendReq](r)
	if err != nil {
		httpres.ValidationErr(w, err, http.StatusBadRequest)
		return
	}

	if err := h.validator.Struct(req); err != nil {
		httpres.ValidationErr(w, err, http.StatusBadRequest)
		return
	}

	err = h.otpSvc.Send(r.Context(), otpservice.OTPSendReq{Phone: req.Phone, Prefix: userOTPPrefix})
	if err != nil {
		httpres.WithErr(w, err)
		return
	}

	httpres.Success(w, nil, http.StatusAccepted)
}
