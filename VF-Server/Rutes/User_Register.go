package Rutes

import (
	"auth-server/database/deployment"
	"auth-server/models"
	"encoding/hex"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func Registrar_Usuario(c *gin.Context) {
	db := deployment.NewThing()
	connect, err := db.Connect()
	if err != nil {
		print(err)
	} else {
		var datos map[string]interface{}
		if err := c.ShouldBindJSON(&datos); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		email_data, ok1 := datos["correo"].(string)
		password_data, ok2 := datos["contrasena"].(string)

		if !ok1 || !ok2 {
			c.JSON(400, gin.H{"error": "Los campos no son del tipo esperado"})
			return
		}

		query := "SELECT COUNT(*) AS contador FROM USER WHERE EMAIL = ?"
		rows, err := connect.Query(query, email_data)
		if err != nil {
			panic(err.Error())
		}
		var (
			contador int
		)
		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(&contador)
			if err != nil {
				panic(err.Error())
			}
			if contador < 1 {
				passbyte := []byte(password_data)
				hash, err_hasheo := bcrypt.GenerateFromPassword(passbyte, bcrypt.DefaultCost)
				if err_hasheo != nil {
					panic(err_hasheo.Error())
				} else {
					query_insert := "INSERT INTO user(email,Password_user) VALUES (?,?)"
					insert, error2 := connect.Query(query_insert, email_data, hash)
					if error2 != nil {
						panic(error2.Error())
					} else {
						godotenv.Load(".env")
						key_hex := os.Getenv("LLAVE_SECRETA")
						key_bin, err_key := hex.DecodeString(key_hex)
						if err_key != nil {
							panic(err_key.Error())
						} else {
							token_model := models.User_build(email_data)
							token := jwt.NewWithClaims(jwt.SigningMethodHS256, token_model)
							final_token, err_token := token.SignedString(key_bin)
							if err_token != nil {
								panic(err_token.Error())
							} else {
								c.JSON(200, final_token)
							}
						}
					}
					defer insert.Close()
				}
			} else {
				c.JSON(404, "El usuario ya esta registrado")
			}
		}
	}
}
