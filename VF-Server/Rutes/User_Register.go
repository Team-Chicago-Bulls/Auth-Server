package Rutes

import (
	"auth-server/database/deployment"
	"auth-server/models"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func Registrar_Usuario(c *gin.Context) {
	db := deployment.NewThing()
	connect, err_database := db.Connect()
	if err_database != nil {
		c.JSON(400, gin.H{"error": err_database.Error()})
		return
	} else {
		var datos map[string]interface{}
		if err_capture := c.ShouldBindJSON(&datos); err_capture != nil {
			c.JSON(400, gin.H{"error": err_capture.Error()})
			return
		}
		
		email_data, ok1 := datos["correo"].(string)
		password_data, ok2 := datos["contrasena"].(string)
		id_data, ok3 := datos["id"].(string)
		
		if !ok1 || !ok2 || !ok3  {
			c.JSON(400, gin.H{"error": "Los campos no son del tipo esperado"})
			return
		}

		query := "SELECT COUNT(*) AS contador FROM user WHERE EMAIL = ?"
		rows, err_cont_select := connect.Query(query, email_data)
		if err_cont_select != nil {
			c.JSON(400, gin.H{"error": err_cont_select.Error()})
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
			}
			if contador < 1 {
				passbyte := []byte(password_data)
				hash, err_hasheo := bcrypt.GenerateFromPassword(passbyte, bcrypt.DefaultCost)
				if err_hasheo != nil {
					c.JSON(400, gin.H{"error": err_hasheo.Error()})
					return
				} else {
					query_insert := "INSERT INTO user(id,email,Password_user) VALUES (?,?,?)"
					insert, error_insert_query := connect.Query(query_insert, id_data ,email_data, hash)
					if error_insert_query != nil {
						c.JSON(400, gin.H{"error": error_insert_query.Error()})
						return
					} else {
						query_select_id := "SELECT id from user where email = ?"
						select_id , error_select_id := connect.Query(query_select_id,email_data)
						if error_select_id != nil{
							c.JSON(400, gin.H{"error": error_select_id.Error()})
							return
						}else{
							var id_post int;
							if select_id.Next() {
								if err_scan := select_id.Scan(&id_post); err_scan != nil {
									c.JSON(400, gin.H{"error": err_scan.Error()})
									return
								}else{
									godotenv.Load(".env")
									key_hex := []byte(os.Getenv("LLAVE_SECRETA"))
									token_model := models.User_build(strconv.Itoa(id_post))
									token := jwt.NewWithClaims(jwt.SigningMethodHS256, token_model)
									final_token, err_token := token.SignedString(key_hex)
									if err_token != nil {
										panic(err_token.Error())
									} else {
										c.JSON(200, final_token)
									}
								}
							} else {
								c.JSON(400, gin.H{"error": "Error interno (SCAN)"})
								return
							}
						}
						defer select_id.Close()
					}
					defer insert.Close()
					
				}
			} else {
				c.JSON(400,"El usuario ya esta registrado")
				return
			}
		}
	}
}
