package main

import (
	modelpackage "auth-server/Rutes/ModelPackage"
	"os"

	"github.com/joho/godotenv"
)

// @title Auth-Server API
// @version 0.1
// @description Descripción del funcionamiento de la API respecto al servidor de autentificación
// @host localhost:8050
func main() {
	godotenv.Load(".env")
	enn, err := modelpackage.NewRutes()
	if err != nil {
		println(err)
		return
	}
	r := enn.Router
	r.Run(":" + os.Getenv("PORT"))
}
