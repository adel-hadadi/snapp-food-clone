package authhandler

import (
	authservice "snapp-food/internal/service/auth"
)

type Handler struct {
	authSvc authservice.Service
}

func New(authSvc authservice.Service) Handler {
	return Handler{
		authSvc: authSvc,
	}
}
