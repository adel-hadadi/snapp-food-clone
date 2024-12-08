package httpreq

import (
	"encoding/json"
	"net/http"
	"snapp-food/internal/delivery/http/middleware"
)

func AuthID(r *http.Request) int {
	userIDRaw := r.Context().Value(middleware.UserCtxKey)

	userID := userIDRaw.(float64)

	return int(userID)
}

func Bind[T any](r *http.Request) (T, error) {
	var req T
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return req, err
	}

	return req, nil
}
