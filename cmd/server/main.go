package main

import (
	"net/http"

	"algogrit.com/yaes-server/internal/config"
	"algogrit.com/yaes-server/internal/db"
	"algogrit.com/yaes-server/pkg/auth"
	"algogrit.com/yaes-server/pkg/routes"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
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

func startDiagnosticsServer(cfg config.Config) {
	r := mux.NewRouter()
	hs := healthService.New(cfg.AppEnv)

	r.HandleFunc("/", hs.Healthz).Methods("GET")
	r.HandleFunc("/healthz", hs.Healthz).Methods("GET")
	r.Handle("/metrics", promhttp.Handler())

	n := negroni.Classic()
	n.UseHandler(r)

	log.Info("Starting diagnostics server on port: ", cfg.DiagnosticsPort)
	http.ListenAndServe(":"+cfg.DiagnosticsPort, n)
}

func main() {
	cfg := config.New()

	err := cfg.Validate()

	if err != nil {
		log.Fatal(err)
	}

	dbInstance := db.New(cfg.AppEnv, cfg.DBUrl, cfg.DBName)

	ur := userRepo.New(dbInstance)
	us := userService.New(ur, cfg.JWTSigningKey)

	er := expenseRepo.New(dbInstance)
	es := expenseService.New(er)

	pr := payableRepo.New(dbInstance)
	ps := payableService.New(pr)

	jwtChain := auth.New(ur, cfg.JWTSigningKey).Middleware()
	appRouter := routes.New(us, es, ps, jwtChain)

	n := negroni.Classic()
	n.UseHandler(appRouter)

	go startDiagnosticsServer(cfg)

	log.Info("Starting server on port: ", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, n))
}
