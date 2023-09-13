package main

import "vf-server/database"

func main() {
	cre := database.NewThing()
	//"usuario:contraseña@tcp(hostname:puerto)/basededatos"
	print(cre.Ip)
}
