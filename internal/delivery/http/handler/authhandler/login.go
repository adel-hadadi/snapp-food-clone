package authhandler

import (
	"net/http"
	"snapp-food/internal/dto"
	"snapp-food/pkg/httpres"
	"snapp-food/pkg/server/httpreq"
)

func (h Handler) LoginRegister(w http.ResponseWriter, r *http.Request) {
	req, err := httpreq.Bind[dto.AuthLoginRegisterReq](r)
	if err != nil {
		httpres.ValidationErr(w, err, http.StatusBadRequest)
		return
	}

	res, err := h.authSvc.LoginRegister(r.Context(), req)
	if err != nil {
		httpres.WithErr(w, err)
		return
	}

	httpres.Success(w, res, http.StatusOK)
}

func (h Handler) SellerLoginRegister(w http.ResponseWriter, r *http.Request) {
	req, err := httpreq.Bind[dto.AuthLoginRegisterReq](r)
	if err != nil {
		httpres.ValidationErr(w, err, http.StatusBadRequest)
		return
	}

	res, err := h.authSvc.SellerLoginRegister(r.Context(), req)
	if err != nil {
		httpres.WithErr(w, err)
		return
	}

	httpres.Success(w, res, http.StatusOK)
}
