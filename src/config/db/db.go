package db

import (
	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/gauravagarwalr/yaes-server/src/models"
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

func InitializeDB(goAppEnvironment string, dbUrl string, dbName string) {
	var dbConnectionString string

	if dbUrl != "" {
		dbConnectionString = dbUrl + dbName
	} else {
		dbConnectionString = "dbname=" + dbName + " sslmode=disable"
	}

	log.Info("DB connection string: " + dbConnectionString)

	localDb, err := gorm.Open("postgres", dbConnectionString)

	if err != nil {
		log.Fatal("Database connection error. Could not connect to database: ", dbConnectionString, ". ", err)
	}

	localDb.LogMode(goAppEnvironment == "development")

	if goAppEnvironment == "production" {
		localDb.DB().SetMaxIdleConns(4)
		localDb.DB().SetMaxOpenConns(20)
	}

	dbInstance = localDb

	// Migrate the schema
	migration()
}

func Instance() *gorm.DB {
	return dbInstance
}
