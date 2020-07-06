package service

import "net/http"

type HealthService interface {
	Healthz(http.ResponseWriter, *http.Request)
}
