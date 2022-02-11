package handler

import (
	"github.com/valyala/fasthttp"
	"go.uber.org/fx"
)

func RegisterEndpoint(target interface{}) fx.Annotated {
	return fx.Annotated{Group: "http_endpoints", Target: target}
}

type Endpoint struct {
	Path    string
	Method  string
	Handler fasthttp.RequestHandler
}

type NewRouterParams struct {
	fx.In
	Endpoints []*Endpoint `group:"http_endpoints"`
}
