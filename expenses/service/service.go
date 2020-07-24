package service

import (
	"context"

	"algogrit.com/yaes-server/entities"
	"algogrit.com/yaes-server/expenses/repository"
)

type expenseService struct {
	repository.ExpenseRepository
}

func (es *expenseService) Create(ctx context.Context, user entities.User, expense entities.Expense) (*entities.Expense, error) {
	expense.User = user

	return es.Save(expense)
}

func (es *expenseService) Index(ctx context.Context, user entities.User) ([]*entities.Expense, error) {
	return es.RetrieveBy(user)
}

// New creates a new instance of ExpenseService
func New(repo repository.ExpenseRepository) ExpenseService {
	return &expenseService{repo}
}
