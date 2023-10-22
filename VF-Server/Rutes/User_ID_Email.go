package Rutes

import (
	"auth-server/database/deployment"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Validar_correo_id(c *gin.Context) {
	user := false
	db := deployment.NewThing()
	connect, err := db.Connect()
	correo := c.Param("correo")
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	} else {
		query := "SELECT id as id_email FROM user WHERE email = ?"
		rows, rows_error := connect.Query(query, correo)
		if rows_error != nil {
			c.JSON(400, gin.H{"error": rows_error.Error()})
			return
		}
		var (
			id_email string
		)
		defer rows.Close()
		for rows.Next() {
			err_search := rows.Scan(&id_email)
			if err_search != nil {
				c.JSON(400, gin.H{"error": err_search.Error()})
				return
			} else {
				user = true
				c.JSON(200, gin.H{"id": id_email})
				return
			}
		}

		if user == false {
			c.JSON(400, gin.H{"error": "Ese correo no esta registrado"})
		}

	}
}
