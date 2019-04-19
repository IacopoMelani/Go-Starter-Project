package models

import (
	"Go-Starter-Project/db"
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
	RecordID int
	Name     string
	Lastname string
	Gender   string
}

// UsersList - Tipo che definisce una lista di struct di User
type UsersList []User

// LoadAllUser - Restituisce la lista di tutti gli utenti
func LoadAllUser() (UsersList, error) {

	db := db.GetConnection()

	querySQL := "SELECT * FROM " + TableUser
	rows, err := db.Query(querySQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usersList UsersList

	for rows.Next() {

		u := User{}

		err := rows.Scan(&u.RecordID, &u.Name, &u.Lastname, &u.Gender)
		if err != nil {
			return nil, err
		}

		usersList = append(usersList, u)
	}

	return usersList, nil
}

// GetSaveQuery - Restituisce una query di inserimento nel caso in cui il record sia nuovo, altrimenti di modifica
func (u *User) GetSaveQuery() string {
	if u.RecordID == 0 {
		return "INSERT INTO " + TableUser + " (" + ColUserName + "," + ColUserLastname + ", " + ColUserGender + ") VALUES (?, ?, ?)"
	}
	return "UPDATE " + TableUser + " SET " + ColUserName + " = ?, " + ColUserLastname + "=?, " + ColUserGender + "=? WHERE " + ColUserRecordID + " = ?"
}

// GetSelectQuery Ritorna una possibile query tra quelle specificate
func (u *User) GetSelectQuery() (string, []interface{}) {

	querySQL := "SELECT * FROM " + TableUser + " WHERE " + ColUserRecordID + " = ?"

	return querySQL, []interface{}{&u.RecordID, &u.Name, &u.Lastname, &u.Gender}

}

// SetRecordID - Imposta il valore della chiave primaria
func (u *User) SetRecordID(id int) {
	u.RecordID = id
}
