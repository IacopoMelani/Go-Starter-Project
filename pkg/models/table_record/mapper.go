package record

import (
	refl "github.com/IacopoMelani/Go-Starter-Project/pkg/helpers/reflect"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/helpers/slice"
)

const dbTagName = "db"

// getFieldsNameNoPrimary - Restituisce tutti i campi del model ad eccezione della chiave primaria
func getFieldsNameNoPrimary(ti TableRecordInterface) []string {

	fName, _ := GetFieldMapper(ti)

	for i, name := range fName {
		if name == ti.GetPrimaryKeyName() {
			return slice.RemoveString(fName, i)
		}
	}

	return fName
}

// getFieldsValueNoPrimary - Restituisce tutti i campi di mappatura ad esclusione della chiave primaria
func getFieldsValueNoPrimary(ti TableRecordInterface) []interface{} {

	fName, fValue := GetFieldMapper(ti)

	for i := 0; i < len(fName); i++ {
		if fName[i] == ti.GetPrimaryKeyName() {
			return slice.Remove(fValue, i)
		}
	}

	return fValue
}

// GetFieldMapper - Si occupa di recuperare in reflection i nomi dei tag "db" e l'indirizzo del valore del campo
func GetFieldMapper(ti TableRecordInterface) (fieldsName []string, fieldsValue []interface{}) {
	fieldsName, fieldsValue = refl.GetStructFieldsMapperByTagName(ti, dbTagName)
	return
}
