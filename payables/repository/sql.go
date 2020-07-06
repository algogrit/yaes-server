package repository

import (
	"algogrit.com/yaes-server/entities"
	"github.com/jinzhu/gorm"
)

type payableRepository struct {
	*gorm.DB
}

func (pr *payableRepository) RetrieveBy(u entities.User) ([]*entities.Payable, error) {
	var payables []*entities.Payable
	err := pr.Model(&u).Related(&payables, "Payables").Error

	return payables, err
}

func (pr *payableRepository) FindBy(payableID uint64) (*entities.Payable, error) {
	var payable entities.Payable
	err := pr.Preload("Expense").Where("id = ?", payableID).First(&payable).Error

	return &payable, err
}

func (pr *payableRepository) Update(payable *entities.Payable) error {
	return pr.Save(payable).Error
}

// New creates an instance of PayableRepository
func New(db *gorm.DB) PayableRepository {
	return &payableRepository{db}
}
