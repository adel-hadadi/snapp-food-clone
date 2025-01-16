package sellerstorehandler

import (
	"net/http"
	"snapp-food/pkg/httpres"
	"snapp-food/pkg/server/httpreq"
)

func (h Handler) List(w http.ResponseWriter, r *http.Request) {
	manager := httpreq.AuthID(r)

	stores, err := h.storeSvc.ListByManagerID(r.Context(), manager)
	if err != nil {
		httpres.WithErr(w, err)
		return
	}

	httpres.Success(w, stores, http.StatusOK)
}
