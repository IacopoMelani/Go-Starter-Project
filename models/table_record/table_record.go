package record

import (
	"Go-Starter-Project/db"
	"strings"
)

// TableRecordInterface -
type TableRecordInterface interface {
	getTableRecord() *TableRecord
	GetPrimaryKeyName() string
	GetPrimaryKeyValue() int64
	GetTableName() string
	GetFieldMapper() ([]string, []*interface{})
}

// TableRecord -
type TableRecord struct {
	isNew bool
}

func executeSaveUpdateQuery(query string, params []*interface{}) (int64, error) {

	db := db.GetConnection()

	res, err := db.Exec(query, getParamsFromSlicePointer(params)...)
	if err != nil {
		return 0, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastID, nil
}

func getParamsFromSlicePointer(params []*interface{}) []interface{} {

	s := make([]interface{}, len(params))

	for _, v := range params {
		s = append(s, *v)
	}

	return s
}

func getSaveFieldParams(ti TableRecordInterface) []string {

	fName, _ := ti.GetFieldMapper()

	s := make([]string, len(fName))

	for i := 0; i < len(fName); i++ {
		if fName[i] == ti.GetPrimaryKeyName() {
			continue
		}
		s[i] = "?"
	}

	return s
}

func getUpdateFiledParams(ti TableRecordInterface) []string {

	var updateStmt []string

	fName, _ := ti.GetFieldMapper()

	for i := 0; i < len(fName); i++ {
		if fName[i] == ti.GetPrimaryKeyName() {
			continue
		}
		updateStmt[i] = fName[i] + " = ?"
	}

	return updateStmt
}

func genSaveQuery(ti TableRecordInterface) string {

	fName, _ := ti.GetFieldMapper()

	query := "INSERT INTO " + ti.GetTableName() + " (" + strings.Join(fName, ", ") + ") VALUES ( " + strings.Join(getSaveFieldParams(ti), ", ") + " )"

	return query
}

func genUpdateQuery(ti TableRecordInterface) string {

	query := "UPDATE  " + ti.GetTableName() + " SET " + strings.Join(getUpdateFiledParams(ti), ", ") + ") WHERE " + ti.GetPrimaryKeyName() + " = ?"
	return query
}

// Save -
func Save(ti TableRecordInterface) (TableRecordInterface, error) {

	t := ti.getTableRecord()

	db := db.GetConnection()

	if t.isNew {

		/* 		query := genSaveQuery(ti)
		   		_, fValue := ti.GetFieldMapper()
		   		executeQuery(query, fValue)
		*/
	} else {

		query := genUpdateQuery(ti)
		_, err := db.Prepare(query)
		if err != nil {
			return nil, err
		}

	}

	return ti, nil
}

// SetIsNew - Si occupa di impostare il valore del campo TableRecord::isNews
func (t *TableRecord) SetIsNew(new bool) *TableRecord {
	t.isNew = new
	return t
}
