package Rutes

import (
	"auth-server/database/deployment"
	"fmt"
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
			c.JSON(400, "Firma no valida")
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		godotenv.Load(".env")
		key_hex := []byte(os.Getenv("LLAVE_SECRETA"))
		return key_hex, nil

	})

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	} else {
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			query := "SELECT COUNT(*) AS contador FROM user WHERE id = ?"
			rows, rows_error := connect.Query(query, claims["sub"])
			if rows_error != nil {
				c.JSON(400, gin.H{"error": rows_error.Error()})
				return
			}
			var (
				contador int
			)
			defer rows.Close()
			for rows.Next() {
				err_search := rows.Scan(&contador)
				if err_search != nil {
					c.JSON(400, gin.H{"error": err_search.Error()})
					return
				} else {
					if contador == 1 {
						c.JSON(200, claims["sub"])
						return
					} else {
						c.JSON(400, gin.H{"error": rows_error.Error()})
						return
					}
				}
			}
		} else {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
	}
}
