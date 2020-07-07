package record

import (
	refl "github.com/IacopoMelani/Go-Starter-Project/pkg/helpers/reflect"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/helpers/slice"
)

const dbTagName = "db"

// getFieldsNameNoPrimary - Returns all fields name except the primary key
func getFieldsNameNoPrimary(ti TableRecordInterface) []string {

	fName, _ := GetFieldMapper(ti)

	for i, name := range fName {
		if name == ti.GetPrimaryKeyName() {
			return slice.RemoveString(fName, i)
		}
	}

	return fName
}

// getFieldsValueNoPrimary -  Return all struct fields with tags "db" execpt primary key struct field
func getFieldsValueNoPrimary(ti TableRecordInterface) []interface{} {

	fName, fValue := GetFieldMapper(ti)

	for i := 0; i < len(fName); i++ {
		if fName[i] == ti.GetPrimaryKeyName() {
			return slice.Remove(fValue, i)
		}
	}

	return fValue
}

// GetFieldMapper - Returns a slice of table name fields and a slice of struct field ptr
func GetFieldMapper(ti TableRecordInterface) (fieldsName []string, fieldsValue []interface{}) {
	fieldsName, fieldsValue = refl.GetStructFieldsMapperByTagName(ti, dbTagName)
	return
}
