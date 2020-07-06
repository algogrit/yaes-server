package repository

import "algogrit.com/yaes-server/entities"

type ExpenseRepository interface {
	RetrieveBy(entities.User) ([]*entities.Expense, error)
	Save(entities.Expense) (*entities.Expense, error)
}
