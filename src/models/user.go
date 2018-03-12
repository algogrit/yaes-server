package model

import (
	log "github.com/sirupsen/logrus"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

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

func HashAndSalt(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Error(err)
	}

	return string(hash)
}

func ComparePasswords(hashedPwd string, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	if err != nil {
		log.Error(err)
		return false
	}

	return true
}

func CreateJWTToken(user User, jwtSigningKey []byte) map[string]string {
	token := jwt.New(jwt.SigningMethodHS256)

	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	claims["user"] = user
	claims["userID"] = user.ID

	/* Sign the token with our secret */
	tokenString, _ := token.SignedString(jwtSigningKey)

	tokenMap := map[string]string{"token": tokenString}

	return tokenMap
}

func FindUserFromToken(jwtToken *jwt.Token, db *gorm.DB) (User, error) {
	userID := jwtToken.Claims.(jwt.MapClaims)["userID"]

	var user User

	err := db.Where("id = ?", userID).First(&user).Error

	return user, err
}
