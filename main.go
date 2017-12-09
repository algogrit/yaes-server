package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Username       string `gorm:"not null;unique"`
	HashedPassword string `json:"-" gorm:"not null"`
	FirstName      string
	LastName       string
	MobileNumber   string `gorm:"not null;unique"`
}

func hashAndSalt(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func CreateUserHandler(w http.ResponseWriter, req *http.Request) {
	body := req.Body
	decoder := json.NewDecoder(body)

	var creds interface{}
	decoder.Decode(&creds)
	cred, _ := creds.(map[string]interface{})

	user := User{
		Username:     cred["username"].(string),
		FirstName:    cred["firstName"].(string),
		LastName:     cred["lastName"].(string),
		MobileNumber: cred["mobileNumber"].(string)}
	user.HashedPassword = hashAndSalt(cred["password"].(string))

	if err := db.Create(&user).Error; err != nil {
		http.Error(w, err.Error(), 428)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func initializeDB() {
	localDB, err := gorm.Open("postgres", "dbname=yaes-dev sslmode=disable")

	if err != nil {
		panic("failed to connect database")
	}

	localDB.LogMode(true)

	db = localDB

	// Migrate the schema
	localDB.AutoMigrate(&User{})

	addCheckForEmptyUsername := "ALTER TABLE users ADD CONSTRAINT check_empty_username CHECK (username <> '');"
	db.Exec(addCheckForEmptyUsername)

	addCheckForEmptyMobileNumber := "ALTER TABLE users ADD CONSTRAINT check_empty_mobile_number CHECK (mobile_number <> '');"
	db.Exec(addCheckForEmptyMobileNumber)
}

func main() {
	initializeDB()

	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/users", CreateUserHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":12345", router))
}
