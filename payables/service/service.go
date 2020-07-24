package service

import (
	"context"

	"algogrit.com/yaes-server/entities"
	"algogrit.com/yaes-server/payables/repository"
	httpError "algogrit.com/yaes-server/pkg/http_error"
)

type payableService struct {
	repository.PayableRepository
}

func (ps *payableService) Index(ctx context.Context, currentUser entities.User) ([]*entities.Payable, error) {
	return ps.PayableRepository.RetrieveBy(currentUser)
}

func (ps *payableService) Update(ctx context.Context, currentUser entities.User, payable entities.Payable) (*entities.Payable, error) {
	existingPayable, err := ps.FindBy(payable.ID)

	if err != nil {
		return nil, httpError.NotFoundErr().Wrap(err)
	}

	if existingPayable.Expense.CreatedBy != currentUser.ID {
		return nil, httpError.UnauthorizedErr()
	}

	if err := ps.PayableRepository.Update(&payable); err != nil {
		return nil, err
	}

	return &payable, nil
}

// New creates a new instance of PayableService
func New(repo repository.PayableRepository) PayableService {
	return &payableService{repo}
}
