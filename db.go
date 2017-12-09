package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Db *gorm.DB

var databaseMap = map[string]string{
	"development": "yaes-dev",
	"testing":     "yaes-test",
	"production":  "yaes",
}

type User struct {
	gorm.Model
	Username       string `gorm:"not null;unique"`
	HashedPassword string `json:"-" gorm:"not null"`
	FirstName      string
	LastName       string
	MobileNumber   string `gorm:"not null;unique"`
}

func migration() {
	Db.AutoMigrate(&User{})

	addCheckForEmptyUsername := "ALTER TABLE users ADD CONSTRAINT check_empty_username CHECK (username <> '');"
	Db.Exec(addCheckForEmptyUsername)

	addCheckForEmptyMobileNumber := "ALTER TABLE users ADD CONSTRAINT check_empty_mobile_number CHECK (mobile_number <> '');"
	Db.Exec(addCheckForEmptyMobileNumber)
}

func initializeDB(goLangEnvironment string) {
	dbName, ok := databaseMap[goLangEnvironment]

	if !ok {
		dbName = "yaes-dev"
	}

	localDB, err := gorm.Open("postgres", "dbname="+dbName+" sslmode=disable")

	if err != nil {
		panic("failed to connect database")
	}

	localDB.LogMode(goLangEnvironment != "production")

	Db = localDB

	// Migrate the schema
	migration()
}
