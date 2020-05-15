package refl

import (
	"errors"
	"reflect"
)

// GetStructFieldValueByTagName - Returns a value ptr to struct field by tag type and tag name
func GetStructFieldValueByTagName(c interface{}, tagType string, tagName string) interface{} {

	vPtr := reflect.ValueOf(c)

	t := reflect.TypeOf(c)
	v := reflect.Indirect(vPtr)

	for i := 0; i < v.NumField(); i++ {

		name := t.Elem().Field(i).Tag.Get(tagType)

		if !v.Field(i).CanInterface() || !v.Field(i).CanSet() || name != tagName {
			continue
		}

		return v.Field(i).Addr().Interface()
	}

	return errors.New("Field value not found for " + tagName)
}

// GetStructFieldsMapperByTagName - Restituisce i campi di mappatura di una struct, i due slice rappresentato:
// il primo uno slice con i nomi dei tag
// il secondo un ptr al valore del campo
func GetStructFieldsMapperByTagName(c interface{}, tagName string) (fieldsName []string, fieldsValue []interface{}) {

	vPtr := reflect.ValueOf(c)

	t := reflect.TypeOf(c)
	v := reflect.Indirect(vPtr)

	for i := 0; i < v.NumField(); i++ {

		if !v.Field(i).CanInterface() || !v.Field(i).CanSet() || t.Elem().Field(i).Tag.Get(tagName) == "" {
			continue
		}

		fieldsValue = append(fieldsValue, v.Field(i).Addr().Interface())
		fieldsName = append(fieldsName, t.Elem().Field(i).Tag.Get(tagName))
	}

	return
}

// GetStructFieldsNameAndTagByTagName - Restituisce u campi di mappatura di una struct, i due slice rappresentano:
// il primo uno slice con i nomi dei tag
// il secondo uno slice con i nomi dei campi della struct
func GetStructFieldsNameAndTagByTagName(c interface{}, tagName string) (tagFields []string, structFields []string) {

	s := reflect.ValueOf(c)
	v := reflect.Indirect(s)
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {

		if t.Field(i).Tag.Get(tagName) == "" {
			continue
		}

		tagFields = append(tagFields, t.Field(i).Tag.Get(tagName))
		structFields = append(structFields, t.Field(i).Name)
	}

	return
}
