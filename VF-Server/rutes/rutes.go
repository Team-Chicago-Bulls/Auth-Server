package rutes

import (
	"vf-server/database/deployment"
	"vf-server/database/model_base"

	"github.com/gin-gonic/gin"
)

type Rutes struct {
	DB *model_base.DB
}

func NewRutes() (*Rutes, error) {
	db := deployment.NewThing()
	return &Rutes{DB: db}, nil
}

func (r *Rutes) Prueba(c *gin.Context) {
	c.JSON(200, "Especatacular")
}

// Registrar usuario
func (r *Rutes) Enviar_data(c *gin.Context) {
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

	db, err := r.DB.Connect()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	c.JSON(200, "nombre:"+campo1+" contrasena: "+campo2)
}
