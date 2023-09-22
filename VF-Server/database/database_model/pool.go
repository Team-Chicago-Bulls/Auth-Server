package database_model

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Estructura
type DB struct {
	link string
}

// Constructor
func NewDB(link string) *DB {
	return &DB{link: link}
}

// Método para realizar la conexión
func (db *DB) Connect() (*sql.DB, error) {
	database, err := sql.Open("mysql", db.link)
	if err != nil {
		return nil, err
	}
	err = database.Ping()
	if err != nil {
		return nil, err
	}
	return database, nil
}
