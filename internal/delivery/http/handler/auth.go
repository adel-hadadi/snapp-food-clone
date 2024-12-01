package handler

import (
	"net/http"
	"snapp-food/pkg/httpres"
	"snapp-food/pkg/server/httpreq"
	"snapp-food/pkg/validate"
)

type AuthHandler struct {
	validator validate.Validator
}

func NewAuthHandler(v validate.Validator) AuthHandler {
	return AuthHandler{
		validator: v,
	}
}

func (h AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You logged in"))
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
