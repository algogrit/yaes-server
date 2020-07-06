package service

import "net/http"

type PayableService interface {
	Index(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	// TODO: Add handler for POST /payables
}
