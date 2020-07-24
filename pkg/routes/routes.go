package routes

import (
	expenseService "algogrit.com/yaes-server/expenses/service"
	userService "algogrit.com/yaes-server/users/service"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

// Router struct implements http.Handler through mux.Router
type Router struct {
	*mux.Router

	us       userService.UserService
	jwtChain alice.Chain
}

func (r *Router) initRoutes() {
	r.HandleFunc("/users", r.us.Create).Methods("POST")
	r.HandleFunc("/login", r.us.Login).Methods("POST")

	r.Handle("/users", r.jwtChain.ThenFunc(r.us.Index)).Methods("GET")
}

// New initializes the Router
func New(us userService.UserService,
	es expenseService.ExpenseService,
	jwtChain alice.Chain) Router {

	r := mux.NewRouter()
	routes := Router{r, us, jwtChain}

	routes.initRoutes()

	return routes
}
