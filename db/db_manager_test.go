package db

import (
	"testing"

	"github.com/subosito/gotenv"
)

// TestGetConnection - Esegue il test della funziona GeTConnection()
func TestGetConnection(t *testing.T) {

	gotenv.Load("../.env")

	db := GetConnection()

	err := db.Ping()

	if err != nil {
		t.Error(err.Error())
	}

}

func TestTableExists(t *testing.T) {

	gotenv.Load("../.env")

	exists, err := TableExists("users")
	if err != nil {
		t.Fatal("Errore durante la ricerca della tabella")
	}

	if !exists {
		t.Fatal("Attenzione tabella users non presente")
	}
}
