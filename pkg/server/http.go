package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

const ErrStartServer = "error on start server: %w"

func RunHttpServerOnAddr(port string, createHandler func(router chi.Router) http.Handler) {
	apiRouter := chi.NewRouter()

	rootRouter := chi.NewRouter()

	setMiddleware(rootRouter)

	createHandler(apiRouter)

	rootRouter.Mount("/api", apiRouter)

	chi.Walk(rootRouter, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("[%s]: '%s' \n", method, route)
		return nil
	})

	log.Printf("server listening on port %s \n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), rootRouter); err != nil {
		panic(fmt.Errorf(ErrStartServer, err))
	}
}

func setMiddleware(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Heartbeat("/ping"))
	router.Use(middleware.Logger)
	router.Use(middleware.CleanPath)

	addCorsMiddleware(router)
}

func addCorsMiddleware(router *chi.Mux) {
	allowedOrigins := strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ";")
	if len(allowedOrigins) == 0 {
		return
	}

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	router.Use(corsMiddleware.Handler)
}
