package migration

import (
	"project-clickhouse/db"

	clickhouseMigrate "github.com/golang-migrate/migrate/v4/database/clickhouse"
	"go.uber.org/fx"
)

func RunMigrations(conn *db.Connection, registry *Registry, dbConfig *db.Config) error {
	driver, err := clickhouseMigrate.WithInstance(conn.DB, &clickhouseMigrate.Config{
		DatabaseName:          dbConfig.DBName,
		MultiStatementEnabled: true,
	})
	if err != nil {
		return err
	}

	for registry.Next() {
		migration, err := registry.Migration()
		if err != nil {
			return err
		}

		err = driver.Run(migration)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewMigrationExecutor() *fx.App {
	return fx.New(
		db.Module,
		fx.Provide(NewRegistry),
		fx.Invoke(
			RunMigrations,
		),
	)
}
