package http

import (
	"net/http"

	"snapp-food/internal/delivery/http/middleware"

	"github.com/go-chi/chi/v5"
)

func (s HttpServer) setRoutes(router chi.Router) http.Handler {
	router.Route("/provinces", func(r chi.Router) {
		r.Get("/", s.Handlers.Province.List)

		r.Get("/{provinceID}/cities", s.Handlers.City.ListByProvince)
	})

	router.Route("/auth", func(r chi.Router) {
		r.Post("/otp", s.Handlers.Auth.SendCode)
		r.Post("/sellers/otp", s.Handlers.Auth.SellersSendCode)

		r.Post("/login-register", s.Handlers.Auth.LoginRegister)
		r.Post("/sellers/login-register", s.Handlers.Auth.SellerLoginRegister)

		r.Post("/refresh", s.Handlers.Auth.Refresh)
	})

	router.Route("/profile", func(r chi.Router) {
		r.Use(middleware.Authenticate(s.TokenSvc))
		r.Get("/", s.Handlers.Profile.Get)
		r.Put("/", s.Handlers.Profile.Update)

		r.Get("/addresses", s.Handlers.Profile.GetAddresses)
		r.Post("/addresses", s.Handlers.Profile.CreateAddress)

		r.Route("/orders", func(r chi.Router) {
			r.Post("/", s.Handlers.Order.Create)
			r.Get("/", s.Handlers.Order.List)

			r.Get("/{orderID}/pay", s.Handlers.Order.Pay)
		})
	})

	router.Route("/panel", func(r chi.Router) {
		r.Use(middleware.Authenticate(s.TokenSvc))
		r.Get("/stores/nearest", s.Handlers.Store.ListNearest)
		r.Get("/products", s.Handlers.Product.FilteredList)

		r.Get("/product-categories/{slug}/stores", s.Handlers.Store.ListByProductCategory)
	})

	router.Route("/sellers/dashboard", func(r chi.Router) {
		r.Use(middleware.DashboardAuthenticate(s.TokenSvc))

		r.Get("/", s.Handlers.Store.Dashboard)

		r.Route("/stores", func(r chi.Router) {
			r.Get("/", s.Handlers.SellerStore.List)
			r.Post("/", s.Handlers.SellerStore.Create)
		})

		r.Route("/products", func(r chi.Router) {
			r.Post("/", s.Handlers.Product.Create)
			r.Get("/", s.Handlers.Product.List)
		})

		r.Route("/categories", func(r chi.Router) {
			r.Post("/", s.Handlers.StoreCategory.Create)
		})
	})

	router.Get("/store-types", s.Handlers.StoreType.Get)
	router.Get("/product-categories", s.Handlers.ProductCategory.List)
	return router
}
