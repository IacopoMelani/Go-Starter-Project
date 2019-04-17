package db

import (
	"testing"
)

// TestGetConnection - Esegue il test della funziona GeTConnection()
func TestGetConnection(t *testing.T) {

	db := GetConnection()

	err := db.Ping()

	if err != nil {
		t.Error(err.Error())
	}

}
