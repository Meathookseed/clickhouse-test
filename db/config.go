package db

import (
	"github.com/joeshaw/envdecode"
	"github.com/pkg/errors"
)

type Config struct {
	Addr       string `env:"CLICKHOUSE_DB_ADDR,required"`
	DBName     string `env:"CLICKHOUSE_DB_NAME,required"`
	DBPassword string `env:"CLICKHOUSE_DB_PASSWORD"`
	DBUsername string `env:"CLICKHOUSE_DB_USERNAME"`
}

func NewConfig() (*Config, error) {
	c := &Config{}

	if err := envdecode.Decode(c); err != nil {
		return nil, errors.WithStack(err)
	}

	return c, nil
}
