package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s HttpServer) setRoutes(router chi.Router) http.Handler {
	router.Route("/auth", func(r chi.Router) {
		r.Post("/login", s.Handlers.Auth.Login)
		r.Post("/register", s.Handlers.Auth.Register)
	})

	return router
}
