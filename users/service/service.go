package service

import (
	"context"

	"algogrit.com/yaes-server/entities"
	httpError "algogrit.com/yaes-server/pkg/http_error"
	"algogrit.com/yaes-server/users/repository"
)

type userService struct {
	repository.UserRepository
	jwtSigningKey string
}

func (us *userService) Index(ctx context.Context, currentUser entities.User) ([]*entities.User, error) {
	return us.RetrieveOthers(currentUser)
}

func (us *userService) Create(ctx context.Context, req CreateUserRequest) (*entities.User, error) {
	user := entities.User{
		Username:     req.Username,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		MobileNumber: req.MobileNumber,
	}

	user.SetPassword(req.Password)

	createdUser, err := us.Save(user)

	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (us *userService) Login(ctx context.Context, credentials LoginRequest) (LoginResponse, error) {
	user, err := us.FindBy(credentials.Username)

	if err != nil || !user.MatchPassword(credentials.Password) {
		return nil, httpError.UnauthorizedErr().Wrap(err)
	}

	token, err := user.NewJWT(us.jwtSigningKey)

	if err != nil {
		return nil, err
	}

	return map[string]string{"token": token}, nil
}

// New creates a new instance of UserService
func New(repo repository.UserRepository, jwtSigningKey string) UserService {
	return &userService{repo, jwtSigningKey}
}
