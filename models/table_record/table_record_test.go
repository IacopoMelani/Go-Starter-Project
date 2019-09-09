package record

import (
	"testing"

	"github.com/subosito/gotenv"
)

// TestStruct - struct di test che implementa TableRecordInterface
type TestStruct struct {
	tr       *TableRecord
	Name     *string `json:"name" db:"name"`
	Lastname *string `json:"lastname" db:"lastname"`
	Gender   *string `json:"gender" db:"gender"`
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

func (t *TestStruct) setGender(value string) *TestStruct {
	t.Gender = &value
	return t
}
func (t *TestStruct) setLastname(value string) *TestStruct {
	t.Lastname = &value
	return t
}
func (t *TestStruct) setName(value string) *TestStruct {
	t.Name = &value
	return t
}

func TestTableRecord(t *testing.T) {

	gotenv.Load("./../../.env")

	ts := NewTestStruct()

	ts.setName("Mario")
	ts.setLastname("Rossi")
	ts.setGender("M")

	err := Save(ts)
	if err != nil {
		t.Fatal(err.Error())
	}

	if ts.tr.RecordID == 0 {
		t.Fatal("Chiave non salvata")
	}

	if ts.tr.IsNew() {
		t.Fatal("Il record non dovrebbe essere isNew = true")
	}

	tempName := *ts.Name
	tempID := ts.tr.RecordID

	err = LoadByID(ts, ts.GetTableRecord().RecordID)
	if err != nil {
		t.Fatal(err.Error())
	}

	if tempName != *ts.Name {
		t.Fatal("Campi non uguali")
	}

	ts.setName("Marco")

	err = Save(ts)
	if err != nil {
		t.Fatal(err.Error())
	}

	if tempID != ts.tr.RecordID {
		t.Fatal("Chiave primaria è cambiata durante l'update")
	}

	ts = NewTestStruct()

	ts.tr.WhereEqual("name", "Marco").OrderByDesc("record_id").WhereOperator("record_id", "<", 23)

	tsList, err := ExecQuery(ts)
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(tsList) == 0 {
		t.Error("La query sembra non aver restituito zero valori")
	}

	ts = NewTestStruct()

	ts.setName("Mario")
	ts.setLastname("Rossi")
	ts.setGender("M")

	err = Save(ts)
	if err != nil {
		t.Fatal(err.Error())
	}

	rows, err := Delete(ts)
	if err != nil {
		t.Fatal(err.Error())
	}

	if rows <= 0 {
		t.Fatal("Errore: nessuna cancellazione effettuata")
	}
}
