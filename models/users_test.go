package models

import (
	"testing"
)

const currentUserTableLen = 3

func TestLoadAll(t *testing.T) {

	users := LoadAllUser()
	if len(users) <= 0 {
		t.Error("Impossibile completare il test nessun dato presente nella tabella")
	}

	if len(users) != currentUserTableLen {
		t.Error("Numero di risultati errati")
	}

	users.PrintAll()

	user := users[0]

	if user.Name != "Iacopo" || user.Lastname != "Melani" || user.Gender != "M" {
		t.Error("Errore risultati primo risultato errato!")
	}

}
