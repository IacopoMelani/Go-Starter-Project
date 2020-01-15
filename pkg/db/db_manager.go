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

// Query - Esegue fisicamente la query e restituisce l'istanza di *Rows
func Query(query string, args ...interface{}) (*sql.Rows, error) {

	db := GetConnection()
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

// QueryOrPanic - Esegue fisicamente la query e restituisce l'istanza di *Rows, panic in caso di errore
func QueryOrPanic(query string, args ...interface{}) *sql.Rows {

	rows, err := Query(query, args...)
	if err != nil {
		panic(err)
	}

	return rows
}

// TableExists - Restituisce true se la tabella esiste altrimenti false
//TODO: Non serve ritornare error
func TableExists(tableName string) (bool, error) {

	query := "SELECT * FROM " + tableName + " LIMIT 1"

	db := GetConnection()

	rows, err := db.Query(query)
	if err != nil {
		return false, nil
	}
	defer rows.Close()

	return true, nil
}
