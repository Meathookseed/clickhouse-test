package handler

import (
	"github.com/valyala/fasthttp"
	"log"
)

// NewRecoverMiddleware return a middleware which can let app recover from a panic in request handler.
func NewRecoverMiddleware(logger *log.Logger) func(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(h fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {
			defer func() {
				if rec := recover(); rec != nil {
					logger.Println("recover")
					ctx.Response.SetStatusCode(fasthttp.StatusInternalServerError)
					return
				}
			}()
			h(ctx)
		}
	}
}
