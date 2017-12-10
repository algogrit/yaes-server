package main

var jwtSigningKey = []byte("483175006c1088c849502ef22406ac4e")

func main() {
	goLangEnvironment := getenv("GO_APP_ENV", "production")

	initializeDB(goLangEnvironment)

	defer Db.Close()
	runServer()
}
