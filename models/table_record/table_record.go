package record

import (
	"database/sql"
	"errors"
	"reflect"
	"strings"

	"github.com/IacopoMelani/Go-Starter-Project/db"
	builder "github.com/IacopoMelani/Go-Starter-Project/db/query_builder"
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
	RecordID   int64
	isNew      bool
	isReadOnly bool
	builder.Builder
}

// executeSaveUpdateQuery - Si occupa di eseguire fisicamente la query, in caso di successo restituisce l'Id appena inserito
func executeSaveUpdateQuery(query string, params []interface{}) (int64, error) {

	db := db.GetConnection()

	res, err := db.Exec(query, params...)
	if err != nil {
		return 0, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastID, nil
}

// genDeleteQuery - Si occupa di generare la query per la cancellazione del record
func genDeleteQuery(ti TableRecordInterface) string {

	query := "DELETE FROM " + ti.GetTableName() + " WHERE " + ti.GetPrimaryKeyName() + " = ?"

	return query
}

// getSaveFieldParams -  Si occupa di generare uno slice di "?" tanti quanti sono i parametri della query di inserimento
func getSaveFieldParams(ti TableRecordInterface) []string {

	fName, _ := GetFieldMapper(ti)

	s := make([]string, len(fName))

	for i := 0; i < len(fName); i++ {
		s[i] = "?"
	}

	return s
}

// genSaveQuery - Si occupa di generare la query di salvataggio
func genSaveQuery(ti TableRecordInterface) string {

	fName, _ := GetFieldMapper(ti)

	query := "INSERT INTO " + ti.GetTableName() + " (" + strings.Join(fName, ", ") + ") VALUES ( " + strings.Join(getSaveFieldParams(ti), ", ") + " )"

	return query
}

// getUpdateFiledParams - Si occupa di generare uno slice di "?" tanti quanti sono i parametri della query di aggiornamento
func getUpdateFieldParams(ti TableRecordInterface) []string {

	fName, _ := GetFieldMapper(ti)

	updateStmt := make([]string, len(fName))

	for i := 0; i < len(fName); i++ {
		updateStmt[i] = fName[i] + " = ?"
	}

	return updateStmt
}

// genUpdateQuery - Si occupa di generare la query di aggiornamento
func genUpdateQuery(ti TableRecordInterface) string {

	query := "UPDATE  " + ti.GetTableName() + " SET " + strings.Join(getUpdateFieldParams(ti), ", ") + " WHERE " + ti.GetPrimaryKeyName() + " = ?"
	return query
}

// AllField - Restitusice tutti i campi per la select *
func AllField(ti TableRecordInterface) string {

	fieldName, _ := GetFieldMapper(ti)

	fieldName = append([]string{ti.GetPrimaryKeyName()}, fieldName...)

	return strings.Join(fieldName, ",")
}

// Delete - Si occupa di cancellare un record sul database
func Delete(ti TableRecordInterface) (int64, error) {

	t := ti.GetTableRecord()

	db := db.GetConnection()

	stmt, err := db.Prepare(genDeleteQuery(ti))
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(t.RecordID)
	if err != nil {
		return 0, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return 0, nil
	}

	return rows, nil
}

// ExecQuery - Esegue la query costruita con QueryBuilder
func ExecQuery(ti TableRecordInterface, ntm NewTableModel) ([]TableRecordInterface, error) {

	t := ti.GetTableRecord()

	stmt, err := t.PrepareStmt(ti.GetTableName())
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(t.Params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tiList []TableRecordInterface

	for rows.Next() {

		nti := ntm()

		if err := LoadFromRow(rows, nti); err != nil {
			return nil, err
		}
		
		tiList = append(tiList, nti)
	}

	ti.GetTableRecord().ResetStmt()

	return tiList, nil
}

// GetFieldMapper - Si occupa di recuperare in reflection i nomi dei tag "db" e l'indirizzo del valore del campo
func GetFieldMapper(ti TableRecordInterface) ([]string, []interface{}) {

	vPtr := reflect.ValueOf(ti)

	t := reflect.TypeOf(ti)
	v := reflect.Indirect(vPtr)

	var fieldName []string
	var fieldValue []interface{}

	for i := 0; i < v.NumField(); i++ {

		if !v.Field(i).CanInterface() || !v.Field(i).CanSet() {
			continue
		}

		fieldValue = append(fieldValue, v.Field(i).Addr().Interface())
		fieldName = append(fieldName, t.Elem().Field(i).Tag.Get("db"))
	}

	return fieldName, fieldValue
}

// LoadByID - Carica l'istanza passata con i valori della sua tabella ricercando per chiave primaria
func LoadByID(ti TableRecordInterface, id int64) error {

	db := db.GetConnection()

	query := "SELECT " + AllField(ti) + " FROM " + ti.GetTableName() + " WHERE " + ti.GetPrimaryKeyName() + " = ?"

	params := []interface{}{interface{}(id)}

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Query(params...)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {

		if err := LoadFromRow(rows, ti); err != nil {
			return err
		}
	}

	return nil
}

// LoadFromRow - Si occupa di caricare la struct dal result - row della query
func LoadFromRow(r *sql.Rows, tri TableRecordInterface) error {

	_, vField := GetFieldMapper(tri)

	dest := append([]interface{}{&tri.GetTableRecord().RecordID}, vField...)

	if err := r.Scan(dest...); err != nil {
		return err
	}

	tri.GetTableRecord().SetIsNew(false)

	return nil
}

// NewTableRecord - Restituisce una nuova istanza di TableRecord
func NewTableRecord(isNew bool, isReadOnly bool) *TableRecord {

	tr := new(TableRecord)
	tr.isNew = isNew
	tr.isReadOnly = isReadOnly

	return tr
}

// Save - Si occupa di eseguire il salvataggio della TableRecord eseguendo un inserimento se TableRecord::isNew risulta false, altrimenti ne aggiorna il valore
func Save(ti TableRecordInterface) error {

	t := ti.GetTableRecord()

	if t.isReadOnly {
		return errors.New("Read-only model")
	}

	if t.isNew {

		query := genSaveQuery(ti)
		_, fValue := GetFieldMapper(ti)
		id, err := executeSaveUpdateQuery(query, fValue)
		if err != nil {
			return err
		}

		t.RecordID = id
		t.SetIsNew(false)
	} else {

		query := genUpdateQuery(ti)
		_, fValue := GetFieldMapper(ti)
		_, err := executeSaveUpdateQuery(query, append(fValue, ti.GetTableRecord().RecordID))
		if err != nil {
			return err
		}
	}

	return nil
}

// IsNew - Restituisce se il record è nuovo
func (t *TableRecord) IsNew() bool {
	return t.isNew
}

// SetIsNew - Si occupa di impostare il valore del campo TableRecord::isNews
func (t *TableRecord) SetIsNew(new bool) *TableRecord {
	t.isNew = new
	return t
}

// PrepareStmt - Restituisce lo stmt della query pronta da essere eseguita
func (t *TableRecord) PrepareStmt(tableName string) (*sql.Stmt, error) {

	db := db.GetConnection()

	query := t.BuildQuery(tableName)

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	return stmt, nil
}
