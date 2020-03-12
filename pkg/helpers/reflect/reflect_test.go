package refl

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

}
