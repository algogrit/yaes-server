package model

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

var jwtSigningKey = []byte("483175006c1088c849502ef22406ac4e")

type User struct {
	gorm.Model
	Username       string `gorm:"not null;unique"`
	HashedPassword string `json:"-" gorm:"not null"`
	FirstName      string
	LastName       string
	MobileNumber   string `gorm:"not null;unique"`
}

func CreateJWTToken(user User) map[string]string {
	token := jwt.New(jwt.SigningMethodHS256)

	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	claims["user"] = user

	/* Sign the token with our secret */
	tokenString, _ := token.SignedString(jwtSigningKey)

	tokenMap := map[string]string{"token": tokenString}

	return tokenMap
}
