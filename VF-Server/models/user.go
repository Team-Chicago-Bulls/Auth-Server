package models

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type User struct {
	id string `json:"id"`
	jwt.StandardClaims
}

func User_build(id_param string) *User {
	user := &User{}
	user.StandardClaims = jwt.StandardClaims{
		Subject:   id_param,
		ExpiresAt: time.Now().AddDate(9999, 0, 0).Unix(),
	}
	return user
}
