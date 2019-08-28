package table

import (
	record "github.com/IacopoMelani/Go-Starter-Project/models/table_record"
	"testing"

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

	if u.tr.RecordID == 0 {
		t.Error("Chiave non salvata")
	}

	tempName := *u.Name
	tempID := u.tr.RecordID

	err = record.LoadByID(u, u.GetTableRecord().RecordID)
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

	if tempID != u.tr.RecordID {
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
