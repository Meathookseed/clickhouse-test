package main

import (
	"log"
	"project-clickhouse/cmd"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("failed to start server: %s", err)
	}
}
