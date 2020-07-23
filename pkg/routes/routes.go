package routes

import (
	"net/http"

	expenseService "algogrit.com/yaes-server/expenses/service"
	userService "algogrit.com/yaes-server/users/service"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

// Router struct implements http.Handler through mux.Router
type Router struct {
	*mux.Router

	us       userService.UserService
	es       expenseService.ExpenseService
	jwtChain alice.Chain
}

func (r *Router) initRoutes() {
	r.HandleFunc("/users", r.us.Create).Methods("POST")
	r.HandleFunc("/login", r.us.Login).Methods("POST")

	r.Handle("/users", r.jwtChain.Then((http.HandlerFunc(r.us.Index)))).Methods("GET")

	r.Handle("/expenses", r.jwtChain.Then((http.HandlerFunc(r.es.Create)))).Methods("POST")
	r.Handle("/expenses", r.jwtChain.Then((http.HandlerFunc(r.es.Index)))).Methods("GET")
}

// New initializes the Router
func New(us userService.UserService,
	es expenseService.ExpenseService,
	jwtChain alice.Chain) Router {

	r := mux.NewRouter()
	routes := Router{r, us, es, jwtChain}

	routes.initRoutes()

	return routes
}
