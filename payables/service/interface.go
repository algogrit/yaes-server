package service

import (
	"context"

	"algogrit.com/yaes-server/entities"
)

// PayableService represents a payable service
type PayableService interface {
	Index(ctx context.Context, user entities.User) ([]*entities.Payable, error)
	Update(ctx context.Context, user entities.User, payable entities.Payable) (*entities.Payable, error)
	// TODO: Add handler for POST /payables
}
