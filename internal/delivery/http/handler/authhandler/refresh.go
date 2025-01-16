package authhandler

import (
	"net/http"
	"snapp-food/pkg/httpres"
	"snapp-food/pkg/server/httpreq"
)

type RefreshReq struct {
	RefreshToken string `json:"refresh_token"`
}

func (h Handler) Refresh(w http.ResponseWriter, r *http.Request) {
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
