package querymanager

import (
	"Go-Starter-Project/models"
	"testing"

	"github.com/subosito/gotenv"
)

// TestSave - Si occupa di testare il salvataggio(insert/update) di una struct che implementa l'interfaccia Salvable
func TestSave(t *testing.T) {

	gotenv.Load("../../.env")

	u := models.User{}
	u.Name = "Filippo"
	u.Lastname = "Neri"
	u.Gender = "M"

	params := []interface{}{u.Name, u.Lastname, u.Gender}

	err := Save(&u, params)
	if err != nil {
		t.Error(err.Error())
	}

	if u.RecordID <= 0 {
		t.Error("Index errato, è minore di 0")
	}

	u.Name = "Mario"

	params = []interface{}{u.Name, u.Lastname, u.Gender, u.RecordID}

	lastID := u.RecordID

	err = Save(&u, params)

	if u.RecordID <= 0 || lastID != u.RecordID {
		t.Error("Index errato, è minore di 0 oppure non corrisponde con il precedente")
	}
}

func TestSelect(t *testing.T) {

	gotenv.Load("../../.env")

	TestSave(t)

	u := models.User{}

	err := Select(&u, 1, 1)
	if err != nil {
		t.Error("Errore nella Select 1, errore restituito:", err.Error())
	}

	if u.RecordID != 1 {
		t.Error("Errore 1: record non valido")
	}

}
