package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	api "github.com/gauravagarwalr/Yet-Another-Expense-Splitter/src/api"
	db "github.com/gauravagarwalr/Yet-Another-Expense-Splitter/src/config/db"
)

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func main() {
	goAppEnvironment := getenv("GO_APP_ENV", "production")
	dbUrl := os.Getenv("DATABASE_URL")
	dbName := os.Getenv("DB_NAME")
	port := getenv("PORT", "12345")

	log.Info("Go Environment: " + goAppEnvironment)

	db.InitializeDB(goAppEnvironment, dbUrl, dbName)

	defer db.Instance().Close()

	api.InitializeRouter()
	api.RunServer(port)
}
