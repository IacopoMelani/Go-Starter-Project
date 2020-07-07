package record

import (
	"strings"

	refl "github.com/IacopoMelani/Go-Starter-Project/pkg/helpers/reflect"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db"
	builder "github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db/query_builder"
	"github.com/jmoiron/sqlx"
)

// NewTableModel - Tipo per definire una funzione che restituisce una TableRecordInterface
type NewTableModel func() TableRecordInterface

// TableRecordInterface - interfaccia che definisce una generica struct che permette l'interazione con TableRecord
type TableRecordInterface interface {
	GetTableRecord() *TableRecord
	GetPrimaryKeyName() string
	GetTableName() string
}

// TableRecord - Struct per l'implementazione di TableRecordInterface
// implementa QueryBuilderInterface
type TableRecord struct {
	builder.Builder
	isLoaded   bool
	isNew      bool
	isReadOnly bool
	db         db.SQLConnector
}

// getTableRecordConnection - Restituisce la connessione di un TableRecordInterface
func getTableRecordConnection(ti TableRecordInterface) db.SQLConnector {
	return ti.GetTableRecord().db
}

// AllField - Restitusice tutti i campi per la select *
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

// LoadFromRow - Si occupa di caricare la struct dal result - row della query
func LoadFromRow(r *sqlx.Rows, tri TableRecordInterface) error {

	if err := r.StructScan(tri); err != nil {
		return err
	}

	tr := tri.GetTableRecord()

	tr.SetIsNew(false)
	tr.SetSQLConnection(tri.GetTableRecord().db)
	tr.isLoaded = true

	return nil
}

// NewTableRecord - Restituisce una nuova istanza di TableRecord
func NewTableRecord(isNew bool, isReadOnly bool) *TableRecord {

	tr := new(TableRecord)
	tr.isNew = isNew
	tr.isReadOnly = isReadOnly

	return tr
}

// GetDB - Restituisce la risorsa di connessione al database
func (t *TableRecord) GetDB() db.SQLConnector {
	return t.db
}

// IsLoaded - Restituisce se TableRecord è stato caricato correttamente
func (t *TableRecord) IsLoaded() bool {
	return t.isLoaded
}

// IsNew - Restituisce se il record è nuovo
func (t *TableRecord) IsNew() bool {
	return t.isNew
}

// DriverName - Returns the sql driver name for TableRecord's instance
func (t *TableRecord) DriverName() string {
	return t.db.DriverName()
}

// PrepareStmt - Restituisce lo stmt della query pronta da essere eseguita
func (t *TableRecord) PrepareStmt(tableName string) (*sqlx.Stmt, error) {

	db := t.db

	query := t.BuildQuery(tableName)

	stmt, err := db.Preparex(query)
	if err != nil {
		return nil, err
	}

	return stmt, nil
}

// SetIsNew - Si occupa di impostare il valore del campo TableRecord::isNews
func (t *TableRecord) SetIsNew(new bool) *TableRecord {
	t.isNew = new
	return t
}

// SetSQLConnection - Imposta la connessione
func (t *TableRecord) SetSQLConnection(db db.SQLConnector) *TableRecord {
	t.db = db
	return t
}
