package rutes

import (
	"vf-server/database"

	"github.com/gin-gonic/gin"
)

type Rutes struct {
	DB *database.DB
}

func NewRutes() (*Rutes, error) {
	cre := database.NewThing()
	print(cre.Ip)
	db := database.NewDB("sdads")
	return &Rutes{DB: db}, nil
}

func (r *Rutes) Prueba(c *gin.Context) {
	c.JSON(200, "Especatacular")
}

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
