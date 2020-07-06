package repository

import (
	"algogrit.com/yaes-server/entities"
	"github.com/jinzhu/gorm"
)

type expenseRepository struct {
	*gorm.DB
}

func (er *expenseRepository) Save(expense entities.Expense) (*entities.Expense, error) {
	err := er.Create(&expense).Error

	return &expense, err
}

func (er *expenseRepository) RetrieveBy(user entities.User) ([]*entities.Expense, error) {
	var expenses []*entities.Expense

	err := er.Preload("Payables").Model(&user).Related(&expenses, "Expenses").Error

	return expenses, err
}

// New creates an instance of ExpenseRepository
func New(db *gorm.DB) ExpenseRepository {
	return &expenseRepository{db}
}
