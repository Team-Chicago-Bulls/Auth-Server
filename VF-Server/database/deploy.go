package database

import (
	"os"

	"github.com/joho/godotenv"
)

type Database struct {
	Ip       string
	User     string
	Password string
	Nam      string
	Port     string
}

func NewThing() *Database {
	godotenv.Load(".env")
	p := new(Database)
	p.Ip = os.Getenv("DB_HOST")
	p.User = os.Getenv("DB_USER")
	p.Nam = os.Getenv("DB_DATABASENAME")
	p.Password = os.Getenv("DB_PASSWORD")
	p.Port = os.Getenv("DB_PORT")
	return p
}
