package db

import (
	"database/sql"

	"github.com/IacopoMelani/Go-Starter-Project/config"

	// Carica il driver mysql per la connessione al db
	_ "github.com/go-sql-driver/mysql"

	"log"
	"sync"
)

var (
	db   *sql.DB
	once sync.Once
)

// GetConnection - restituisce un'istanza di connessione al database
func GetConnection() *sql.DB {

	once.Do(func() {

		config := config.GetInstance()
		conn, err := sql.Open("mysql", config.StringConnection)
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

// TableExists - Restituisce true se la tabella esiste altrimenti false
func TableExists(tableName string) (bool, error) {

	query := "SELECT * FROM " + tableName + " LIMIT 1"

	db := GetConnection()

	rows, err := db.Query(query)
	if err != nil {
		return false, nil
	}
	defer rows.Close()

	if rows.Next() {

		return true, nil
	}
	return false, nil
}
