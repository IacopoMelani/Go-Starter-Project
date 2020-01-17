package db

import (
	"testing"

	"github.com/subosito/gotenv"
)

// TestGetConnection - Esegue il test della funziona GeTConnection()
func TestGetConnection(t *testing.T) {

	loadEnv()

	db := GetConnection()

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

	exists := TableExists("migrations")
	if !exists {
		t.Fatal("Attenzione tabella migrations non presente")
	}

	exists = TableExists("migrationss")
	if exists {
		t.Fatal("Errore, la tabella non dovrebbe esistere")
	}

}

func loadEnv() {
	gotenv.Load("../../.env")
}
