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

// NewTestStruct - Restitusice una nuova istaza di TestStruct
func NewTestStruct() *TestStruct {

	ts := new(TestStruct)
	ts.tr = new(TableRecord)
	ts.tr.SetIsNew(true)

	return ts
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
func (t TestStruct) New() TableRecordInterface {
	return NewTestStruct()
}

func TestTableRecord(t *testing.T) {
	gotenv.Load("./../../.env")

	ts := NewTestStruct()

	ts.Name = "Mario"
	ts.Lastname = "Rossi"
	ts.Gender = "M"

	err := Save(ts)
	if err != nil {
		t.Fatal(err.Error())
	}

	if ts.tr.RecordID == 0 {
		t.Fatal("Chiave non salvata")
	}

	tempName := ts.Name
	tempID := ts.tr.RecordID

	err = LoadByID(ts, ts.GetTableRecord().RecordID)
	if err != nil {
		t.Fatal(err.Error())
	}

	if tempName != ts.Name {
		t.Fatal("Campi non uguali")
	}

	ts.Name = "Marco"

	err = Save(ts)
	if err != nil {
		t.Fatal(err.Error())
	}

	if tempID != ts.tr.RecordID {
		t.Fatal("Chiave primaria è cambiata durante l'update")
	}

	ts = NewTestStruct()

	ts.tr.WhereEqual("name", "Marco").OrderByDesc("record_id")

	tsList, err := ExecQuery(ts)
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(tsList) == 0 {
		t.Error("La query sembra non aver restituito zero valori")
	}

}
