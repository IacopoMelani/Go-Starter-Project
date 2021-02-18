package record

import (
	"os"
	"testing"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db"
	"github.com/subosito/gotenv"
)

// TestMapperTableRecord - Example
type TestMapperTableRecord struct {
	tr       *TableRecord
	RecordID int64   `json:"id" db:"record_id"`
	Name     *string `json:"name" db:"name"`
	Lastname *string `json:"lastname" db:"lastname"`
	Gender   *string `json:"gender" db:"gender"`
}

// NewTestMapperTableRecord - Example
func NewTestMapperTableRecord(db db.SQLConnector) *TestMapperTableRecord {

	ts := new(TestMapperTableRecord)
	ts.tr = NewTableRecord(db, true, false)

	return ts
}

// GetTableRecord - Example
func (t TestMapperTableRecord) GetTableRecord() *TableRecord {
	return t.tr
}

// GetPrimaryKeyName - Example
func (t TestMapperTableRecord) GetPrimaryKeyName() string {
	return "wrong_id"
}

// GetTableName - Example
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
