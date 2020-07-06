package service

import (
	"encoding/json"
	"net/http"

	"algogrit.com/yaes-server/entities"
	"algogrit.com/yaes-server/expenses/repository"
)

type expenseService struct {
	repository.ExpenseRepository
}

const loggedInUserKey = "LoggedInUser"

func (es *expenseService) Create(w http.ResponseWriter, req *http.Request) {
	user := req.Context().Value(loggedInUserKey).(entities.User)

	var expense entities.Expense

	json.NewDecoder(req.Body).Decode(&expense)
	expense.User = user

	createdExpense, err := es.Save(expense)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	json.NewEncoder(w).Encode(createdExpense)
}

func (es *expenseService) Index(w http.ResponseWriter, req *http.Request) {
	user := req.Context().Value(loggedInUserKey).(entities.User)

	expenses, err := es.RetrieveBy(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	json.NewEncoder(w).Encode(expenses)
}

// New creates a new instance of ExpenseService
func New(repo repository.ExpenseRepository) ExpenseService {
	return &expenseService{repo}
}
