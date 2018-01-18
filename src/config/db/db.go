package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/gauravagarwalr/Yet-Another-Expense-Splitter/src/models"
)

var dbInstance *gorm.DB

var databaseMap = map[string]string{
	"development": "yaes-dev",
	"test":        "yaes-test",
	"production":  "yaes"}

func migration() {
	dbInstance.AutoMigrate(&model.User{})
	dbInstance.AutoMigrate(&model.Expense{})
	dbInstance.AutoMigrate(&model.Payable{})

	addCheckForEmptyUsername := "ALTER TABLE users ADD CONSTRAINT check_empty_username CHECK (username <> '');"
	dbInstance.Exec(addCheckForEmptyUsername)

	addCheckForEmptyMobileNumber := "ALTER TABLE users ADD CONSTRAINT check_empty_mobile_number CHECK (mobile_number <> '');"
	dbInstance.Exec(addCheckForEmptyMobileNumber)
}

func InitializeDB(goAppEnvironment string) {
	dbName, ok := databaseMap[goAppEnvironment]

	if !ok {
		dbName = "yaes-dev"
	}

	localDb, err := gorm.Open("postgres", "dbname="+dbName+" sslmode=disable")

	if err != nil {
		panic("failed to connect database")
	}

	localDb.LogMode(goAppEnvironment != "production")

	localDb.DB().SetMaxIdleConns(4)
	localDb.DB().SetMaxOpenConns(20)

	dbInstance = localDb

	// Migrate the schema
	migration()
}

func Instance() *gorm.DB {
	return dbInstance
}
