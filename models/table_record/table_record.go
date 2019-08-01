package record

import (
	"Go-Starter-Project/db"
	"reflect"
	"strings"
)

// TableRecordInterface - interfaccia che definisce una generica struct che permette l'interazione con TableRecord
type TableRecordInterface interface {
	GetTableRecord() *TableRecord
	GetPrimaryKeyName() string
	GetTableName() string
	New()
}

// TableRecord - Struct per l'implementazione di TableRecordInterface
type TableRecord struct {
	RecordID int64
	isNew    bool
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

// GetFieldMapper - Si occupa di recuperare in reflection i nomi dei tag "db" e l'indirizzo del valore del campo
func GetFieldMapper(ti TableRecordInterface) ([]string, []interface{}) {

	vPtr := reflect.ValueOf(ti)

	t := reflect.TypeOf(ti)
	v := reflect.Indirect(vPtr)

	var fieldName []string
	var fieldValue []interface{}

	for i := 0; i < v.NumField(); i++ {

		if !v.Field(i).CanSet() {
			continue
		}

		fieldValue = append(fieldValue, v.Field(i).Addr().Interface())
		fieldName = append(fieldName, t.Elem().Field(i).Tag.Get("db"))
	}

	return fieldName, fieldValue
}

// getSaveFieldParams -  Si occupa di generare uno slice di "?" tanti quanti sono i parametri della query di inserimento
func getSaveFieldParams(ti TableRecordInterface) []string {

	fName, _ := GetFieldMapper(ti)

	s := make([]string, len(fName))

	for i := 0; i < len(fName); i++ {
		if fName[i] == ti.GetPrimaryKeyName() {
			continue
		}
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
		if fName[i] == ti.GetPrimaryKeyName() {
			continue
		}
		updateStmt[i] = fName[i] + " = ?"
	}

	return updateStmt
}

// genUpdateQuery - Si occupa di generare la query di aggiornamento
func genUpdateQuery(ti TableRecordInterface) string {

	query := "UPDATE  " + ti.GetTableName() + " SET " + strings.Join(getUpdateFieldParams(ti), ", ") + " WHERE " + ti.GetPrimaryKeyName() + " = ?"
	return query
}

// LoadByID - Carica l'istanza passata con i valori della sua tabella ricercando per chiave primaria
func LoadByID(ti TableRecordInterface, id int64) error {

	db := db.GetConnection()

	query := "SELECT * FROM " + ti.GetTableName() + " WHERE " + ti.GetPrimaryKeyName() + " = ?"

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

		_, vField := GetFieldMapper(ti)

		params := append([]interface{}{&ti.GetTableRecord().RecordID}, vField...)

		err := rows.Scan(params...)
		if err != nil {
			return err
		}

		ti.GetTableRecord().SetIsNew(false)
	}

	return nil
}

// Save - Si occupa di eseguire il salvataggio della TableRecord eseguendo un inserimento se TableRecord::isNew risulta false, altrimenti ne aggiorna il valore
func Save(ti TableRecordInterface) error {

	t := ti.GetTableRecord()

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

// SetIsNew - Si occupa di impostare il valore del campo TableRecord::isNews
func (t *TableRecord) SetIsNew(new bool) *TableRecord {
	t.isNew = new
	return t
}
