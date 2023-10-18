package testing

import (
	"auth-server/testing/functions_test"
	"math/rand"
	"strings"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	caracteres := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	for i := 0; i < 1; i++ {
		var builder_email strings.Builder
		var builder_password strings.Builder
		for i := 0; i < 10; i++ {
			randomIndex := rand.Intn(len(caracteres) - 1)
			randomIndex_2 := rand.Intn(len(caracteres) - 1)
			builder_email.WriteByte(caracteres[randomIndex])
			builder_password.WriteByte(caracteres[randomIndex_2])
		}

		email := builder_email.String()
		password := builder_password.String()
		functions_test.RegisterUser(t, email, password)
	}
}
