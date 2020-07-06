package main

func migrate(instance *gorm.DB) {
	instance.AutoMigrate(&entities.User{})
	instance.AutoMigrate(&entities.Expense{})
	instance.AutoMigrate(&entities.Payable{})

	addCheckForEmptyUsername := "ALTER TABLE users ADD CONSTRAINT check_empty_username CHECK (username <> '');"
	instance.Exec(addCheckForEmptyUsername)

	addCheckForEmptyMobileNumber := "ALTER TABLE users ADD CONSTRAINT check_empty_mobile_number CHECK (mobile_number <> '');"
	instance.Exec(addCheckForEmptyMobileNumber)
}

func main() {
	cfg := config.New()

	err := cfg.Validate()

	if err != nil {
		log.Fatal(err)
	}

	dbInstance := db.New(cfg.AppEnv, dbURL, dbName)

	migrate(dbInstance)
}
