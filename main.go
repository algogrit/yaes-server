package main

func main() {
	goLangEnvironment := getenv("GO_APP_ENV", "production")

	initializeDB(goLangEnvironment)

	defer Db.Close()
	runServer()
}
