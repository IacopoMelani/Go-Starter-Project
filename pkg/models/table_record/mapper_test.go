package record

import (
	"os"
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

// GetTableRecord - Restituisce l'istanza di TableRecord
func (t TestMapperTableRecord) GetTableRecord() *TableRecord {
	return t.tr
}

// GetPrimaryKeyName - Restituisce il nome della chiave primaria
func (t TestMapperTableRecord) GetPrimaryKeyName() string {
	return "wrong_id"
}

// GetTableName - Restituisce il nome della tabella
func (t TestMapperTableRecord) GetTableName() string {
	return "users"
}

func TestMapper(t *testing.T) {

	if err := gotenv.Load("./../../../.env"); err != nil {
		t.Fatal("Errore caricamento configurazione")
	}
	db.InitConnection(os.Getenv("SQL_DRIVER"), os.Getenv("STRING_CONNECTION"))

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
