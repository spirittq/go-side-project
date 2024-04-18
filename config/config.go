package config

import (
	"os"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	Log
	GORM
	HTTP
}

type Log struct {
	Level string `env:"LOG_LEVEL"`
}

type GORM struct {
	URL     string `env:"DB_URL"`
}

type HTTP struct {
	Address string `env:"HTTP_ADDRESS"`
	Timeout int `env:"HTTP_TIMEOUT"`
}

func NewConfig() (*Config, error) {
	if os.Getenv("ENV") == "" {
		err := godotenv.Load(".env.local")
		if err != nil {
			return nil, err
		}
	}

	cfg := Config{}
	opts := env.Options{RequiredIfNoDef: true}
	if err := env.ParseWithOptions(&cfg, opts); err != nil {
		return nil, err
	}

	return &cfg, nil
}
