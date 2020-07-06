package repository

import (
	"algogrit.com/yaes-server/entities"
)

type UserRepository interface {
	RetrieveOthers(entities.User) ([]*entities.User, error)
	FindBy(string) (*entities.User, error)
	FindByID(interface{}) (*entities.User, error)
	Save(entities.User) (*entities.User, error)
}
