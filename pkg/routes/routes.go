package routes

import (
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
}
