package models

import (
	"strings"
	"testing"
)

func TestGetSaveQueryAndSetRecordID(t *testing.T) {

	u := new(User)

	u.Name = "Paolo"
	u.Lastname = "Rossi"
	u.Gender = "M"

	queryString := strings.Split(u.GetSaveQuery(), " ")

	if queryString[0] != "INSERT" {
		t.Error("Errore: la prima parola dovrebbe essere INSERT")
	}

	u.RecordID = 1

	u.SetRecordID(2)

	if u.RecordID != 2 {
		t.Error("Errore: il record dovrebbe essere 2")
	}

	queryString = strings.Split(u.GetSaveQuery(), " ")

	if queryString[0] != "UPDATE" {
		t.Error("Errore: la prima parola dovrebbe essere UPDATE")
	}

}

func TestLoadAllUser(t *testing.T) {

	usersList, err := LoadAllUser()
	if err != nil {
		t.Error(err.Error())
	}

	if len(usersList) <= 0 {
		t.Error("Errore nessun risultato prelevato")
	}

	if usersList[0].RecordID == usersList[1].RecordID {
		t.Error("Errore record duplicati")
	}
}
