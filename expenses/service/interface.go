package service

import "net/http"

type ExpenseService interface {
	Create(http.ResponseWriter, *http.Request)
	Index(http.ResponseWriter, *http.Request)
}
