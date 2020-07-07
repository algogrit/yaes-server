package auth

import (
	"context"
	"net/http"

	"algogrit.com/yaes-server/internal/config"
	"algogrit.com/yaes-server/users/repository"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/justinas/alice"
)

type jwtAuth struct {
	repository.UserRepository
	jwtSigningKey string
}

func (j *jwtAuth) Middleware() alice.Chain {
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(j.jwtSigningKey), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	return alice.New(jwtMiddleware.Handler, j.setUser)
}

func (j *jwtAuth) setUser(h http.Handler) http.Handler {
	jwtHandler := func(w http.ResponseWriter, req *http.Request) {
		jwtToken := req.Context().Value("user").(*jwt.Token)
		userID := jwtToken.Claims.(jwt.MapClaims)["userID"]

		user, err := j.UserRepository.FindByID(userID)

		if err != nil {
			http.Error(w, "Not Authorized", http.StatusUnauthorized)
			return
		}

		newRequest := req.WithContext(context.WithValue(req.Context(), config.LoggedInUser, *user))

		h.ServeHTTP(w, newRequest)
	}

	return http.HandlerFunc(jwtHandler)
}

// New creates a new instance of jwtAuth
func New(repo repository.UserRepository, jwtSigningKey string) Auth {
	return &jwtAuth{repo, jwtSigningKey}
}
