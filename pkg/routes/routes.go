package routes

import (
	expenseService "algogrit.com/yaes-server/expenses/service"
	payableService "algogrit.com/yaes-server/payables/service"
	userService "algogrit.com/yaes-server/users/service"

	"github.com/gorilla/mux"
)

type Router struct {
	*mux.Router
}

func New() Router {
	r := mux.NewRouter()
	return Router{r}
}

func (r *Router) SetUserRoutes(us userService.UserService) {
	r.HandleFunc("/users", us.Create).Methods("POST")
	r.HandleFunc("/login", us.Login).Methods("POST")

	r.HandleFunc("/users", us.Index).Methods("GET")
}

func (r *Router) SetExpenseRoutes(es expenseService.ExpenseService) {
	r.HandleFunc("/expenses", es.Create).Methods("POST")
	r.HandleFunc("/expenses", es.Index).Methods("GET")
}

func (r *Router) SetPayableRoutes(ps payableService.PayableService) {
	r.HandleFunc("/payables", ps.Index).Methods("GET")
	r.HandleFunc("/payables/{payableID}", ps.Update).Methods("PUT")
}
