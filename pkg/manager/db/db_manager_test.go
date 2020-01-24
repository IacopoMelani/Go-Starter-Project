package db

import (
	"testing"

	"github.com/subosito/gotenv"
)

// createTableTest - Si occupa di create la tabella di test
func createTableTest() error {

	conn := GetConnection()

	query := `CREATE TABLE IF NOT EXISTS testTable (
		record_id INT AUTO_INCREMENT,
		PRIMARY KEY (record_id)
		)`

	_, err := conn.Exec(query)

	return err
}

// dropTableTest - Si occupa di distruggere la tabella di test
func dropTableTest() error {

	conn := GetConnection()

	query := "DROP TABLE IF EXISTS testTable"

	_, err := conn.Exec(query)

	return err
}

// TestGetConnection - Esegue il test della funziona GeTConnection()
func TestGetConnection(t *testing.T) {

	loadEnv()

	GetConnection()

	err := db.Ping()

	if err != nil {
		t.Error(err.Error())
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

func loadEnv() {
	if err := gotenv.Load("../../../.env"); err != nil {
		panic(err)
	}
}
