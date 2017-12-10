package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func createUserHandler(w http.ResponseWriter, req *http.Request) {
	var creds = make(map[string]interface{})

	json.NewDecoder(req.Body).Decode(&creds)

	user := User{
		Username:     creds["username"].(string),
		FirstName:    creds["firstName"].(string),
		LastName:     creds["lastName"].(string),
		MobileNumber: creds["mobileNumber"].(string)}
	user.HashedPassword = hashAndSalt(creds["password"].(string))

	if err := Db.Create(&user).Error; err != nil {
		http.Error(w, err.Error(), UNPROCESSABLE_ENTITY)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func createSessionHandler(w http.ResponseWriter, req *http.Request) {
	var creds = make(map[string]string)

	json.NewDecoder(req.Body).Decode(&creds)

	var user User
	Db.Where("username = ?", creds["username"]).First(&user)

	if comparePasswords(user.HashedPassword, creds["password"]) {
		tokenMap := createJWTToken(user)

		json.NewEncoder(w).Encode(tokenMap)
	} else {
		http.Error(w, "Not Authorized", UNAUTHORIZED)
		return
	}
}

func runServer() {
	router := mux.NewRouter()
	router.HandleFunc("/users", createUserHandler).Methods("POST")
	router.HandleFunc("/login", createSessionHandler).Methods("POST")

	port := getenv("PORT", "12345")

	log.Fatal(http.ListenAndServe(":"+port, router))
}
