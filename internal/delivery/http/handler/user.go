package handler

import (
	"net/http"
)

type ProfileHandler struct {
}

func NewProfileHandler() ProfileHandler {
	return ProfileHandler{}
}

func (h ProfileHandler) PersonalInfo(w http.ResponseWriter, r *http.Request) {
    h.
}
