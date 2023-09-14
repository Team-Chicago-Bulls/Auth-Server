package deployment

import (
	"os"
	"vf-server/database/model_base"

	"github.com/joho/godotenv"
)

type Deploy struct {
	ip       string
	user     string
	password string
	nam      string
	port     string
}

func NewThing() *model_base.DB {
	godotenv.Load(".env")
	p := new(Deploy)
	p.ip = os.Getenv("DB_HOST")
	p.user = os.Getenv("DB_USER")
	p.nam = os.Getenv("DB_DATABASENAME")
	p.password = os.Getenv("DB_PASSWORD")
	p.port = os.Getenv("DB_PORT")
	//"usuario:contraseña@tcp(hostname:puerto)/basededatos"
	x := model_base.NewDB(p.user + ":" + p.password + "@tcp(" + p.ip + ":" + p.port + ")/" + p.nam)
	return x
}
