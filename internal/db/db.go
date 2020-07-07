package db

import (
	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var databaseMap = map[string]string{
	"development": "yaes-dev",
	"test":        "yaes-test",
	"production":  "yaes",
}

func getConnectionString(dbURL string, dbName string) string {
	var dbConnectionString string

	if dbURL != "" {
		dbConnectionString = dbURL
	} else {
		dbConnectionString = "dbname=" + dbName + " sslmode=disable"
	}

	log.Info("DB connection string: " + dbConnectionString)

	return dbConnectionString
}

// New returns an instance of db connection
func New(goAppEnvironment string, dbURL string, dbName string) *gorm.DB {
	dbConnectionString := getConnectionString(dbURL, dbName)

	localDB, err := gorm.Open("postgres", dbConnectionString)

	if err != nil {
		log.Fatal("Database connection error. Could not connect to database: ", dbConnectionString, ". ", err)
	}

	localDB.LogMode(goAppEnvironment == "development")

	if goAppEnvironment == "production" {
		localDB.DB().SetMaxIdleConns(4)
		localDB.DB().SetMaxOpenConns(20)
	}

	return localDB
}
