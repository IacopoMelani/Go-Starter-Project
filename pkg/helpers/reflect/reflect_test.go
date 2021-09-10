package refl

//lint:file-ignore U1000 Ignore all unused code, it's test code.

import (
	"testing"
)

type useless struct{}

type testStruct struct {
	useless
	One   int `number:"ONE"`
	Two   int `number:"TWO"`
	Three int `number:"THREE"`
}

const testTagName = "number"

func TestRefl(t *testing.T) {

	ts := &testStruct{}

	fValue, fName := GetStructFieldsMapperByTagName(ts, testTagName)

	if len(fValue) != len(fName) || len(fValue) != 3 {
		t.Error("Lunghezza slice errata")
	}

	fValue, nfName := GetStructFieldsNameAndTagByTagName(ts, testTagName)

	if len(fValue) != len(nfName) || len(fValue) != 3 {
		t.Error("Lunghezza slice errata")
	}

	ts.One = 1

	value, err := GetStructFieldValueByTagName(ts, testTagName, "ONE")
	if err != nil {
		t.Error(err.Error())
	}

	if *value.(*int) != ts.One {
		t.Error("Errore: dovrebbe essere lo stesso valore")
	}

	_, err = GetStructFieldValueByTagName(ts, testTagName, "ON2")
	if err == nil {
		t.Error("Dovrebbe essere Error")
	}

	valueType := "test"
	valueTypePtr := &valueType

	if GetType(valueType) != "string" {
		t.Error("value should be 'string'")
	}
	if GetType(valueTypePtr) != "*string" {
		t.Error("value should be '*string'")
	}
}
