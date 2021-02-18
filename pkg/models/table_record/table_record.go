package record

import (
	"strings"

	refl "github.com/IacopoMelani/Go-Starter-Project/pkg/helpers/reflect"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db"
	builder "github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db/query_builder"
	"github.com/jmoiron/sqlx"
)

// NewTableModel -  Defines a func that return a TableRecordInterface
type NewTableModel func() TableRecordInterface

// TableRecordInterface - Defines a generics struct to interact with database
type TableRecordInterface interface {
	GetTableRecord() *TableRecord
	GetPrimaryKeyName() string
	GetTableName() string
}

// TableRecord - Common struct to access all utility func with TableRecordInterface
// implements QueryBuilderInterface
type TableRecord struct {
	builder.Builder
	isLoaded   bool
	isNew      bool
	isReadOnly bool
	db         db.SQLConnector
}

// getTableRecordConnection - Returns the connection of TableRecordInterface
func getTableRecordConnection(ti TableRecordInterface) db.SQLConnector {
	return ti.GetTableRecord().db
}

// AllField - Return all select fields
func AllField(ti TableRecordInterface) string {

	fieldName, _ := GetFieldMapper(ti)

	return strings.Join(fieldName, ",")
}

// GetPrimaryKeyValue - Returs the primary key value
func GetPrimaryKeyValue(ti TableRecordInterface) interface{} {
	value, err := refl.GetStructFieldValueByTagName(ti, "db", ti.GetPrimaryKeyName())
	if err != nil {
		panic(err)
	}
	return value
}

// LoadFromRow - Loads the passed TableRecordInterface with the row sql result
func LoadFromRow(r *sqlx.Rows, tri TableRecordInterface) error {

	if err := r.StructScan(tri); err != nil {
		return err
	}

	tr := tri.GetTableRecord()

	tr.SetIsNew(false)
	tr.isLoaded = true

	return nil
}

// NewTableRecord - Returns a new instance of TableRecord
func NewTableRecord(db db.SQLConnector, isNew bool, isReadOnly bool) *TableRecord {

	tr := new(TableRecord)
	tr.isNew = isNew
	tr.isReadOnly = isReadOnly
	tr.db = db

	return tr
}

// GetDB - Returns TableRecord db resource instance
func (t *TableRecord) GetDB() db.SQLConnector {
	return t.db
}

// IsLoaded - Returns if TableRecord is loaded successfully, might useful after a LoadByID or similar func
func (t *TableRecord) IsLoaded() bool {
	return t.isLoaded
}

// IsNew - Returns if the record is new
func (t *TableRecord) IsNew() bool {
	return t.isNew
}

// DriverName - Returns the sql driver name for TableRecord's instance
func (t *TableRecord) DriverName() string {
	return t.db.DriverName()
}

// PrepareStmt - Returns the query stmt built with QueryBuilder
func (t *TableRecord) PrepareStmt(tableName string) (*sqlx.Stmt, error) {

	db := t.db

	query := t.BuildQuery(tableName)

	stmt, err := db.Preparex(query)
	if err != nil {
		return nil, err
	}

	return stmt, nil
}

// SetIsNew - Sets isNew field
func (t *TableRecord) SetIsNew(new bool) *TableRecord {
	t.isNew = new
	return t
}

// SetSQLConnection - Sets the db resource instance
func (t *TableRecord) SetSQLConnection(db db.SQLConnector) *TableRecord {
	t.db = db
	return t
}
