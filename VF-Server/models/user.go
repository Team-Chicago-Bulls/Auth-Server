package models

import "github.com/golang-jwt/jwt"

type User struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
	jwt.StandardClaims
}

func user_build() *User {
	user := new(User)
	return user
}
