package config

import (
	"errors"

	"github.com/caarlos0/env"
)

// LookupKey for context
type LookupKey string

const (
	// LoggedInUser is the key in context
	LoggedInUser LookupKey = "LoggedInUser"
)

// Config reads values from environment variables
type Config struct {
	DBUser          string `env:"DB_USER" envDefault:"yaesuser"`
	DBPassword      string `env:"DB_PASSWORD"`
	DBHost          string `env:"DB_HOST" envDefault:"yaes-db-postgresql"`
	DBPort          string `env:"DB_PORT" envDefault:"5432"`
	DBName          string `env:"DB_NAME" envDefault:"yaes"`
	DBUrl           string `env:"DATABASE_URL" envDefault:""`
	AppEnv          string `env:"GO_APP_ENV" envDefault:"production"`
	Port            string `env:"PORT" envDefault:"12345"`
	DiagnosticsPort string `env:"DIAGNOSTICS_PORT" envDefault:"8080"`
	JWTSigningKey   string `env:"JWT_KEY" envDefault:"483175006c1088c849502ef22406ac4e"`
}

// New initializes the config from environment variables
func New() Config {
	var cfg Config
	env.Parse(&cfg)

	return cfg
}

// Validate validates the config
func (cfg *Config) Validate() error {
	dbURL := cfg.DBUrl
	dbName := cfg.DBName

	if len(dbURL) == 0 && len(dbName) == 0 {
		return errors.New("no database config provided! You can set it using DATABASE_URL or DB_NAME env variable")
	}

	return nil
}
