package modelpackage

import (
	"auth-server/Rutes"

	"github.com/gin-gonic/gin"
)

type Funcions struct {
	Router *gin.Engine
}

func NewRutes() (*Funcions, error) {
	r := gin.Default()

	r.GET("/", Rutes.Prueba)
	r.POST("/user/register_user", Rutes.Registrar_Usuario)
	r.POST("/user/log_user", Rutes.Validar_usuario)

	return &Funcions{Router: r}, nil
}