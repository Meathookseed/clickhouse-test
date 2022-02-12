package handler

import (
	"github.com/buaazp/fasthttprouter"
	"log"
)

func NewRouter(
	params NewRouterParams,
	logger *log.Logger,
) *fasthttprouter.Router {
	router := fasthttprouter.New()

	for _, endpoint := range params.Endpoints {
		router.Handle(
			endpoint.Method,
			endpoint.Path,
			NewRecoverMiddleware(logger)(endpoint.Handler),
		)
	}

	return router
}
