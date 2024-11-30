package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s HttpServer) setRoutes(router chi.Router) http.Handler {
	router.Get("/auth/login", s.Handlers.Auth.Login)

	return router
}
