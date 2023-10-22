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

	r.GET("/user/prueba", Rutes.Prueba)
	r.GET("/user/documentacion", Rutes.Documentacion)
	r.GET("/user/identificacion/:correo", Rutes.Validar_correo_id)
	r.GET("/user/log_user_token/:token", Rutes.Validar_usuario_token)
	r.POST("/user/register_user", Rutes.Registrar_Usuario)
	r.POST("/user/log_user", Rutes.Validar_usuario)

	return &Funcions{Router: r}, nil
}
