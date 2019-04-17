package db

import (
	"database/sql"
	// Carica il driver mysql per la connessione al db
	_ "github.com/go-sql-driver/mysql"

	"log"
	"sync"
)

const stringConnection = "root:root@tcp(127.0.0.1:3306)/test"

var db *sql.DB

var once sync.Once

// GetConnection -
func GetConnection() *sql.DB {

	once.Do(func() {

		conn, err := sql.Open("mysql", stringConnection)
		if err != nil {
			log.Panic(err.Error())
		}

		if err := conn.Ping(); err != nil {
			log.Panic(err.Error())
		}

		db = conn
	})

	return db
}
