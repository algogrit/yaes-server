package main

import (
	"encoding/json"
	"log"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gauravagarwalr/Yet-Another-Expense-Splitter/src/models"
	"github.com/gorilla/mux"
)

var jwtSigningKey = []byte("483175006c1088c849502ef22406ac4e")

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
		tokenMap := model.CreateJWTToken(user, jwtSigningKey)

		json.NewEncoder(w).Encode(tokenMap)
	} else {
		http.Error(w, "Not Authorized", Unauthorized)
		return
	}
}

func createExpenseHandler(w http.ResponseWriter, req *http.Request) {
	jwtToken := req.Context().Value("user").(*jwt.Token)

	user := model.FindUserFromToken(jwtToken, db)

	var expense model.Expense

	json.NewDecoder(req.Body).Decode(&expense)
	expense.User = user

	if err := db.Create(&expense).Error; err != nil {
		http.Error(w, err.Error(), Unprocessable_Entity)
		return
	}

	json.NewEncoder(w).Encode(expense)
}

func runServer(port string) {
	router := mux.NewRouter()
	router.HandleFunc("/users", createUserHandler).Methods("POST")
	router.HandleFunc("/login", createSessionHandler).Methods("POST")

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return jwtSigningKey, nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	NegroniRoute(router, "/expenses", "POST", createExpenseHandler, jwtMiddleware.HandlerWithNext)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
