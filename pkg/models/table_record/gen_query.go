package record

import (
	"strings"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/helpers/slice"
)

// genDeleteQuery - Si occupa di generare la query per la cancellazione del record
func genDeleteQuery(ti TableRecordInterface) string {

	query := "DELETE FROM " + ti.GetTableName() + " WHERE " + ti.GetPrimaryKeyName() + " = ?"

	return query
}

// genSaveQuery - Si occupa di generare la query di salvataggio
func genSaveQuery(ti TableRecordInterface) string {

	fName, _ := GetFieldMapper(ti)

	query := "INSERT INTO " + ti.GetTableName() + " (" + strings.Join(fName, ", ") + ") VALUES ( " + strings.Join(getSaveFieldParams(ti), ", ") + " )"

	return query
}

// genUpdateQuery - Si occupa di generare la query di aggiornamento
func genUpdateQuery(ti TableRecordInterface) string {

	query := "UPDATE  " + ti.GetTableName() + " SET " + strings.Join(getUpdateFieldParams(ti), ", ") + " WHERE " + ti.GetPrimaryKeyName() + " = ?"
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

// getUpdateFiledParams - Si occupa di generare uno slice di "?" tanti quanti sono i parametri della query di aggiornamento
func getUpdateFieldParams(ti TableRecordInterface) []string {

	fName, _ := GetFieldMapper(ti)

	tempFName := fName

	for i, name := range fName {
		if name == ti.GetPrimaryKeyName() {
			tempFName = slice.RemoveString(fName, i)
			break
		}
	}

	fName = tempFName

	updateStmt := make([]string, len(fName))

	for i := 0; i < len(fName); i++ {
		updateStmt[i] = fName[i] + " = ?"
	}

	return updateStmt
}
