package db

import (
	"database/sql"
	goClickhouse "github.com/mailru/go-clickhouse"
)

type Connection struct {
	*sql.DB
}

func NewDatabaseConnection(cfg *Config) (*Connection, error) {
	conn, err := sql.Open("clickhouse", newClickHouseConfig(cfg).FormatDSN())
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return &Connection{DB: conn}, nil
}

func newClickHouseConfig(cfg *Config) *goClickhouse.Config {
	return &goClickhouse.Config{
		Scheme:   "http",
		Host:     cfg.Addr,
		Database: cfg.DBName,
		Password: cfg.DBPassword,
		User:     cfg.DBUsername,
	}
}
