package service

// CreateUserRequest contains the user fields
type CreateUserRequest struct {
	Username     string
	FirstName    string
	LastName     string
	MobileNumber string
	Password     string
}

// LoginRequest contains the credentials
type LoginRequest struct {
	Username string
	Password string
}

// LoginResponse contains the valid auth token
type LoginResponse map[string]string
