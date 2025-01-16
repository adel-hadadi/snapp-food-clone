package authhandler

import (
	"net/http"
	"snapp-food/internal/dto"
	"snapp-food/pkg/httpres"
	"snapp-food/pkg/server/httpreq"
)

func (h Handler) SendCode(w http.ResponseWriter, r *http.Request) {
	req, err := httpreq.Bind[dto.AuthSendOTPReq](r)
	if err != nil {
		httpres.ValidationErr(w, err, http.StatusBadRequest)
		return
	}

	err = h.authSvc.SendCode(r.Context(), req.Phone)
	if err != nil {
		httpres.WithErr(w, err)
		return
	}

	httpres.Success(w, nil, http.StatusAccepted)
}

func (h Handler) SellersSendCode(w http.ResponseWriter, r *http.Request) {
	req, err := httpreq.Bind[dto.AuthSendSellerOTPReq](r)
	if err != nil {
		httpres.ValidationErr(w, err, http.StatusBadRequest)
		return
	}

	err = h.authSvc.SellerSendCode(r.Context(), req)
	if err != nil {
		httpres.WithErr(w, err)
		return
	}

	httpres.Success(w, nil, http.StatusAccepted)
}
