package events

import (
	"log"
	"net/http"
	"project-clickhouse/handler"

	realip "github.com/Ferluci/fast-realip"
	"github.com/valyala/fasthttp"
)

func NewCreateGameEventEndpoint(h *Handler) *handler.Endpoint {
	return &handler.Endpoint{Method: http.MethodPost, Path: "/game-event", Handler: h.createGameEventEndpoint}
}

type Handler struct {
	service *Service
	logger  *log.Logger
}

func NewHandler(service *Service, logger *log.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

func (h *Handler) createGameEventEndpoint(ctx *fasthttp.RequestCtx) {
	events, err := h.service.Deserialize(ctx.Request.Body())
	if err != nil {
		h.logger.Println(err)
		ctx.SetStatusCode(http.StatusBadRequest)

		return
	}

	err = h.service.CreateEvents(ctx, events, realip.FromRequest(ctx))
	if err != nil {
		h.logger.Println(err)
		ctx.SetStatusCode(http.StatusInternalServerError)

		return
	}

	ctx.SetStatusCode(http.StatusOK)
}
