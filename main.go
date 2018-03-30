package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	api "github.com/gauravagarwalr/Yet-Another-Expense-Splitter/src/api"
	db "github.com/gauravagarwalr/Yet-Another-Expense-Splitter/src/config/db"
)

var goAppEnvironment string

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func initDB() {
	dbURL := os.Getenv("DATABASE_URL")
	dbName := os.Getenv("DB_NAME")

	if len(dbURL) == 0 && len(dbName) == 0 {
		log.Fatal("No databases provided! You can set it using DATABASE_URL or DB_NAME env variable.")
	}

	db.InitializeDB(goAppEnvironment, dbURL, dbName)
}

func initServer() {
	port := getenv("PORT", "12345")

	api.InitializeRouter(goAppEnvironment)
	api.RunServer(port)
}

func main() {
	goAppEnvironment = getenv("GO_APP_ENV", "production")

	log.Info("Go Environment: " + goAppEnvironment)

	initDB()
	initServer()

	defer db.Instance().Close()
}
