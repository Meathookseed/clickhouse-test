package cmd

import (
	"project-clickhouse/app"

	"github.com/spf13/cobra"
)

func newServerStartCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return app.NewApp()
		},
	}

	return cmd
}
