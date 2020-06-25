package api

import (
	"encoding/json"
	"net/http"

	db "algogrit.com/yaes-server/src/config/db"
	model "algogrit.com/yaes-server/src/models"
)

const loggedInUserKey = "LoggedInUser"

var jwtSigningKey = []byte("483175006c1088c849502ef22406ac4e")

type newUser struct {
	Username     string
	FirstName    string
	LastName     string
	MobileNumber string
	Password     string
}

func CreateUserHandler(w http.ResponseWriter, req *http.Request) {
	var newUser newUser
	json.NewDecoder(req.Body).Decode(&newUser)

	user := model.User{
		Username:     newUser.Username,
		FirstName:    newUser.FirstName,
		LastName:     newUser.LastName,
		MobileNumber: newUser.MobileNumber}
	user.HashedPassword = model.HashAndSalt(newUser.Password)

	if err := db.Instance().Create(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	json.NewEncoder(w).Encode(user)
}

type credentials struct {
	Username string
	Password string
}

func CreateSessionHandler(w http.ResponseWriter, req *http.Request) {
	var creds credentials
	json.NewDecoder(req.Body).Decode(&creds)

	var user model.User
	db.Instance().Where("username = ?", creds.Username).First(&user)

	if model.ComparePasswords(user.HashedPassword, creds.Password) {
		tokenMap := model.CreateJWTToken(user, jwtSigningKey)

		json.NewEncoder(w).Encode(tokenMap)
	} else {
		http.Error(w, "Not Authorized", http.StatusUnauthorized)
		return
	}
}

func GetUsersHandler(w http.ResponseWriter, req *http.Request) {
	user := req.Context().Value(loggedInUserKey).(model.User)

	var users []model.User
	db.Instance().Where("id != ?", user.ID).Find(&users)

	json.NewEncoder(w).Encode(users)
}
