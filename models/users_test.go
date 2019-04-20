package models

import (
	"strings"
	"testing"

	"github.com/subosito/gotenv"
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

func TestGetSelectQuery(t *testing.T) {

	u := User{}

	querySQL, _ := u.GetSelectQuery()

	queryArray := strings.Split(querySQL, " ")
	if queryArray[0] != "SELECT" {
		t.Error("Query errata, dovrebbe iniziare con 'SELECT'")
	}

}

func TestLoadAllUser(t *testing.T) {

	gotenv.Load("../.env")

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
