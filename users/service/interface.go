package service

import "net/http"

type UserService interface {
	Create(http.ResponseWriter, *http.Request)
	Index(http.ResponseWriter, *http.Request)
	Login(http.ResponseWriter, *http.Request)

	Middleware(http.ResponseWriter, *http.Request, http.HandlerFunc)
}
