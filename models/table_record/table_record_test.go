package record

import (
	"testing"

	"github.com/subosito/gotenv"
)

// TestStruct - struct di test che implementa TableRecordInterface
type TestStruct struct {
	tr       *TableRecord
	Name     string `json:"name" db:"name"`
	Lastname string `json:"lastname" db:"lastname"`
	Gender   string `json:"gender" db:"gender"`
}

// GetTableRecord - Restituisce l'istanza di TableRecord
func (t TestStruct) GetTableRecord() *TableRecord {
	return t.tr
}

// GetPrimaryKeyName - Restituisce il nome della chiave primaria
func (t TestStruct) GetPrimaryKeyName() string {
	return "record_id"
}

// GetTableName - Restituisce il nome della tabella
func (t TestStruct) GetTableName() string {
	return "users"
}

// New - Si occupa di istanziare una nuova struct andando ad istaziare table record e settanto il campo isNew a true
func (t *TestStruct) New() {

	t.tr = new(TableRecord)
	t.tr.SetIsNew(true)
}

func TestTableRecord(t *testing.T) {
	gotenv.Load("./../../.env")

	ts := &TestStruct{}
	ts.New()

	ts.Name = "Mario"
	ts.Lastname = "Rossi"
	ts.Gender = "M"

	err := Save(ts)
	if err != nil {
		t.Error(err.Error())
	}

	if ts.tr.RecordID == 0 {
		t.Error("Chiave non salvata")
	}

	tempName := ts.Name
	tempID := ts.tr.RecordID

	err = LoadByID(ts, ts.GetTableRecord().RecordID)
	if err != nil {
		t.Error(err.Error())
	}

	if tempName != ts.Name {
		t.Error("Campi non uguali")
	}

	ts.Name = "Marco"

	err = Save(ts)
	if err != nil {
		t.Error(err.Error())
	}

	if tempID != ts.tr.RecordID {
		t.Error("Chiave primaria è cambiata durante l'update")
	}
}
