package db

import (
	"fmt"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/IacopoMelani/Go-Starter-Project/config"
	"github.com/subosito/gotenv"
)

// createTableTest - Query to create table test
func createTableTest() error {

	conn := GetConnection()

	var query string

	switch conn.DriverName() {

	case DriverMySQL:
		query = `CREATE TABLE IF NOT EXISTS testTable (
			record_id INT AUTO_INCREMENT,
			PRIMARY KEY (record_id)
			)`

	case DriverSQLServer:
		query = `IF NOT EXISTS (SELECT * FROM sysobjects WHERE name='testTable' and xtype='U')
		CREATE TABLE testTable (
			record_id BIGINT IDENTITY(1, 1) NOT NULL PRIMARY KEY
		)`
	}

	_, err := conn.Exec(query)

	return err
}

// dropTableTest - Query to destoy the table
func dropTableTest() error {

	conn := GetConnection()

	query := "DROP TABLE IF EXISTS testTable"

	_, err := conn.Exec(query)

	return err
}

func TestGetConnection(t *testing.T) {

	loadEnv()

	db := GetSQLXFromSQLConnector(GetConnection())

	err := db.Ping()

	if err != nil {
		t.Error(err.Error())
	}
}

func TestMultiConnection(t *testing.T) {

	loadEnv()

	config := config.GetInstance()

	firstKey := "first"
	secondKey := "second"
	errorKey := "error"

	InitConnectionWithKey(firstKey, config.SQLDriver, config.StringConnection)

	connFirst, err := GetConnectionWithKey(firstKey)
	if err != nil {
		t.Error(err)
	}

	_, err = GetConnectionWithKey(errorKey)
	if err == nil {
		t.Error(fmt.Sprintf("Chiave '%s' non valida, dovrebbe essere error", errorKey))
	}

	InitConnectionWithKey(secondKey, config.SQLDriver, config.StringConnection)

	connSecond, err := GetConnectionWithKey(secondKey)
	if err != nil {
		t.Error(err)
	}

	if connFirst == connSecond {
		t.Error("Errore, le due connessione dovrebbero essere diverse")
	}
}

func TestQuery(t *testing.T) {

	loadEnv()

	query1 := "SELECT * FROM users"

	_, err := Query(query1)
	if err != nil {
		t.Fatal("Errore esecuzione query")
	}

	query2 := "SELECT * FROM users WHERE record_id = ?"
	_, err = Query(query2, 1)
	if err != nil {
		t.Fatal("Errore esecuzione query parametrizzata")
	}

	query3 := "SELECT * FORM users"
	_, err = Query(query3)
	if err == nil {
		t.Fatal("Errore, query dovrebbe essere sbagliata")
	}
}

func TestQueryOrPanic(t *testing.T) {

	loadEnv()

	query1 := "SELECT * FROM users"
	QueryOrPanic(query1)

	query2 := "SELECT * FROM users WHERE record_id = ?"
	QueryOrPanic(query2, 1)

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("Errore recover non riuscito")
		}
	}()

	query3 := "SELECT * FORM users"
	QueryOrPanic(query3)
}

func TestTableExists(t *testing.T) {

	loadEnv()

	if err := createTableTest(); err != nil {
		t.Fatal("Attenzione, impossibile creare la tabella di test")
	}

	exists := TableExists("testTable")
	if !exists {
		t.Fatal("Attenzione tabella di test non presente")
	}

	if err := dropTableTest(); err != nil {
		t.Fatal("Attenzione, impossibile cancellare la tabella di test")
	}

	exists = TableExists("testTable")
	if exists {
		t.Fatal("Errore, la tabella non dovrebbe esistere")
	}

}

func TestConnectionConcurrency(t *testing.T) {

	if err := gotenv.Load("../../../.env"); err != nil {
		panic(err)
	}

	wg := sync.WaitGroup{}

	wg.Add(4)

	go func() {
		defer wg.Done()
		GetConnection()
	}()
	go func() {
		defer wg.Done()
		GetConnection()
	}()
	go func() {
		defer wg.Done()
		GetConnection()
	}()
	go func() {
		defer wg.Done()
		GetConnection()
	}()

	time.Sleep(500 * time.Millisecond)

	InitConnection(os.Getenv("SQL_DRIVER"), os.Getenv("STRING_CONNECTION"))

	GetConnection()

	wg.Wait()
}

func loadEnv() {
	if err := gotenv.Load("../../../.env"); err != nil {
		panic(err)
	}
	InitConnection(os.Getenv("SQL_DRIVER"), os.Getenv("STRING_CONNECTION"))
}
