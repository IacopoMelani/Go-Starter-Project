package db

import (
	"database/sql"

	"github.com/jmoiron/sqlx"

	// Carica il driver sql per la connessione al db
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"

	"log"
	"sync"
)

// SQLConnector - Interfaccia per gestire operazioni sotto transaction
type SQLConnector interface {
	DriverName() string
	Exec(string, ...interface{}) (sql.Result, error)
	Prepare(string) (*sql.Stmt, error)
	Preparex(string) (*sqlx.Stmt, error)
	Query(string, ...interface{}) (*sql.Rows, error)
	Queryx(string, ...interface{}) (*sqlx.Rows, error)
	QueryRow(string, ...interface{}) *sql.Row
}

// Defines all possible sql drivers
const (
	DriverSQLServer = "mssql"
	DriverMySQL     = "mysql"
)

var (
	db   *sqlx.DB
	ok   = make(chan bool, 1)
	once sync.Once
)

// DriverName - Returns driver name
func DriverName() string {
	return GetConnection().DriverName()
}

// GetConnection - restituisce un'istanza di connessione al database
func GetConnection() SQLConnector {
	<-ok
	return db
}

// InitConnection - Inizializza la connessione impostando il driver e la stringa di connessione
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

// Query - Esegue fisicamente la query e restituisce l'istanza di *Rows
func Query(query string, args ...interface{}) (*sqlx.Rows, error) {

	db := GetConnection()
	rows, err := db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}

	return rows, err
}

// QueryOrPanic - Esegue fisicamente la query e restituisce l'istanza di *Rows, panic in caso di errore
func QueryOrPanic(query string, args ...interface{}) *sqlx.Rows {

	rows, err := Query(query, args...)
	if err != nil {
		panic(err)
	}

	return rows
}

// TableExists - Restituisce true se la tabella esiste altrimenti false
func TableExists(tableName string) bool {

	var query string

	switch DriverName() {
	case DriverMySQL:
		query = "SELECT * FROM " + tableName + " LIMIT 1"
	case DriverSQLServer:
		query = "SELECT TOP 1 * FROM " + tableName
	}

	db := GetConnection()

	rows, err := db.Queryx(query)
	if err != nil {
		return false
	}
	defer rows.Close()

	return true
}
