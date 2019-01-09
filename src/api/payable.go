package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	db "github.com/gauravagarwalr/yaes-server/src/config/db"
	model "github.com/gauravagarwalr/yaes-server/src/models"
	"github.com/gorilla/mux"
)

func GetPayablesHandler(w http.ResponseWriter, req *http.Request) {
	user := req.Context().Value(loggedInUserKey).(model.User)
	var payables []model.Payable

	db.Instance().Model(&user).Related(&payables, "Payables")

	json.NewEncoder(w).Encode(payables)
}

func UpdatePayableHandler(w http.ResponseWriter, req *http.Request) {
	user := req.Context().Value(loggedInUserKey).(model.User)
	payableID, err := strconv.ParseUint(mux.Vars(req)["payableID"], 10, 32)

	var payable model.Payable
	dbErr := db.Instance().Preload("Expense").Where("id = ?", payableID).First(&payable).Error

	if err == nil {
		err = dbErr
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if payable.Expense.CreatedBy != user.ID {
		http.Error(w, "Not Authorized", http.StatusUnprocessableEntity)
		return
	}

	json.NewDecoder(req.Body).Decode(&payable)
	payable.ID = uint(payableID)
	db.Instance().Save(&payable)

	json.NewEncoder(w).Encode(payable)
}

// TODO: Add handler for POST /payables
