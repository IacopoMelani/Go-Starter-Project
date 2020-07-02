package refl

import (
	"errors"
	"reflect"
)

// GetStructFieldValueByTagName - Returns a value ptr to struct field by tag type and tag name
func GetStructFieldValueByTagName(c interface{}, tagType string, tagName string) (interface{}, error) {

	vPtr := reflect.ValueOf(c)

	t := reflect.TypeOf(c)
	v := reflect.Indirect(vPtr)

	for i := 0; i < v.NumField(); i++ {

		name := t.Elem().Field(i).Tag.Get(tagType)

		if !v.Field(i).CanInterface() || !v.Field(i).CanSet() || name != tagName {
			continue
		}

		return v.Field(i).Addr().Interface(), nil
	}

	return nil, errors.New("Field value not found for " + tagName)
}

// GetStructFieldsMapperByTagName - Returns two slices filtered by a tag name,
// the first is a slice of strings with struct tags value,
// the seconds is a slice of ptr to struct fields value
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

// GetStructFieldsNameAndTagByTagName - Returns two slices filtered by a tag name,
// the first is a slice of strings with struct tags value,
// the seconds is a slice of string with struct fields name
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

// GetType - Returns the name of var, appends "*" if var is ptr to value
func GetType(val interface{}) (typeName string) {
	typeName = reflect.TypeOf(val).String()
	return
}
