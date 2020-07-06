package repository

import (
	"algogrit.com/yaes-server/entities"
	"github.com/jinzhu/gorm"
)

type userRepository struct {
	*gorm.DB
}

func (ur *userRepository) RetrieveOthers(u entities.User) ([]*entities.User, error) {
	var users []*entities.User

	err := ur.Where("id != ?", u.ID).Find(&users).Error

	return users, err
}

func (ur *userRepository) FindBy(username string) (*entities.User, error) {
	user := new(entities.User)
	err := ur.Where("username = ?", username).First(user).Error

	return user, err
}

func (ur *userRepository) FindByID(id interface{}) (*entities.User, error) {
	user := new(entities.User)
	err := ur.Where("id = ?", id).First(user).Error

	return user, err
}

func (ur *userRepository) Save(u entities.User) (*entities.User, error) {
	err := ur.Create(&u).Error

	return &u, err
}

// New creates an instance of UserRepository
func New(db *gorm.DB) UserRepository {
	return &userRepository{db}
}
