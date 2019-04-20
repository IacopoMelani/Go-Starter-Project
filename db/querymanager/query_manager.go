package querymanager

import "Go-Starter-Project/db"

// Salvable - Interfaccia per permettere di generalizzare il salvataggio di un model sul database
type Salvable interface {
	GetSaveQuery() string
	SetRecordID(id int)
}

// Selecter - Interfaccia per permettere di generalizzare una select di un model
type Selecter interface {
	GetSelectQuery() (string, []interface{})
}

// Save - Metodo che si occcupa del salvataggio fisico sul database
func Save(s Salvable, params []interface{}) error {

	db := db.GetConnection()

	stmt, err := db.Prepare(s.GetSaveQuery())
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(params...)
	if err != nil {
		return err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	if lastID > 0 {
		s.SetRecordID(int(lastID))
	}
	return nil
}

// Select - Metodo che si occupa di eseguire una select sul database
func Select(s Selecter, query int, params ...interface{}) error {

	db := db.GetConnection()

	querySQL, scanFields := s.GetSelectQuery()

	stmt, err := db.Prepare(querySQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Query(params...)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(scanFields...)
		if err != nil {
			return err
		}
	}

	return nil
}
