package main

import (
	"net/http"

	"algogrit.com/yaes-server/internal/config"
	"algogrit.com/yaes-server/internal/db"
	"algogrit.com/yaes-server/pkg/auth"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"

	healthService "algogrit.com/yaes-server/health/service"

	userHTTP "algogrit.com/yaes-server/users/http"
	userRepo "algogrit.com/yaes-server/users/repository"
	userService "algogrit.com/yaes-server/users/service"

	expenseHTTP "algogrit.com/yaes-server/expenses/http"
	expenseRepo "algogrit.com/yaes-server/expenses/repository"
	expenseService "algogrit.com/yaes-server/expenses/service"

	payableHTTP "algogrit.com/yaes-server/payables/http"
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

	dbInstance := db.New(cfg)

	ur := userRepo.New(dbInstance)
	jwtChain := auth.New(ur, cfg.JWTSigningKey).Middleware()

	us := userService.New(ur, cfg.JWTSigningKey)
	userHandler := userHTTP.NewHTTPHandler(us, jwtChain)

	er := expenseRepo.New(dbInstance)
	es := expenseService.New(er)
	expenseHandler := expenseHTTP.NewHTTPHandler(es, jwtChain)

	pr := payableRepo.New(dbInstance)
	ps := payableService.New(pr)
	payableHandler := payableHTTP.NewHTTPHandler(ps, jwtChain)

	payableHandler.Setup(userHandler.Router)
	expenseHandler.Setup(userHandler.Router)

	n := negroni.Classic()
	n.UseHandler(userHandler.Router)

	go startDiagnosticsServer(cfg)

	log.Infof("Starting server on port: %s in %s mode\n", cfg.Port, cfg.AppEnv)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, n))
}
