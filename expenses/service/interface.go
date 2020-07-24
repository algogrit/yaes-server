package service

import (
	"context"

	"algogrit.com/yaes-server/entities"
)

type ExpenseService interface {
	Index(ctx context.Context, currentUser entities.User) ([]*entities.Expense, error)
	Create(ctx context.Context, currentUser entities.User, expense entities.Expense) (*entities.Expense, error)
}
