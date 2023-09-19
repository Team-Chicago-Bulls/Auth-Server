package modelpackage

import (
	Rutes "vf-server/rutes"

	"github.com/gin-gonic/gin"
)

type Funcions struct {
	Router *gin.Engine
}

func NewRutes() (*Funcions, error) {
	r := gin.Default()

	r.GET("/", Rutes.Prueba)
	r.POST("/user/enviar_data", Rutes.Registrar_Usuario)

	return &Funcions{Router: r}, nil
}
