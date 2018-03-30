package api

import (
	"context"
	"net/http"

	log "github.com/sirupsen/logrus"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gauravagarwalr/Yet-Another-Expense-Splitter/src/config/db"
	model "github.com/gauravagarwalr/Yet-Another-Expense-Splitter/src/models"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

var routerInstance *negroni.Negroni

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

func wrapHandler(
	m *mux.Router,
	path string,
	pathType string,
	f http.HandlerFunc,
	mids ...func(http.ResponseWriter, *http.Request, http.HandlerFunc),
) {
	n := negroni.New()
	for _, mid := range mids {
		n.Use(negroni.HandlerFunc(mid))
	}

	n.UseHandler(f)

	m.Handle(path, n).Methods(pathType)
}

func InitializeRouter() {
	router := mux.NewRouter()

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return jwtSigningKey, nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	wrapHandler(router, "/users", "POST", CreateUserHandler)

	wrapHandler(router, "/login", "POST", CreateSessionHandler)

	wrapHandler(router, "/users", "GET", GetUsersHandler, jwtMiddleware.HandlerWithNext, userLogInHandlerWithNext)
	wrapHandler(router, "/expenses", "POST", CreateExpenseHandler, jwtMiddleware.HandlerWithNext, userLogInHandlerWithNext)
	wrapHandler(router, "/expenses", "GET", GetExpensesHandler, jwtMiddleware.HandlerWithNext, userLogInHandlerWithNext)
	wrapHandler(router, "/payables", "GET", GetPayablesHandler, jwtMiddleware.HandlerWithNext, userLogInHandlerWithNext)
	wrapHandler(router, "/payables/{payableID}", "PUT", UpdatePayableHandler, jwtMiddleware.HandlerWithNext, userLogInHandlerWithNext)

	n := negroni.Classic()
	n.UseHandler(router)

	routerInstance = n
}

func RunServer(port string) {
	handler := cors.Default().Handler(routerInstance)

	log.Fatal(http.ListenAndServe(":"+port, handler))
}

func Instance() *negroni.Negroni {
	return routerInstance
}
