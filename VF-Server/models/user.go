package models

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type User struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func User_build(email string) *User {
	user := &User{
		Email: email,
	}
	user.ExpiresAt = time.Now().Add(time.Hour * 48).Unix()
	return user

}
