package testing

import (
	"auth-server/testing/functions_test"
	"testing"
)

func TestValidation_token(t *testing.T) {
	//token de prueba
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTc3NzU0NDMsInN1YiI6IjNjZDg4MGYxLTZkNmQtMTFlZS05Njk0LWQ4Y2I4YWNhYTk1MSJ9.8_2gFqrwSdDQawEM-rdaLGPR38d4oKrS0YntJhIOW8Q"
	functions_test.Validation_token(t, token)
}
