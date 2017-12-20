package api

import (
	"encoding/json"
	"net/http"

	db "github.com/gauravagarwalr/Yet-Another-Expense-Splitter/src/config/db"
	model "github.com/gauravagarwalr/Yet-Another-Expense-Splitter/src/models"
)

const loggedInUserKey = "LoggedInUser"

var jwtSigningKey = []byte("483175006c1088c849502ef22406ac4e")

func CreateUserHandler(w http.ResponseWriter, req *http.Request) {
	var creds = make(map[string]interface{})

	json.NewDecoder(req.Body).Decode(&creds)

	user := model.User{
		Username:     creds["username"].(string),
		FirstName:    creds["firstName"].(string),
		LastName:     creds["lastName"].(string),
		MobileNumber: creds["mobileNumber"].(string)}
	user.HashedPassword = model.HashAndSalt(creds["password"].(string))

	if err := db.Instance().Create(&user).Error; err != nil {
		http.Error(w, err.Error(), unprocessableEntity)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func CreateSessionHandler(w http.ResponseWriter, req *http.Request) {
	var creds = make(map[string]string)

	json.NewDecoder(req.Body).Decode(&creds)

	var user model.User
	db.Instance().Where("username = ?", creds["username"]).First(&user)

	if model.ComparePasswords(user.HashedPassword, creds["password"]) {
		tokenMap := model.CreateJWTToken(user, jwtSigningKey)

		json.NewEncoder(w).Encode(tokenMap)
	} else {
		http.Error(w, "Not Authorized", unauthorized)
		return
	}
}

func GetUsersHandler(w http.ResponseWriter, req *http.Request) {
	user := req.Context().Value(loggedInUserKey).(model.User)

	var users []model.User
	db.Instance().Where("id != ?", user.ID).Find(&users)

	json.NewEncoder(w).Encode(users)
}
