package token

import (
	"auth-server/models"
	"os"

	"github.com/joho/godotenv"
)

type Token struct {
	secretkey []byte
}

func newtoken_user(u models.User) *Token {
	token := new(Token)
	token.initSecretKey()
	return token
}

func (t *Token) initSecretKey() {
	godotenv.Load(".env")
	t.secretkey = []byte(os.Getenv("LLAVE_SECRETA"))
}
