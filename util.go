package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"golang.org/x/crypto/bcrypt"
)

const Unauthorized = 401
const Unprocessable_Entity = 422

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func hashAndSalt(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd))
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func NegroniRoute(m *mux.Router,
	path string,
	pathType string,
	f func(http.ResponseWriter, *http.Request), // Your Route Handler
	mids ...func(http.ResponseWriter, *http.Request, http.HandlerFunc), // Middlewares
) {
	_routes := mux.NewRouter()
	_routes.HandleFunc(path, f).Methods(pathType)

	_n := negroni.New()
	for _, mid := range mids {
		_n.Use(negroni.HandlerFunc(mid))
	}

	_n.UseHandler(_routes)
	m.Handle(path, _n)
}
