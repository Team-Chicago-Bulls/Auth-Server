package main

import (
	"net/http"
	"vf-server/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// Crear un enrutador Gin
	r := gin.Default()

	// Definir una ruta para la solicitud GET a la raíz del sitio
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "¡Hola desde Gintonic!",
		})
	})

	// Iniciar el servidor en el puerto 8080
	link := "root:@tcp(127.0.0.1:3306)/token"
	db := database.NewDB(link)
	pool, fix := db.Connect()
	if fix != nil {
		println(fix)
	}
	// Consulta SQL
	query := "SELECT * from  users"
	rows, err := pool.Query(query)
	if err != nil {
		println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var correo string
		var contrasena string
		err := rows.Scan(&correo, &contrasena)
		if err != nil {
			println(err)
		}
		println("correo: " + correo)
	}

	r.Run(":8080")
}
