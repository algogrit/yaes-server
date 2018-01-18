package api

import (
	"context"
	"log"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gauravagarwalr/Yet-Another-Expense-Splitter/src/config/db"
	model "github.com/gauravagarwalr/Yet-Another-Expense-Splitter/src/models"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

var routerInstance *mux.Router

func userLogInHandlerWithNext(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	jwtToken := r.Context().Value("user").(*jwt.Token)
	user, err := model.FindUserFromToken(jwtToken, db.Instance())

	if err != nil {
		http.Error(w, "Not Authorized", http.StatusUnauthorized)
		return
	}

	newRequest := r.WithContext(context.WithValue(r.Context(), loggedInUserKey, user))

	next(w, newRequest)
}

func negroniRoute(m *mux.Router,
	path string,
	pathType string,
	f func(http.ResponseWriter, *http.Request), // Your Route Handler
	mids ...func(http.ResponseWriter, *http.Request, http.HandlerFunc), // Middlewares
) {
	_routes := mux.NewRouter()
	_routes.HandleFunc(path, f).Methods(pathType)

	_n := negroni.Classic()
	for _, mid := range mids {
		_n.Use(negroni.HandlerFunc(mid))
	}

	_n.UseHandler(_routes)
	m.Handle(path, _n).Methods(pathType)
}

func InitializeRouter() {
	router := mux.NewRouter()

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return jwtSigningKey, nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	negroniRoute(router, "/users", "POST", CreateUserHandler)
	negroniRoute(router, "/login", "POST", CreateSessionHandler)

	negroniRoute(router, "/users", "GET", GetUsersHandler, jwtMiddleware.HandlerWithNext, userLogInHandlerWithNext)
	negroniRoute(router, "/expenses", "POST", CreateExpenseHandler, jwtMiddleware.HandlerWithNext, userLogInHandlerWithNext)
	negroniRoute(router, "/expenses", "GET", GetExpensesHandler, jwtMiddleware.HandlerWithNext, userLogInHandlerWithNext)
	negroniRoute(router, "/payables", "GET", GetPayablesHandler, jwtMiddleware.HandlerWithNext, userLogInHandlerWithNext)
	negroniRoute(router, "/payables/{payableID}", "PUT", UpdatePayableHandler, jwtMiddleware.HandlerWithNext, userLogInHandlerWithNext)

	routerInstance = router
}

func RunServer(port string) {
	handler := cors.Default().Handler(routerInstance)

	log.Fatal(http.ListenAndServe(":"+port, handler))
}

func Instance() *mux.Router {
	return routerInstance
}
