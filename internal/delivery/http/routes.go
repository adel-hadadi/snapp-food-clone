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

		r.Get("/addresses", s.Handlers.Profile.GetAddresses)
		r.Post("/addresses", s.Handlers.Profile.CreateAddress)
	})

	router.Route("/panel", func(r chi.Router) {
		r.Use(middleware.Authenticate(s.TokenSvc))
		r.Get("/stores/nearest", s.Handlers.Store.ListNearest)
		r.Get("/products", s.Handlers.Product.FilteredList)
	})

	router.Route("/stores", func(r chi.Router) {
		r.Post("/", s.Handlers.Store.Create)
		r.Get("/{slug}", s.Handlers.Store.Find)
		r.Get("/", s.Handlers.Store.List)

		r.Post("/otp", s.Handlers.StoreManager.SendOTP)
		r.Post("/login", s.Handlers.StoreManager.Login)

		r.Route("/dashboard", func(r chi.Router) {
			r.Use(middleware.DashboardAuthenticate(s.TokenSvc))

			r.Get("/", s.Handlers.Store.Dashboard)

			r.Route("/products", func(r chi.Router) {
				r.Post("/", s.Handlers.Product.Create)
				r.Get("/", s.Handlers.Product.List)
			})

			r.Route("/categories", func(r chi.Router) {
				r.Post("/", s.Handlers.StoreCategory.Create)
			})
		})
	})

	router.Get("/store-types", s.Handlers.StoreType.Get)
	router.Get("/product-categories", s.Handlers.ProductCategory.List)
	return router
}
