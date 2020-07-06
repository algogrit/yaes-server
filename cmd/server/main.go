package main

import (
	"net/http"

	"algogrit.com/yaes-server/internal/config"
	"algogrit.com/yaes-server/internal/db"
	"algogrit.com/yaes-server/pkg/routes"
	log "github.com/sirupsen/logrus"

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

	r := routes.New()

	ur := userRepo.New(dbInstance)
	us := userService.New(ur)

	r.SetUserRoutes(us)

	er := expenseRepo.New(dbInstance)
	es := expenseService.New(er)

	r.SetExpenseRoutes(es)

	pr := payableRepo.New(dbInstance)
	ps := payableService.New(pr)

	r.SetPayableRoutes(ps)

	http.ListenAndServe(":"+cfg.Port, r)
}
