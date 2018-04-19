package api

import (
	"encoding/json"
	"net/http"

	db "github.com/gauravagarwalr/yaes-server/src/config/db"
	model "github.com/gauravagarwalr/yaes-server/src/models"
)

func CreateExpenseHandler(w http.ResponseWriter, req *http.Request) {
	user := req.Context().Value(loggedInUserKey).(model.User)

	var expense model.Expense

	json.NewDecoder(req.Body).Decode(&expense)
	expense.User = user

	if err := db.Instance().Create(&expense).Error; err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	json.NewEncoder(w).Encode(expense)
}

func GetExpensesHandler(w http.ResponseWriter, req *http.Request) {
	user := req.Context().Value(loggedInUserKey).(model.User)
	var expenses []model.Expense

	db.Instance().Preload("Payables").Model(&user).Related(&expenses, "Expenses")

	json.NewEncoder(w).Encode(expenses)
}
