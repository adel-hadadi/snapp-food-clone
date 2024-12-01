package httpres

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool              `json:"success"`
	Data    any               `json:"data,omitempty"`
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors,omitempty"`
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

func ValidationErr(w http.ResponseWriter, errs map[string]string, status int) {
	res := Response{
		Success: false,
		Errors:  errs,
	}

	w.WriteHeader(status)

	bytes, _ := json.Marshal(res)

	w.Write(bytes)
}
