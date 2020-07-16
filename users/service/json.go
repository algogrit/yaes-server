package service

import (
	"encoding/json"
	"net/http"

	"algogrit.com/yaes-server/entities"
	"algogrit.com/yaes-server/internal/config"
	"algogrit.com/yaes-server/users/repository"
)

type userService struct {
	repository.UserRepository
	jwtSigningKey string
}

func (us *userService) Create(w http.ResponseWriter, req *http.Request) {
	nUser := new(createUserForm)
	json.NewDecoder(req.Body).Decode(&nUser)

	user := entities.User{
		Username:     nUser.Username,
		FirstName:    nUser.FirstName,
		LastName:     nUser.LastName,
		MobileNumber: nUser.MobileNumber,
	}

	user.SetPassword(nUser.Password)

	createdUser, err := us.Save(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	json.NewEncoder(w).Encode(createdUser)
}

func (us *userService) Index(w http.ResponseWriter, req *http.Request) {
	user := req.Context().Value(config.LoggedInUser).(entities.User)

	users, err := us.RetrieveOthers(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if users == nil {
		users = []*entities.User{}
	}

	json.NewEncoder(w).Encode(users)
}

func (us *userService) Login(w http.ResponseWriter, req *http.Request) {
	var creds loginForm
	json.NewDecoder(req.Body).Decode(&creds)

	user, err := us.FindBy(creds.Username)

	if err != nil || !user.MatchPassword(creds.Password) {
		http.Error(w, "Not Authorized", http.StatusUnauthorized)
		return
	}

	token, err := user.NewJWT(us.jwtSigningKey)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}

	tokenMap := map[string]string{"token": token}

	json.NewEncoder(w).Encode(tokenMap)
}

// New creates a new instance of UserService
func New(repo repository.UserRepository, jwtSigningKey string) UserService {
	return &userService{repo, jwtSigningKey}
}
