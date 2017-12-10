package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/gauravagarwalr/Yet-Another-Expense-Splitter/src/models"
)

var databaseMap = map[string]string{
	"development": "yaes-dev",
	"testing":     "yaes-test",
	"production":  "yaes"}

func migration(db *gorm.DB) {
	db.AutoMigrate(&model.User{})

	addCheckForEmptyUsername := "ALTER TABLE users ADD CONSTRAINT check_empty_username CHECK (username <> '');"
	db.Exec(addCheckForEmptyUsername)

	addCheckForEmptyMobileNumber := "ALTER TABLE users ADD CONSTRAINT check_empty_mobile_number CHECK (mobile_number <> '');"
	db.Exec(addCheckForEmptyMobileNumber)
}

func initializeDB(goLangEnvironment string) *gorm.DB {
	dbName, ok := databaseMap[goLangEnvironment]

	if !ok {
		dbName = "yaes-dev"
	}

	db, err := gorm.Open("postgres", "dbname="+dbName+" sslmode=disable")

	if err != nil {
		panic("failed to connect database")
	}

	db.LogMode(goLangEnvironment != "production")

	// Migrate the schema
	migration(db)

	return db
}
