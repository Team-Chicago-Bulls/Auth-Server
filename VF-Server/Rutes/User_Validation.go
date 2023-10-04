package Rutes

import (
	"auth-server/database/deployment"
	"auth-server/models"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func Validar_usuario(c *gin.Context) {
	result_found := false
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

		if !ok1 && !ok2 {
			c.JSON(400, gin.H{"error": "Los campos no son del tipo esperado"})
			return
		}

		query := "SELECT * FROM user WHERE EMAIL = ?"
		rows, err := connect.Query(query, email_data)
		if err != nil {
			panic(err.Error())
		}
		var (
			email    string
			password string
		)
		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(&email, &password)
			if err != nil {
				panic(err.Error())
			} else {
				result_found = true
				erro_has := bcrypt.CompareHashAndPassword([]byte(password), []byte(password_data))
				if email_data == email && erro_has == nil {
					godotenv.Load(".env")
					key_hex := []byte(os.Getenv("LLAVE_SECRETA"))
					fmt.Println(key_hex)
					token_model := models.User_build(email_data)
					token := jwt.NewWithClaims(jwt.SigningMethodHS256, token_model)
					final_token, err_token := token.SignedString(key_hex)
					if err_token != nil {
						c.JSON(500, gin.H{"error": err_token.Error()})
						return
					}
					c.JSON(202, final_token)
				} else {
					c.JSON(408, "Las contraseñas no coinciden")
				}
			}
		}
		if !result_found {
			c.JSON(205, "No se encontro usuario")
		}
	}
}
