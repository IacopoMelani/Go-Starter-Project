package db

import (
	"database/sql"
	// Carica il driver mysql per la connessione al db
	_ "github.com/go-sql-driver/mysql"

	"log"
	"sync"
)

// Salvable - Interfaccia per permettere di generalizzare il salvataggio di un model sul database
type Salvable interface {
	GetSaveQuery() string
	SetRecordID(id int)
}

// Selecter - Interfaccia per permettere di generalizzare una select di un model
type Selecter interface {
	GetSelectQuery(s int) (string, []interface{})
}

const stringConnection = "root:Suite&Table@2017@tcp(10.10.10.9:3306)/test"

var db *sql.DB

var once sync.Once

// GetConnection - restituisce un'istanza di connessione al database
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

// Save - Metodo che si occcupa del salvataggio fisico sul database
func Save(s Salvable, params []interface{}) error {

	db := GetConnection()

	stmt, err := db.Prepare(s.GetSaveQuery())
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(params...)
	if err != nil {
		return err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	if lastID > 0 {
		s.SetRecordID(int(lastID))
	}
	return nil
}

// Select - Metodo che si occupa di eseguire una select sul database
func Select(s Selecter, query int, params []interface{}) error {

	db := GetConnection()

	querySQL, scanFields := s.GetSelectQuery(query)

	stmt, err := db.Prepare(querySQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Query(params...)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(scanFields...)
		if err != nil {
			return err
		}
	}

	return nil
}
