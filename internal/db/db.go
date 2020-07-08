package db

import (
	"fmt"

	"algogrit.com/yaes-server/internal/config"
	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // For init connection to postgres
)

var databaseMap = map[string]string{
	"development": "yaes-dev",
	"test":        "yaes-test",
	"production":  "yaes",
}

func getConnectionString(cfg config.Config) string {
	var dbURL string

	if cfg.DBUrl != "" {
		dbURL = cfg.DBUrl
	} else {
		dbURL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBSSLMode)
	}

	log.Info("DB connection string: " + dbURL)

	return dbURL
}

// New returns an instance of db connection
func New(cfg config.Config) *gorm.DB {
	dbURL := getConnectionString(cfg)

	localDB, err := gorm.Open("postgres", dbURL)

	if err != nil {
		log.Fatal("Database connection error. Could not connect to database: ", dbURL, ". ", err)
	}

	localDB.LogMode(cfg.AppEnv == "development")

	if cfg.AppEnv == "production" {
		localDB.DB().SetMaxIdleConns(4)
		localDB.DB().SetMaxOpenConns(20)
	}

	return localDB
}
