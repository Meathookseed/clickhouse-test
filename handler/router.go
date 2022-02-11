package handler

import (
	"github.com/buaazp/fasthttprouter"
)

func NewRouter(
	params NewRouterParams,
) *fasthttprouter.Router {
	router := fasthttprouter.New()

	for _, endpoint := range params.Endpoints {
		router.Handle(
			endpoint.Method,
			endpoint.Path,
			endpoint.Handler,
		)
	}

	return router
}
