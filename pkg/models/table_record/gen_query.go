package record

import (
	"strings"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db"
)

// genDeleteQuery - Returns the delete query
func genDeleteQuery(ti TableRecordInterface) string {

	query := "DELETE FROM " + ti.GetTableName() + " WHERE " + ti.GetPrimaryKeyName() + " = ?"

	return query
}

// genSaveQuery - Returns the insert query
func genSaveQuery(ti TableRecordInterface) string {

	fName := getFieldsNameNoPrimary(ti)

	var query string

	switch ti.GetTableRecord().DriverName() {

	case db.DriverSQLServer:
		query = "INSERT INTO " + ti.GetTableName() + " (" + strings.Join(fName, ", ") + ") OUTPUT INSERTED.* VALUES ( " + strings.Join(getSaveFieldParams(ti), ", ") + " )"

	case db.DriverMySQL:
		query = "INSERT INTO " + ti.GetTableName() + " (" + strings.Join(fName, ", ") + ") VALUES ( " + strings.Join(getSaveFieldParams(ti), ", ") + " )"
	}

	return query
}

// genUpdateQuery - Returns the update query
func genUpdateQuery(ti TableRecordInterface) string {

	var query string

	switch ti.GetTableRecord().DriverName() {

	case db.DriverSQLServer:
		query = "UPDATE  " + ti.GetTableName() + " SET " + strings.Join(getUpdateFieldParams(ti), ", ") + " OUTPUT INSERTED.* WHERE " + ti.GetPrimaryKeyName() + " = ?"

	case db.DriverMySQL:
		query = "UPDATE  " + ti.GetTableName() + " SET " + strings.Join(getUpdateFieldParams(ti), ", ") + " WHERE " + ti.GetPrimaryKeyName() + " = ?"
	}

	return query
}

// getSaveFieldParams - Returns a slice of "?" as many as the parameters of the insert query
func getSaveFieldParams(ti TableRecordInterface) []string {

	fName := getFieldsNameNoPrimary(ti)

	s := make([]string, len(fName))

	for i := 0; i < len(fName); i++ {
		s[i] = "?"
	}

	return s
}

// getUpdateFiledParams - Returns a slice of "?" as many as the parameters of the update query
func getUpdateFieldParams(ti TableRecordInterface) []string {

	fName := getFieldsNameNoPrimary(ti)
	updateStmt := make([]string, len(fName))

	for i := 0; i < len(fName); i++ {
		updateStmt[i] = fName[i] + " = ?"
	}

	return updateStmt
}
