package refl

import (
	"reflect"
)

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
	i := reflect.Indirect(s)
	t := i.Type()

	for i := 0; i < t.NumField(); i++ {

		if t.Field(i).Tag.Get(tagName) == "" {
			continue
		}

		tagFields = append(tagFields, t.Field(i).Tag.Get(tagName))
		structFields = append(structFields, t.Field(i).Name)
	}

	return
}
