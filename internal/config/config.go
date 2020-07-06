package config

import (
	"github.com/caarlos0/env"
	"errors"
)

// Config reads values from environment variables
type Config struct {
	DBName string `env:"DB_NAME"`
	DBUrl  string `env:"DATABASE_URL"`
	AppEnv string `env:"GO_APP_ENV" envDefault:"production"`
	Port   string `env:"PORT" envDefault:"12345"`
}

// New initializes the config from environment variables
func New() Config {
	var cfg Config
	env.Parse(&cfg)

	return cfg
}

func (cfg *Config) Validate() error {
	dbURL := cfg.DBUrl
	dbName := cfg.DBName

	if len(dbURL) == 0 && len(dbName) == 0 {
		return errors.New("no database config provided! You can set it using DATABASE_URL or DB_NAME env variable.")
	}

	return nil
}