package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"algogrit.com/yaes-server/entities"
	"algogrit.com/yaes-server/internal/config"
	"algogrit.com/yaes-server/payables/service"
	httpError "algogrit.com/yaes-server/pkg/http_error"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

// Handler represents a payable service handler
type Handler struct {
	ps       service.PayableService
	jwtChain alice.Chain
	*mux.Router
}

func (h *Handler) index(w http.ResponseWriter, req *http.Request) {
	user := req.Context().Value(config.LoggedInUser).(entities.User)

	payables, err := h.ps.Index(req.Context(), user)

	if err != nil {
		httpError.Write(w, err)
		return
	}

	json.NewEncoder(w).Encode(payables)
}

func (h *Handler) update(w http.ResponseWriter, req *http.Request) {
	user := req.Context().Value(config.LoggedInUser).(entities.User)

	payableID, err := strconv.ParseUint(mux.Vars(req)["payableID"], 10, 32)

	if err != nil {
		httpError.Write(w, err)
		return
	}

	var payable entities.Payable
	err = json.NewDecoder(req.Body).Decode(&payable)

	if err != nil {
		httpError.Write(w, err)
		return
	}

	payable.ID = uint(payableID)

	updatedPayable, err := h.ps.Update(req.Context(), user, payable)

	if err != nil {
		httpError.Write(w, err)
		return
	}

	json.NewEncoder(w).Encode(updatedPayable)
}

// Setup routes on an existing Router instance
func (h *Handler) Setup(r *mux.Router) {
	r.Handle("/payables", h.jwtChain.ThenFunc(h.index)).Methods("GET")

	r.Handle("/payables/{payableID}", h.jwtChain.ThenFunc(h.update)).Methods("PUT")
}

// NewHTTPHandler create a new http.Handler
func NewHTTPHandler(ps service.PayableService, jwtChain alice.Chain) Handler {
	h := Handler{ps: ps, jwtChain: jwtChain}

	h.Router = mux.NewRouter()

	h.Setup(h.Router)

	return h
}
