package slice

import (
	"testing"
)

func TestSliceInterface(t *testing.T) {

	var slice = []interface{}{"1", 2, 3.2, "44"}

	slice = Remove(slice, 3)

	if slice[2] != 3.2 || len(slice) != 3 {
		t.Fatal("Errore, non è stato eseguita la corretta eliminazione")
	}

	slice = []interface{}{}

	slice = Remove(slice, 4)

	if len(slice) > 0 {
		t.Fatal("Errore durante cancellazione slice vuoto")
	}
}

func TestSliceString(t *testing.T) {

	var sliceString = []string{"1", "2", "3", "4"}

	sliceString = RemoveString(sliceString, 3)

	if sliceString[2] != "3" || len(sliceString) != 3 {
		t.Fatal("Errore, non è stato eseguita la corretta eliminazione")
	}

	sliceString = []string{}

	sliceString = RemoveString(sliceString, 4)

	if len(sliceString) > 0 {
		t.Fatal("Errore durante cancellazione slice vuoto")
	}
}
