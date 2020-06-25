package api

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/algogrit/raven-go"
	"github.com/algogrit/yaes-server/src/config/db"
	model "github.com/algogrit/yaes-server/src/models"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

var appEnvironment string
var routerInstance *negroni.Negroni

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

	ravenHandler := raven.RecoveryHandler(f)

	n.UseHandler(ravenHandler)

	m.Handle(path, n).Methods(pathType)
}

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

func HealthHandler(w http.ResponseWriter, req *http.Request) {
	health := make(map[string]string)

	health["Hostname"], _ = os.Hostname()
	health["Description"] = "API for the yaes mobile app."
	health["GO_APP_ENV"] = appEnvironment
	health["Host"] = req.Host
	health["End-Point"] = req.URL.Path

	json.NewEncoder(w).Encode(health)
}

func InitializeRouter(goAppEnvironment string) {
	appEnvironment = goAppEnvironment
	router := mux.NewRouter()

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return jwtSigningKey, nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	router.HandleFunc("/", raven.RecoveryHandler(HealthHandler)).Methods("GET")
	router.HandleFunc("/healthz", raven.RecoveryHandler(HealthHandler)).Methods("GET")

	wrapHandler(router, "/users", "POST", CreateUserHandler)

	wrapHandler(router, "/login", "POST", CreateSessionHandler)

	wrapHandler(router, "/users", "GET", GetUsersHandler, jwtMiddleware.HandlerWithNext, userLogInHandlerWithNext)
	wrapHandler(router, "/expenses", "POST", CreateExpenseHandler, jwtMiddleware.HandlerWithNext, userLogInHandlerWithNext)
	wrapHandler(router, "/expenses", "GET", GetExpensesHandler, jwtMiddleware.HandlerWithNext, userLogInHandlerWithNext)
	wrapHandler(router, "/payables", "GET", GetPayablesHandler, jwtMiddleware.HandlerWithNext, userLogInHandlerWithNext)
	// TODO: wrapHandler(router, "/expenses/{expenseID}/payables", "POST", PostPayablesHandler, jwtMiddleware.HandlerWithNext, userLogInHandlerWithNext)
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
