package entities

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// User represents a single user in the system
type User struct {
	gorm.Model
	Username       string `gorm:"not null;unique"`
	HashedPassword string `json:"-" gorm:"not null"`
	FirstName      string
	LastName       string
	MobileNumber   string    `gorm:"not null;unique"`
	Expenses       []Expense `gorm:"ForeignKey:CreatedBy" json:"-"`
	Payables       []Payable `json:"-"`
}

// SetPassword sets the hashed password
func (u *User) SetPassword(pwd string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Error(err)
	}

	u.HashedPassword = string(hash)
}

// MatchPassword compares plain text password with the saved hashed password
func (u *User) MatchPassword(pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(pwd))

	if err != nil {
		log.Error(err)
		return false
	}

	return true
}

// NewJWT creates a jwt token with claims for a given user
func (u *User) NewJWT(jwtSigningKey string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	claims["user"] = u
	claims["userID"] = u.ID

	/* Sign the token with our secret */
	return token.SignedString([]byte(jwtSigningKey))
}
