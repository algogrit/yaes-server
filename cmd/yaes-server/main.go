package main

import (
	"github.com/caarlos0/env"
	log "github.com/sirupsen/logrus"

	api "github.com/gauravagarwalr/yaes-server/src/api"
	db "github.com/gauravagarwalr/yaes-server/src/config/db"

	"github.com/gauravagarwalr/raven-go"
)

// Config slurps the environment variables
type Config struct {
	DBName        string `env:"DB_NAME"`
	DBUrl         string `env:"DATABASE_URL"`
	AppEnv        string `env:"GO_APP_ENV" envDefault:"production"`
	Port          string `env:"PORT" envDefault:"12345"`
	SentryDsn     string `env:SENTRY_DSN`
	SentryRelease string `env:SENTRY_RELEASE envDefault:"production"`
}

var cfg Config

func init() {
	raven.SetDSN(cfg.SentryDsn)
}

func initDB() {
	dbURL := cfg.DBUrl
	dbName := cfg.DBName

	if len(dbURL) == 0 && len(dbName) == 0 {
		log.Fatal("No databases provided! You can set it using DATABASE_URL or DB_NAME env variable.")
	}

	db.InitializeDB(cfg.AppEnv, dbURL, dbName)
}

func initServer() {
	api.InitializeRouter(cfg.AppEnv)
	api.RunServer(cfg.Port)
}

func main() {
	env.Parse(&cfg)

	log.Info("Go Environment: " + cfg.AppEnv)

	raven.CapturePanic(initDB, nil)
	raven.CapturePanic(initServer, nil)

	defer db.Instance().Close()
}
