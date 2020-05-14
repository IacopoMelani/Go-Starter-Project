package table

import (
	"os"
	"testing"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/helpers/copy"
	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db"

	record "github.com/IacopoMelani/Go-Starter-Project/pkg/models/table_record"
	"github.com/subosito/gotenv"
)

func TestTableMirror(t *testing.T) {

	if err := gotenv.Load("./../../.env"); err != nil {
		t.Fatal("Errore caricamento configurazione")
	}
	db.InitConnection(os.Getenv("SQL_DRIVER"), os.Getenv("STRING_CONNECTION"))

	u := NewUser(db.GetConnection())

	u.Name = copy.String("Mario")
	u.Lastname = copy.String("Rossi")
	u.Gender = copy.String("M")

	err := record.Save(u)
	if err != nil {
		t.Error(err.Error())
	}

	if u.RecordID == 0 {
		t.Error("Chiave non salvata")
	}

	tempName := *u.Name
	tempID := u.RecordID

	err = record.LoadByID(u, u.RecordID)
	if err != nil {
		t.Error(err.Error())
	}

	if tempName != *u.Name {
		t.Error("Campi non uguali")
	}

	u.Name = copy.String("Marco")

	err = record.Save(u)
	if err != nil {
		t.Error(err.Error())
	}

	if tempID != u.RecordID {
		t.Error("Chiave primaria è cambiata durante l'update")
	}

	usersList, err := LoadAllUsers(db.GetConnection())
	if err != nil {
		t.Error(err.Error())
	}

	if len(usersList) == 0 {
		t.Error("Lunghezza inferiore al minimo nel contesto del test")
	}
}
