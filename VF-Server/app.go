package main

import (
	"os"
	"vf-server/rutes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	godotenv.Load(".env")
	enn, err := rutes.NewRutes()
	if err != nil {
		println("No se logro la conexión con la BD")
		return
	}
	// Ruta de ejemplo con comentario Swagger
	// @Summary Obtiene la lista de elementos
	// @Description Obtiene la lista de elementos
	// @Tags elementos
	// @Produce json
	// @Success 200 {array} string "Lista de elementos"
	// @Router /elementos [get]
	r.GET("/", enn.Prueba)
	r.POST("/user/enviar_data", enn.Enviar_data)

	r.Run(":" + os.Getenv("PORT"))
}
