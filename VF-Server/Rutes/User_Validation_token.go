package Rutes

import (
	"auth-server/database/deployment"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func Validar_usuario_token(c *gin.Context) {
	db := deployment.NewThing()
	connect, err := db.Connect()

	tokenString := string(c.Param("token"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Método de firma no válido: %v", token.Header["alg"])
		}
		godotenv.Load(".env")
		key_hex := []byte(os.Getenv("LLAVE_SECRETA"))
		fmt.Println(key_hex)
		return key_hex, nil

	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	} else {
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			query := "SELECT COUNT(*) AS contador FROM user WHERE EMAIL = ?"
			rows, err := connect.Query(query, claims["email"])
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
				} else {
					if contador == 1 {
						c.JSON(202, "yes")
					} else {
						c.JSON(200, "no")
					}
				}
			}
		} else {
			fmt.Println(err)
		}
	}

}
