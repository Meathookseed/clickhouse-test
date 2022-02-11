package dependency

import (
	"project-clickhouse/db"
	"project-clickhouse/handler"
	events "project-clickhouse/module/game_events"

	"go.uber.org/fx"
)

func Provider() fx.Option {
	return fx.Options(
		fx.Provide(NewLogger),
		events.Module,
		handler.Module,
		db.Module,
	)
}
