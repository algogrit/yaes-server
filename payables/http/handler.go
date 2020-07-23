package http

import (
	"net/http"

	"algogrit.com/yaes-server/payables/service"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

// Handler represents a payable service handler
type Handler struct {
	ps       service.PayableService
	jwtChain alice.Chain
	*mux.Router
}

// Setup routes on an existing Router instance
func (h Handler) Setup(r *mux.Router) {
	r.Handle("/payables", h.jwtChain.Then((http.HandlerFunc(h.ps.Index)))).Methods("GET")

	r.Handle("/payables/{payableID}", h.jwtChain.Then((http.HandlerFunc(h.ps.Update)))).Methods("PUT")
}

// NewHTTPHandler create a new http.Handler
func NewHTTPHandler(ps service.PayableService, jwtChain alice.Chain) Handler {
	h := Handler{ps: ps, jwtChain: jwtChain}

	h.Router = mux.NewRouter()

	h.Setup(h.Router)

	return h
}
