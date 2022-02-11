package app

import (
	"project-clickhouse/dependency"
	"project-clickhouse/server"

	"go.uber.org/fx"
)

func NewApp() error {
	app := fx.New(
		dependency.Provider(),
		server.Module,
	)

	app.Run()

	return app.Err()
}
