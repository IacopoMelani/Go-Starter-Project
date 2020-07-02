package record

import (
	"strings"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db"
)

// genDeleteQuery - Si occupa di generare la query per la cancellazione del record
func genDeleteQuery(ti TableRecordInterface) string {

	query := "DELETE FROM " + ti.GetTableName() + " WHERE " + ti.GetPrimaryKeyName() + " = ?"

	return query
}

// genSaveQuery - Si occupa di generare la query di salvataggio
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

// genUpdateQuery - Si occupa di generare la query di aggiornamento
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

// getSaveFieldParams -  Si occupa di generare uno slice di "?" tanti quanti sono i parametri della query di inserimento
func getSaveFieldParams(ti TableRecordInterface) []string {

	fName := getFieldsNameNoPrimary(ti)

	s := make([]string, len(fName))

	for i := 0; i < len(fName); i++ {
		s[i] = "?"
	}

	return s
}

// getUpdateFiledParams - Si occupa di generare uno slice di "?" tanti quanti sono i parametri della query di aggiornamento
func getUpdateFieldParams(ti TableRecordInterface) []string {

	fName := getFieldsNameNoPrimary(ti)
	updateStmt := make([]string, len(fName))

	for i := 0; i < len(fName); i++ {
		updateStmt[i] = fName[i] + " = ?"
	}

	return updateStmt
}
