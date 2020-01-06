package record

import (
	"testing"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/helpers/copy"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/db"
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
	ts.tr = NewTableRecord(true, false)
	return ts
}

// loadAllTestTableRecordStruct - carica tutte le istanze della classe
func loadAllTestTableRecordStruct() ([]*TestStruct, error) {

	db := db.GetConnection()

	ts := NewTestStruct()

	query := "SELECT " + AllField(ts) + " FROM " + ts.GetTableName()

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	var result []*TestStruct

	for rows.Next() {

		ts := NewTestStruct()

		if err := LoadFromRow(rows, ts); err != nil {
			return nil, err
		}

		result = append(result, ts)
	}

	return result, nil
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

// TestStructReadOnly - Struct di test readonly che implementa TableRecordInterface
type TestStructReadOnly struct {
	tr       *TableRecord
	Name     *string `json:"name" db:"name"`
	Lastname *string `json:"lastname" db:"lastname"`
	Gender   *string `json:"gender" db:"gender"`
}

// NewTestStructReadOnly - Restituisce una nuova istanza di TestStructReadOnly
func NewTestStructReadOnly() *TestStructReadOnly {

	tsro := new(TestStructReadOnly)
	tsro.tr = NewTableRecord(true, true)

	return tsro
}

// GeetTableRecord - Restituisce l'istanza di TableRecord
func (t TestStructReadOnly) GeetTableRecord() *TableRecord {
	return t.tr
}

// GetPrimaryKeyName - Restituisce il nome della chiave primaria
func (t TestStructReadOnly) GetPrimaryKeyName() string {
	return "record_id"
}

// GetTableName - Restituisce il nome della tabella

func (t TestStructReadOnly) GetTableName() string {
	return "users"
}

// GetTableRecord - Restituisce l'istanza di TableRecord
func (t TestStructReadOnly) GetTableRecord() *TableRecord {
	return t.tr
}

func TestTableRecord(t *testing.T) {

	gotenv.Load("./../../../.env")

	ts := NewTestStruct()

	ts.Name = copy.String("Mario")
	ts.Lastname = copy.String("Rossi")
	ts.Gender = copy.String("M")

	err := Save(ts)
	if err != nil {
		t.Fatal(err.Error())
	}

	if ts.tr.recordID == 0 {
		t.Fatal("Chiave non salvata")
	}

	if ts.tr.IsNew() {
		t.Fatal("Il record non dovrebbe essere isNew = true")
	}

	tempName := *ts.Name
	tempID := ts.tr.recordID

	err = LoadByID(ts, ts.GetTableRecord().recordID)
	if err != nil {
		t.Fatal(err.Error())
	}

	if tempName != *ts.Name {
		t.Fatal("Campi non uguali")
	}

	ts.Name = copy.String("Marco")

	err = Save(ts)
	if err != nil {
		t.Fatal(err.Error())
	}

	if tempID != ts.tr.recordID {
		t.Fatal("Chiave primaria è cambiata durante l'update")
	}

	ts = NewTestStruct()

	ts.tr.WhereEqual("name", "Marco").OrderByDesc("record_id").WhereOperator("record_id", "<", 23)

	tsList, err := ExecQuery(ts, func() TableRecordInterface {
		return NewTestStruct()
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(tsList) == 0 {
		t.Error("La query sembra non aver restituito zero valori")
	}

	ts = NewTestStruct()

	ts.Name = copy.String("Mario")
	ts.Lastname = copy.String("Rossi")
	ts.Gender = copy.String("G")

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

	allResult, err := loadAllTestTableRecordStruct()
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(allResult) == 0 {
		t.Error("La lista restituita sembra vuota")
	}

	testAll, err := All(func() TableRecordInterface {
		return NewTestStruct()
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(testAll) == 0 {
		t.Error("La lista restituita sembra vuota")
	}

	tsr := NewTestStructReadOnly()

	tsr.Name = copy.String("foffo")
	tsr.Lastname = copy.String("bomba")
	tsr.Gender = copy.String("M")

	err = Save(tsr)
	if err == nil {
		t.Fatal("Errore: il model dovrebbe essere read-only")
	}

	err = LoadByID(tsr, 23)
	if err != nil {
		t.Fatal(err)
	}

	if tsr.tr.GetID() < 0 {
		t.Fatal("Errore: record id non valido")
	}

}
