package httpres

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message"`
	Errors  any    `json:"errors,omitempty"`
}

func Success(w http.ResponseWriter, data any, status int) {
	res := Response{
		Data:    data,
		Success: true,
	}

	w.WriteHeader(status)

	bytes, _ := json.Marshal(res)

	w.Write(bytes)
}

func ValidationErr(w http.ResponseWriter, errs any, status int) {
	res := Response{
		Success: false,
		Errors:  errs,
	}

	w.WriteHeader(status)

	bytes, _ := json.Marshal(res)

	w.Write(bytes)
}

func WithErr(w http.ResponseWriter, err error) {
	code, err := toHTTPCode(err)

	res := Response{
		Success: false,
		Message: err.Error(),
	}

	bytes, _ := json.Marshal(res)

	w.WriteHeader(code)
	w.Write(bytes)
}
