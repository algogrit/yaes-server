package main

import (
	"net/http"

	"algogrit.com/yaes-server/internal/config"
	"algogrit.com/yaes-server/internal/db"
	"algogrit.com/yaes-server/pkg/routes"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"

	healthService "algogrit.com/yaes-server/health/service"

	userRepo "algogrit.com/yaes-server/users/repository"
	userService "algogrit.com/yaes-server/users/service"

	expenseRepo "algogrit.com/yaes-server/expenses/repository"
	expenseService "algogrit.com/yaes-server/expenses/service"

	payableRepo "algogrit.com/yaes-server/payables/repository"
	payableService "algogrit.com/yaes-server/payables/service"
)

func main() {
	cfg := config.New()

	err := cfg.Validate()

	if err != nil {
		log.Fatal(err)
	}

	dbInstance := db.New(cfg.AppEnv, cfg.DBUrl, cfg.DBName)

	ur := userRepo.New(dbInstance)
	us := userService.New(ur)

	er := expenseRepo.New(dbInstance)
	es := expenseService.New(er)

	pr := payableRepo.New(dbInstance)
	ps := payableService.New(pr)

	r := routes.New(us, es, ps)

	hs := healthService.New(cfg.AppEnv)

	r.HandleFunc("/", hs.Healthz).Methods("GET")
	r.HandleFunc("/healthz", hs.Healthz).Methods("GET")

	r.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":"+cfg.Port, r)
}
