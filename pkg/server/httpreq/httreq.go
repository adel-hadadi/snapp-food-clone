package httpreq

import (
	"encoding/json"
	"net/http"
	"snapp-food/internal/delivery/http/middleware"
	tokenservice "snapp-food/internal/service/token"
)

func AuthID(r *http.Request) int {
	userIDRaw := r.Context().Value(middleware.UserCtxKey)

	return userIDRaw.(tokenservice.User).ID
}

func Bind[T any](r *http.Request) (T, error) {
	var req T
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return req, err
	}

	return req, nil
}
