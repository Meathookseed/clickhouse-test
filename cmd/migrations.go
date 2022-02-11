package cmd

import (
	"context"
	"log"
	"project-clickhouse/migration"

	"github.com/spf13/cobra"
)

func newMigrationExecuteCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "migrations:execute",
		Short: "Execute migrations",
		RunE: func(cmd *cobra.Command, args []string) error {
			app := migration.NewMigrationExecutor()

			startCtx := context.Background()

			if err := app.Start(startCtx); err != nil {
				log.Fatalf("Could not migrate: %v", err)
			}

			log.Printf("Migration did run successfully")

			return nil
		},
	}
}
