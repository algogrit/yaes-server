package routes

import (
	"net/http"

	expenseService "algogrit.com/yaes-server/expenses/service"
	payableService "algogrit.com/yaes-server/payables/service"
	"algogrit.com/yaes-server/pkg/auth"
	userService "algogrit.com/yaes-server/users/service"

	"github.com/gorilla/mux"
)

type Router struct {
	*mux.Router

	us   userService.UserService
	es   expenseService.ExpenseService
	ps   payableService.PayableService
	auth auth.Auth
}

func New(us userService.UserService,
	es expenseService.ExpenseService,
	ps payableService.PayableService,
	auth auth.Auth) Router {

	r := mux.NewRouter()
	routes := Router{r, us, es, ps, auth}

	routes.initRoutes()

	return routes
}

func (r *Router) initRoutes() {
	r.HandleFunc("/users", r.us.Create).Methods("POST")
	r.HandleFunc("/login", r.us.Login).Methods("POST")

	r.Handle("/users", r.auth.Middleware().Then((http.HandlerFunc(r.us.Index)))).Methods("GET")

	r.Handle("/expenses", r.auth.Middleware().Then((http.HandlerFunc(r.es.Create)))).Methods("POST")
	r.Handle("/expenses", r.auth.Middleware().Then((http.HandlerFunc(r.es.Index)))).Methods("GET")

	r.Handle("/payables", r.auth.Middleware().Then((http.HandlerFunc(r.ps.Index)))).Methods("GET")
	r.Handle("/payables/{payableID}", r.auth.Middleware().Then((http.HandlerFunc(r.ps.Update)))).Methods("PUT")
}
