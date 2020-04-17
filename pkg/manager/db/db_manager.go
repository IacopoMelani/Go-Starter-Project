package db

import (
	"database/sql"

	"github.com/jmoiron/sqlx"

	// Carica il driver mysql per la connessione al db
	_ "github.com/go-sql-driver/mysql"

	"log"
	"sync"
)

// SQLConnector - Generalizes a sql connection
type SQLConnector interface {
	Exec(string, ...interface{}) (sql.Result, error)
	Prepare(string) (*sql.Stmt, error)
	Preparex(string) (*sqlx.Stmt, error)
	Query(string, ...interface{}) (*sql.Rows, error)
	Queryx(string, ...interface{}) (*sqlx.Rows, error)
	QueryRow(string, ...interface{}) *sql.Row
}

var (
	db   *sqlx.DB
	ok   = make(chan bool, 1)
	once sync.Once
)

// GetConnection - Returns an instance of the db connection
func GetConnection() SQLConnector {
	<-ok
	return db
}

// InitConnection - Initialize the connection with driver and connection string
func InitConnection(drvName string, connection string) {

	once.Do(func() {

		conn, err := sqlx.Open(drvName, connection)
		if err != nil {
			log.Panic(err.Error())
		}

		if err := conn.Ping(); err != nil {
			log.Panic(err.Error())
		}
		db = conn

		ok <- true
		close(ok)
	})
}

// Query - Executes the query and return a *Rows instance
func Query(query string, args ...interface{}) (*sqlx.Rows, error) {

	db := GetConnection()
	rows, err := db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}

	return rows, err
}

// QueryOrPanic - Executes the query and return a *Rows instance, panics if error occurs
func QueryOrPanic(query string, args ...interface{}) *sqlx.Rows {

	rows, err := Query(query, args...)
	if err != nil {
		panic(err)
	}

	return rows
}

// TableExists - Returns true if table exists otherwise false
func TableExists(tableName string) bool {

	query := "SELECT * FROM " + tableName + " LIMIT 1"

	db := GetConnection()

	rows, err := db.Queryx(query)
	if err != nil {
		return false
	}
	defer rows.Close()

	return true
}
