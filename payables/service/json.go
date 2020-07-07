package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"algogrit.com/yaes-server/entities"
	"algogrit.com/yaes-server/internal/config"
	"algogrit.com/yaes-server/payables/repository"
	"github.com/gorilla/mux"
)

type payableService struct {
	repository.PayableRepository
}

func (ps *payableService) Index(w http.ResponseWriter, req *http.Request) {
	user := req.Context().Value(config.LoggedInUser).(entities.User)

	payables, err := ps.RetrieveBy(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	json.NewEncoder(w).Encode(payables)
}

func (ps *payableService) Update(w http.ResponseWriter, req *http.Request) {
	user := req.Context().Value(config.LoggedInUser).(entities.User)

	payableID, err := strconv.ParseUint(mux.Vars(req)["payableID"], 10, 32)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	payable, err := ps.FindBy(payableID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if payable.Expense.CreatedBy != user.ID {
		http.Error(w, "Not Authorized", http.StatusUnprocessableEntity)
		return
	}

	json.NewDecoder(req.Body).Decode(payable)
	payable.ID = uint(payableID)

	if err := ps.PayableRepository.Update(payable); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}

	json.NewEncoder(w).Encode(payable)
}

// New creates a new instance of PayableService
func New(repo repository.PayableRepository) PayableService {
	return &payableService{repo}
}
