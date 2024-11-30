package http

import (
	"fmt"
	"net/http"
)

type HttpServer struct {
}

func New() HttpServer {
	return HttpServer{}
}

func (s HttpServer) Run() {
	srv := &http.Server{}

	if err := srv.ListenAndServe(); err != nil {
		panic(fmt.Errorf("error on start serving: %w", err))
	}
}
