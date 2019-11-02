package table

import (
	"testing"

	record "github.com/IacopoMelani/Go-Starter-Project/models/table_record"

	"github.com/subosito/gotenv"
)

func TestTableMirror(t *testing.T) {

	gotenv.Load("./../../../.env")

	u := NewUser()

	u.SetName("Mario").SetLastname("Rossi").SetGender("M")

	err := record.Save(u)
	if err != nil {
		t.Error(err.Error())
	}

	if u.tr.GetID() == 0 {
		t.Error("Chiave non salvata")
	}

	tempName := *u.Name
	tempID := u.tr.GetID()

	err = record.LoadByID(u, u.GetTableRecord().GetID())
	if err != nil {
		t.Error(err.Error())
	}

	if tempName != *u.Name {
		t.Error("Campi non uguali")
	}

	u.SetName("Marco")

	err = record.Save(u)
	if err != nil {
		t.Error(err.Error())
	}

	if tempID != u.tr.GetID() {
		t.Error("Chiave primaria è cambiata durante l'update")
	}

	usersList, err := LoadAllUsers()
	if err != nil {
		t.Error(err.Error())
	}

	if len(usersList) == 0 {
		t.Error("Lunghezza inferiore al minimo nel contesto del test")
	}
}
