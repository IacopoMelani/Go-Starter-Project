package models

import (
	record "Go-Starter-Project/models/table_record"
	"fmt"
	"testing"

	"github.com/subosito/gotenv"
)

func TestTableMirror(t *testing.T) {

	gotenv.Load("./../.env")

	u := NewUser()

	u.Name = "Mario"
	u.Lastname = "Rossi"
	u.Gender = "M"

	err := record.Save(u)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(u)

	if u.tr.RecordID == 0 {
		t.Error("Chiave non salvata")
	}

	tempName := u.Name
	tempID := u.tr.RecordID

	err = record.LoadByID(u, u.GetTableRecord().RecordID)
	if err != nil {
		panic(err.Error())
	}

	if tempName != u.Name {
		t.Error("Campi non uguali")
	}

	u.Name = "Marco"

	err = record.Save(u)
	if err != nil {
		panic(err.Error())
	}

	if tempID != u.tr.RecordID {
		t.Error("Chiave primaria è cambiata durante l'update")
	}

	fmt.Println(u)

	usersList, err := LoadAllUsers()
	if err != nil {
		panic(err.Error())
	}

	if len(usersList) == 0 {
		t.Error("Lunghezza inferiore al minimo nel contesto del test")
	}

	for _, user := range usersList {
		fmt.Println(user)
	}

}
