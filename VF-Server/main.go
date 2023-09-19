package main

import (
	"os"
	modelpackage "vf-server/Rutes/ModelPackage"

	"github.com/joho/godotenv"
)

// @title Auth-Server API
// @version 0.1
// @description Descripción del funcionamiento de la API respecto al servidor de autentificación
// @host localhost:1022
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
