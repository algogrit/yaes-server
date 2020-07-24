package service

import (
	"context"

	"algogrit.com/yaes-server/entities"
)

// UserService is used for creating a user service
type UserService interface {
	Index(ctx context.Context, currentUser entities.User) ([]*entities.User, error)
	Create(ctx context.Context, req CreateUserRequest) (*entities.User, error)
	Login(ctx context.Context, credentials LoginRequest) (LoginResponse, error)
}
