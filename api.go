package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gauravagarwalr/Yet-Another-Expense-Splitter/src/models"
	"github.com/gorilla/mux"
)

var jwtSigningKey = []byte("483175006c1088c849502ef22406ac4e")
var loggedInUserKey = "LoggedInUser"

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

func getUsersHandler(w http.ResponseWriter, req *http.Request) {
	user := req.Context().Value(loggedInUserKey).(model.User)

	var users []model.User
	db.Where("id != ?", user.ID).Find(&users)

	json.NewEncoder(w).Encode(users)
}

func createExpenseHandler(w http.ResponseWriter, req *http.Request) {
	user := req.Context().Value(loggedInUserKey).(model.User)

	var expense model.Expense

	json.NewDecoder(req.Body).Decode(&expense)
	expense.User = user

	if err := db.Create(&expense).Error; err != nil {
		http.Error(w, err.Error(), Unprocessable_Entity)
		return
	}

	json.NewEncoder(w).Encode(expense)
}

func getExpensesHandler(w http.ResponseWriter, req *http.Request) {
	user := req.Context().Value(loggedInUserKey).(model.User)
	var expenses []model.Expense

	db.Preload("Payables").Model(&user).Related(&expenses, "Expenses")

	json.NewEncoder(w).Encode(expenses)
}

func getPayablesHandler(w http.ResponseWriter, req *http.Request) {
	user := req.Context().Value(loggedInUserKey).(model.User)
	var payables []model.Payable

	db.Model(&user).Related(&payables, "Payables")

	json.NewEncoder(w).Encode(payables)
}

func updatePayableHandler(w http.ResponseWriter, req *http.Request) {
	// req.Context().Value(loggedInUserKey).(model.User)
}

func userLogInHandlerWithNext(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	jwtToken := r.Context().Value("user").(*jwt.Token)
	user, err := model.FindUserFromToken(jwtToken, db)

	if err != nil {
		http.Error(w, "Not Authorized", Unauthorized)
		return
	}

	newRequest := r.WithContext(context.WithValue(r.Context(), loggedInUserKey, user))

	next(w, newRequest)
}

func runServer(port string) {
	router := mux.NewRouter()

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return jwtSigningKey, nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	NegroniRoute(router, "/users", "POST", createUserHandler)
	NegroniRoute(router, "/login", "POST", createSessionHandler)

	NegroniRoute(router, "/users", "GET", getUsersHandler, jwtMiddleware.HandlerWithNext, userLogInHandlerWithNext)
	NegroniRoute(router, "/expenses", "POST", createExpenseHandler, jwtMiddleware.HandlerWithNext, userLogInHandlerWithNext)
	NegroniRoute(router, "/expenses", "GET", getExpensesHandler, jwtMiddleware.HandlerWithNext, userLogInHandlerWithNext)
	NegroniRoute(router, "/payables", "GET", getPayablesHandler, jwtMiddleware.HandlerWithNext, userLogInHandlerWithNext)
	NegroniRoute(router, "/payables/{payableID}", "PUT", updatePayableHandler, jwtMiddleware.HandlerWithNext, userLogInHandlerWithNext)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
