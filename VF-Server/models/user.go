package models

import "github.com/golang-jwt/jwt"

type User struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func User_build(email string) *User {
	temp := new(User)
	temp.Email = email
	claims := jwt.StandardClaims{
		ExpiresAt: 15000,
	}
	temp.StandardClaims = claims
	return temp

}
