package core_http_server

import (
	"net/http"

	core_http_middleware "github.com/AndreyAndrian/golang-todoapp/internal/core/transport/http/middleware"
)

type Route struct {
	Method     string                            //post,get,update,put
	Path       string                            //url
	Handler    http.HandlerFunc                  //handler
	Middleware []core_http_middleware.Middleware //Middleware set
}


func (r *Route) WithMiddleware() http.Handler{
	return core_http_middleware.ChainMiddleware(
		r.Handler,
		r.Middleware...,
	)
}