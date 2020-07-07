package service

import "net/http"

// UserService is used for creating a user service
type UserService interface {
	Create(http.ResponseWriter, *http.Request)
	Index(http.ResponseWriter, *http.Request)
	Login(http.ResponseWriter, *http.Request)
}
