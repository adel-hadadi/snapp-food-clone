package httpres

import (
	"net/http"
	"snapp-food/pkg/apperr"
)

var errToHTTPStatusCode = map[apperr.Type]int{
	apperr.NotFound:     http.StatusBadRequest,
	apperr.Conflict:     http.StatusBadRequest,
	apperr.Unauthorized: http.StatusUnauthorized,
	apperr.Forbidden:    http.StatusForbidden,
	apperr.Invalid:      http.StatusBadRequest,
	apperr.Unexpected:   http.StatusUnprocessableEntity,
}

func toHTTPCode(err error) (int, error) {
	apperr, ok := err.(*apperr.AppErr)
	if !ok {
		return http.StatusInternalServerError, err
	}

	code, exists := errToHTTPStatusCode[apperr.Type]
	if !exists {
		code = http.StatusInternalServerError
	}

	return code, err
}
