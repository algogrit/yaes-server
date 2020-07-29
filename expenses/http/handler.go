package http

import (
	"encoding/json"
	"net/http"

	"algogrit.com/yaes-server/entities"
	"algogrit.com/yaes-server/expenses/service"
	"algogrit.com/yaes-server/internal/config"
	httpError "algogrit.com/yaes-server/pkg/http_error"
	"algogrit.com/yaes-server/pkg/metrics"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

// Handler represents a payable service handler
type Handler struct {
	es       service.ExpenseService
	jwtChain alice.Chain
	*mux.Router
	observer metrics.HTTPSummary
}

func (h *Handler) index(w http.ResponseWriter, req *http.Request) {
	user := req.Context().Value(config.LoggedInUser).(entities.User)

	expenses, err := h.es.Index(req.Context(), user)

	if err != nil {
		httpError.Write(w, err)
		return
	}

	json.NewEncoder(w).Encode(expenses)
}

func (h *Handler) create(w http.ResponseWriter, req *http.Request) {
	user := req.Context().Value(config.LoggedInUser).(entities.User)

	var expense entities.Expense

	json.NewDecoder(req.Body).Decode(&expense)

	createdExpense, err := h.es.Create(req.Context(), user, expense)

	if err != nil {
		httpError.Write(w, err)
	}

	json.NewEncoder(w).Encode(createdExpense)
}

// Setup routes on an existing Router instance
func (h *Handler) Setup(r *mux.Router) {
	commonChain := alice.New(h.observer.Middleware)

	r.Handle("/expenses", commonChain.Extend(h.jwtChain).ThenFunc(h.index)).Methods("GET")
	r.Handle("/expenses", commonChain.Extend(h.jwtChain).ThenFunc(h.create)).Methods("POST")
}

// NewHTTPHandler create a new http.Handler
func NewHTTPHandler(ps service.ExpenseService, jwtChain alice.Chain) Handler {
	h := Handler{es: ps, jwtChain: jwtChain}

	h.observer = metrics.NewHTTPSummary(config.MetricsNamespace, "expenses")

	h.Router = mux.NewRouter()
	h.Setup(h.Router)

	return h
}
