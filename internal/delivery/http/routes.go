package http

import (
	"net/http"

	"snapp-food/internal/delivery/http/middleware"

	"github.com/go-chi/chi/v5"
)

func (s HttpServer) setRoutes(router chi.Router) http.Handler {
	router.Route("/auth", func(r chi.Router) {
		r.Post("/otp", s.Handlers.OTP.Send)
		r.Post("/login-register", s.Handlers.Auth.LoginRegister)
	})

	router.Route("/profile", func(r chi.Router) {
		r.Use(middleware.Authenticate(s.TokenSvc))
		r.Get("/personal-info", s.Handlers.Profile.PersonalInfo)
		r.Put("/personal-info", s.Handlers.Profile.Update)
	})

	router.Get("/store-types", s.Handlers.StoreType.Get)
	return router
}
