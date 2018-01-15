package main

import (
	"os"

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
	dbName := getenv("DB_NAME", "yaes")
	goAppEnvironment := getenv("GO_APP_ENV", "production")
	port := getenv("PORT", "12345")

	db.InitializeDB(dbName, goAppEnvironment)

	defer db.Instance().Close()

	api.RunServer(port)
}
