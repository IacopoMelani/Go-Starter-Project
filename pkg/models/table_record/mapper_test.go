package record

import (
	"testing"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db"
	"github.com/subosito/gotenv"
)

// TestMapperTableRecord . Struct per test dei metodi di mapper
type TestMapperTableRecord struct {
	tr       *TableRecord
	RecordID int64   `json:"id" db:"record_id"`
	Name     *string `json:"name" db:"name"`
	Lastname *string `json:"lastname" db:"lastname"`
	Gender   *string `json:"gender" db:"gender"`
}

// NewTestMapperTableRecord - Restitusice una nuova istaza di TestMapperTableRecord
func NewTestMapperTableRecord(db db.SQLConnector) *TestMapperTableRecord {

	ts := new(TestMapperTableRecord)
	ts.tr = NewTableRecord(true, false)
	ts.tr.SetSQLConnection(db)

	return ts
}

// loadAllTestMapperTableRecord - carica tutte le istanze della classe
func loadAllTestMapperTableRecord() ([]*TestMapperTableRecord, error) {

	db := db.GetConnection()

	ts := NewTestMapperTableRecord(db)

	query := "SELECT " + AllField(ts) + " FROM " + ts.GetTableName()

	rows, err := db.Queryx(query)
	if err != nil {
		return nil, err
	}

	var result []*TestMapperTableRecord

	for rows.Next() {

		ts := NewTestMapperTableRecord(db)

		if err := LoadFromRow(rows, ts); err != nil {
			return nil, err
		}

		result = append(result, ts)
	}

	return result, nil
}

// GetTableRecord - Restituisce l'istanza di TableRecord
func (t TestMapperTableRecord) GetTableRecord() *TableRecord {
	return t.tr
}

// GetPrimaryKeyName - Restituisce il nome della chiave primaria
func (t TestMapperTableRecord) GetPrimaryKeyName() string {
	return "wrong_id"
}

// GetPrimaryKeyValue - Restituisce l'indirizzo di memoria del valore della chiave primaria
func (t TestMapperTableRecord) GetPrimaryKeyValue() int64 {
	return t.RecordID
}

// GetTableName - Restituisce il nome della tabella
func (t TestMapperTableRecord) GetTableName() string {
	return "users"
}

func TestMapper(t *testing.T) {

	if err := gotenv.Load("./../../../.env"); err != nil {
		t.Fatal("Errore caricamento configurazione")
	}

	tr := NewTestMapperTableRecord(db.GetConnection())

	fName := getFieldsNameNoPrimary(tr)

	if len(fName) != 4 {
		t.Error("Lunghezza errata")
	}

	fValue := getFieldsValueNoPrimary(tr)

	if len(fValue) != 4 {
		t.Error("Lunghezza errata")
	}
}
