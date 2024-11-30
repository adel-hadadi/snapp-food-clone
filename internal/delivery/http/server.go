package http

import (
	"snapp-food/cmd/app"
	"snapp-food/pkg/server"
)

type HttpServer struct {
	Handlers app.Handlers
}

func New(handlers app.Handlers) HttpServer {
	return HttpServer{
		Handlers: handlers,
	}
}

func (s HttpServer) Run(port string) {
	server.RunHttpServerOnAddr(port, s.setRoutes)
}
