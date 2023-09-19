package Rutes

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Registrar_Usuario(c *gin.Context) {
	var datos map[string]interface{}
	if err := c.ShouldBindJSON(&datos); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	campo1, ok1 := datos["correo"].(string)
	campo2, ok2 := datos["contrasena"].(string)

	if !ok1 || !ok2 {
		c.JSON(400, gin.H{"error": "Los campos no son del tipo esperado"})
		return
	}

	c.JSON(200, "nombre:"+campo1+" contrasena: "+campo2)
}
