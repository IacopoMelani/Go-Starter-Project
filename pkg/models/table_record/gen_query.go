package record

import (
	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db/driver"
)

// genDeleteQuery - Returns the delete query
func genDeleteQuery(ti TableRecordInterface) string {

	query := "DELETE FROM " + ti.GetTableName() + " WHERE " + ti.GetPrimaryKeyName() + " = ?"

	return query
}

// genSaveQuery - Returns the insert query
func genSaveQuery(ti TableRecordInterface) string {
	return driver.GetInsertQuery(ti.GetTableRecord().DriverName(), ti.GetTableName(), getFieldsNameNoPrimary(ti), getSaveFieldParams(ti))
}

// genUpdateQuery - Returns the update query
func genUpdateQuery(ti TableRecordInterface) string {
	return driver.GetUpdateQuery(ti.GetTableRecord().DriverName(), ti.GetTableName(), ti.GetPrimaryKeyName(), getUpdateFieldParams(ti))
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
