package service

type createUserForm struct {
	Username     string
	FirstName    string
	LastName     string
	MobileNumber string
	Password     string
}

type loginForm struct {
	Username string
	Password string
}
