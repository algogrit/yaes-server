package model

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username       string `gorm:"not null;unique"`
	HashedPassword string `json:"-" gorm:"not null"`
	FirstName      string
	LastName       string
	MobileNumber   string    `gorm:"not null;unique"`
	Expenses       []Expense `gorm:"ForeignKey:CreatedBy"`
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

func FindUserFromToken(jwtToken *jwt.Token, db *gorm.DB) User {
	userID := jwtToken.Claims.(jwt.MapClaims)["userID"]

	var user User
	db.Where("id = ?", userID).First(&user)

	return user
}
