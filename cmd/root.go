package cmd

import "github.com/spf13/cobra"

func newRootCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "project-clickhouse",
		Short: "project clickhouse",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Usage()
		},
	}

	return command
}
