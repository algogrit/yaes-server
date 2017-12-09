package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

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

	if err := Db.Create(&user).Error; err != nil {
		http.Error(w, err.Error(), UNPROCESSABLE_ENTITY)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func runServer() {
	router := mux.NewRouter()
	router.HandleFunc("/users", CreateUserHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":12345", router))
}
