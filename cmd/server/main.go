package main

import (
	"net/http"
	"algogrit.com/yaes-server/internal/config"
	log "github.com/sirupsen/logrus"
	"algogrit.com/yaes-server/internal/db"
	"algogrit.com/yaes-server/pkg/routes"

	userService "algogrit.com/yaes-server/users/service"
	userRepo "algogrit.com/yaes-server/users/repository"
)

func main() {
	cfg := config.New()

	err := cfg.Validate()

	if err != nil {
		log.Fatal(err)
	}

	dbInstance := db.New(cfg.AppEnv, cfg.DBUrl, cfg.DBName)

	r := routes.New()

	ur := userRepo.New(dbInstance)
	us := userService.New(ur)

	r.SetUserRoutes(us)

	http.ListenAndServe(":"+cfg.Port, r)
}
