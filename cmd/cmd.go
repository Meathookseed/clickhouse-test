package cmd

func Execute() error {
	rootCmd := newRootCommand()
	rootCmd.AddCommand(newMigrationExecuteCommand())
	rootCmd.AddCommand(newServerStartCommand())

	return rootCmd.Execute()
}
