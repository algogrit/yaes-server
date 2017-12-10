package main

import "github.com/jinzhu/gorm"

var db *gorm.DB

func main() {
	goLangEnvironment := getenv("GO_APP_ENV", "production")
	port := getenv("PORT", "12345")

	db = initializeDB(goLangEnvironment)
	defer db.Close()

	runServer(port)
}
