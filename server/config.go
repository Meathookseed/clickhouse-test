package server

import (
	"github.com/joeshaw/envdecode"
	"github.com/pkg/errors"
)

type Config struct {
	APIBindingAddr string `env:"API_BINDING_ADDRESS,default=0.0.0.0"`
	APIBindingPort string `env:"API_BINDING_PORT,default=9191"`
}

func NewConfig() (*Config, error) {
	c := &Config{}

	if err := envdecode.Decode(c); err != nil {
		return nil, errors.WithStack(err)
	}

	return c, nil
}
