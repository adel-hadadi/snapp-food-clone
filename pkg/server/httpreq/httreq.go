package httpreq

import (
	"encoding/json"
	"net/http"
)

func Bind[T any](r *http.Request) (T, error) {
	var req T
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return req, err
	}

	return req, nil
}
