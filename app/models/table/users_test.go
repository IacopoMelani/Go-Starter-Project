package table

import (
	"os"
	"testing"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db"
	"gopkg.in/guregu/null.v4"

	record "github.com/IacopoMelani/Go-Starter-Project/pkg/models/table_record"
	"github.com/subosito/gotenv"
)

func TestTableMirror(t *testing.T) {

	if err := gotenv.Load("./../../../.env"); err != nil {
		t.Fatal("Errore caricamento configurazione")
	}
	db.InitConnection(os.Getenv("SQL_DRIVER"), os.Getenv("STRING_CONNECTION"))

	u := NewUser(db.GetConnection())

	u.Name = null.StringFrom("Mario")
	u.Lastname = null.StringFrom("Rossi")
	u.Gender = null.StringFrom("M")

	err := record.Save(u)
	if err != nil {
		t.Error(err.Error())
	}

	if u.RecordID == 0 {
		t.Error("Chiave non salvata")
	}

	tempName := u.Name.ValueOrZero()
	tempID := u.RecordID

	err = record.LoadByID(u, u.RecordID)
	if err != nil {
		t.Error(err.Error())
	}

	if tempName != u.Name.ValueOrZero() {
		t.Error("Campi non uguali")
	}

	u.Name = null.StringFrom("Marco")

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
