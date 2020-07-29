package http

import (
	"encoding/json"
	"net/http"

	"algogrit.com/yaes-server/entities"
	"algogrit.com/yaes-server/internal/config"
	httpError "algogrit.com/yaes-server/pkg/http_error"
	"algogrit.com/yaes-server/pkg/metrics"
	"algogrit.com/yaes-server/users/service"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

// Handler represents a payable service handler
type Handler struct {
	us       service.UserService
	jwtChain alice.Chain
	*mux.Router
	observer metrics.HTTPSummary
}

func (h *Handler) index(w http.ResponseWriter, req *http.Request) {
	user := req.Context().Value(config.LoggedInUser).(entities.User)

	users, err := h.us.Index(req.Context(), user)

	if err != nil {
		httpError.Write(w, err)
		return
	}

	if users == nil {
		users = []*entities.User{}
	}

	json.NewEncoder(w).Encode(users)
}

func (h *Handler) create(w http.ResponseWriter, req *http.Request) {
	nUser := new(service.CreateUserRequest)
	json.NewDecoder(req.Body).Decode(&nUser)

	createdUser, err := h.us.Create(req.Context(), *nUser)

	if err != nil {
		httpError.Write(w, err)
		return
	}

	json.NewEncoder(w).Encode(createdUser)
}

func (h *Handler) login(w http.ResponseWriter, req *http.Request) {
	var creds service.LoginRequest
	json.NewDecoder(req.Body).Decode(&creds)

	token, err := h.us.Login(req.Context(), creds)

	if err != nil {
		httpError.Write(w, err)
		return
	}

	json.NewEncoder(w).Encode(token)
}

// Setup routes on an existing Router instance
func (h *Handler) Setup(r *mux.Router) {
	commonChain := alice.New(h.observer.Middleware)

	r.Handle("/users", commonChain.ThenFunc(h.create)).Methods("POST")
	r.Handle("/login", commonChain.ThenFunc(h.login)).Methods("POST")

	r.Handle("/users", commonChain.Extend(h.jwtChain).ThenFunc(h.index)).Methods("GET")
}

// NewHTTPHandler create a new http.Handler
func NewHTTPHandler(us service.UserService, jwtChain alice.Chain) Handler {
	h := Handler{us: us, jwtChain: jwtChain}

	h.observer = metrics.NewHTTPSummary(config.MetricsNamespace, "orders")

	h.Router = mux.NewRouter()
	h.Setup(h.Router)

	return h
}
