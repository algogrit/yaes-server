package repository

import (
	"algogrit.com/yaes-server/entities"
)

// UserRepository describes the behavior of a user repository
type UserRepository interface {
	RetrieveOthers(u entities.User) (users []*entities.User, err error)
	FindBy(username string) (u *entities.User, err error)
	FindByID(ID interface{}) (u *entities.User, err error)
	Save(u entities.User) (user *entities.User, err error)
}
