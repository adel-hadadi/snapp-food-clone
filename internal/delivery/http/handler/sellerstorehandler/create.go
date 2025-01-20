package sellerstorehandler

import (
	"net/http"
	"snapp-food/internal/dto"
	"snapp-food/pkg/httpres"
	"snapp-food/pkg/server/httpreq"
)

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	req, err := httpreq.Bind[dto.StoreCreateReq](r)
	if err != nil {
		httpres.ValidationErr(w, err, http.StatusBadRequest)
		return
	}

	req.ManagerID = httpreq.AuthID(r)

	err = h.storeSvc.Create(r.Context(), req)
	if err != nil {
		httpres.WithErr(w, err)
		return
	}

	httpres.Success(w, nil, http.StatusCreated)
}
