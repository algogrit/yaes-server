package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gauravagarwalr/Yet-Another-Expense-Splitter/src/models"
	"github.com/gorilla/mux"
)

func createUserHandler(w http.ResponseWriter, req *http.Request) {
	var creds = make(map[string]interface{})

	json.NewDecoder(req.Body).Decode(&creds)

	user := model.User{
		Username:     creds["username"].(string),
		FirstName:    creds["firstName"].(string),
		LastName:     creds["lastName"].(string),
		MobileNumber: creds["mobileNumber"].(string)}
	user.HashedPassword = hashAndSalt(creds["password"].(string))

	if err := db.Create(&user).Error; err != nil {
		http.Error(w, err.Error(), Unprocessable_Entity)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func createSessionHandler(w http.ResponseWriter, req *http.Request) {
	var creds = make(map[string]string)

	json.NewDecoder(req.Body).Decode(&creds)

	var user model.User
	db.Where("username = ?", creds["username"]).First(&user)

	if comparePasswords(user.HashedPassword, creds["password"]) {
		tokenMap := model.CreateJWTToken(user)

		json.NewEncoder(w).Encode(tokenMap)
	} else {
		http.Error(w, "Not Authorized", Unauthorized)
		return
	}
}

func runServer(port string) {
	router := mux.NewRouter()
	router.HandleFunc("/users", createUserHandler).Methods("POST")
	router.HandleFunc("/login", createSessionHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":"+port, router))
}
