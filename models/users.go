package models

import (
	"fmt"
	"log"
	"testDB/db"
)

// ColUserRecordID - definisce il nome del campo RecordID
const ColUserRecordID = "record_id"

// ColUserName - definisce il nome del campo name
const ColUserName = "name"

// ColUserLastname - definisce il nome del campo lastname
const ColUserLastname = "lastname"

// ColUserGender - definisce il nome del campo gender
const ColUserGender = "gender"

// TableUser - Definisce il nome della tabella "users"
const TableUser = "users"

// User - Struct che definisce la tabella "users"
type User struct {
	TableMirror
	recordID int
	name     string
	lastname string
	gender   string
}

// UserList - Tipo che definisce una lista di struct di User
type UserList []User

// LoadAllUser - Carica tutti gli utenti presenti in tabella
func LoadAllUser() UserList {

	db := db.GetConnection()

	rows, err := db.Query("SELECT * FROM USERS")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	var users []User
	for rows.Next() {

		u := User{}
		err := rows.Scan(&u.recordID, &u.name, &u.lastname, &u.gender)
		if err != nil {
			log.Fatal(err.Error())
		}

		users = append(users, u)
	}
	return users
}

// GetPrimaryKeyName - Restituisce il nome della chiave primaria
func (u User) GetPrimaryKeyName() string {
	return ColUserRecordID
}

// GetTableName - Restituisce il nome della tabella
func (u User) GetTableName() string {
	return TableUser
}

// PrintAll - Stampa tutti i valori degli utenti
func (u UserList) PrintAll() {
	for _, user := range u {
		fmt.Println(user)
	}
}

// Save -
func (u *User) Save() (int, error) {
	id, err := Save(&u.TableMirror, u)
	if err != nil {
		return -1, err
	}
	u.recordID = id
	return id, nil
}

// SetName - imposta il nome dell'utente
func (u *User) SetName(name string) *User {
	u.name = name
	u.TableMirror.SetField(ColUserName, u.name)
	return u
}
