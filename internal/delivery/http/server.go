package http

import (
	"snapp-food/cmd/app"
	tokenservice "snapp-food/internal/service/token"
	"snapp-food/pkg/server"
)

type HttpServer struct {
	Handlers app.Handlers
	TokenSvc tokenservice.Service
}

func New(handlers app.Handlers, tokenSvc tokenservice.Service) HttpServer {
	return HttpServer{
		Handlers: handlers,
		TokenSvc: tokenSvc,
	}
}

func (s HttpServer) Run(port string) {
	server.RunHttpServerOnAddr(port, s.setRoutes)
}
