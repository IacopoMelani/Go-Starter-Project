package db

import (
	"Go-Starter-Project/models"
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

// TestSave - Si occupa di testare il salvataggio(insert/update) di una struct che implementa l'interfaccia Salvable
func TestSave(t *testing.T) {

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

	TestSave(t)

	u := models.User{}

	err := Select(&u, 1, []interface{}{7})
	if err != nil {
		t.Error("Errore nella Select, errore restituito:", err.Error())
	}
}
