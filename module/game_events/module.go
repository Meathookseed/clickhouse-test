package events

import (
	"project-clickhouse/handler"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	NewGameEventRepository,
	NewService,
	NewHandler,
	handler.RegisterEndpoint(NewCreateGameEventEndpoint),
)
