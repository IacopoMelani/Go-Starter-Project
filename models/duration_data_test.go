package models

import (
	"testing"
)

func TestDurationData(t *testing.T) {
	d := GetData()
	if d == nil {
		t.Error("Errore, impossibile istanziare duration data")
	}

	tempData := 1234567890
	d.SetContent(tempData, 0)

	dataContent, err := d.GetContent()
	if err != nil {
		d.SetContent(tempData, 15)
	}
	dataContent, _ = d.GetContent()

	if tempData != dataContent.(int) {
		t.Error("Dati non coerenti")
	}

}
