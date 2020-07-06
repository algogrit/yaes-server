package service

import (
	"encoding/json"
	"net/http"
	"os"
)

type healthService struct {
	appEnvironment string
}

func (hs *healthService) Healthz(w http.ResponseWriter, req *http.Request) {
	health := make(map[string]string)

	health["Hostname"], _ = os.Hostname()
	health["Description"] = "API for the yaes mobile app."
	health["GO_APP_ENV"] = hs.appEnvironment
	health["Host"] = req.Host
	health["End-Point"] = req.URL.Path

	json.NewEncoder(w).Encode(health)
}

// New creates an instance of HealthService
func New(appEnvironment string) HealthService {
	return &healthService{appEnvironment}
}
