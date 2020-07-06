package routes

import (
	expenseService "algogrit.com/yaes-server/expenses/service"
	payableService "algogrit.com/yaes-server/payables/service"
	userService "algogrit.com/yaes-server/users/service"

	"github.com/gorilla/mux"
)

type Router struct {
	*mux.Router

	us userService.UserService
	es expenseService.ExpenseService
	ps payableService.PayableService
}

func New(us userService.UserService,
	es expenseService.ExpenseService,
	ps payableService.PayableService) Router {

	r := mux.NewRouter()
	routes := Router{r, us, es, ps}

	routes.initRoutes()

	return routes
}

func (r *Router) initRoutes() {
	r.HandleFunc("/users", r.us.Create).Methods("POST")
	r.HandleFunc("/login", r.us.Login).Methods("POST")

	r.HandleFunc("/users", r.us.Index).Methods("GET")

	r.HandleFunc("/expenses", r.es.Create).Methods("POST")
	r.HandleFunc("/expenses", r.es.Index).Methods("GET")

	r.HandleFunc("/payables", r.ps.Index).Methods("GET")
	r.HandleFunc("/payables/{payableID}", r.ps.Update).Methods("PUT")
}
