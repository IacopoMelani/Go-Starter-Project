package record

import (
	"errors"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db/driver"
)

// executeSaveUpdateQuery - Execs the insert/update query and update the current TableRecordInterface passed
func executeSaveUpdateQuery(ti TableRecordInterface, query string, params []interface{}) error {

	conn := getTableRecordConnection(ti)

	switch ti.GetTableRecord().DriverName() {

	case driver.DriverSQLServer:

		rows, err := conn.Queryx(query, params...)
		if err != nil {
			return err
		}

		if rows.Next() {

			if err := LoadFromRow(rows, ti); err != nil {
				return err
			}
		}

	case driver.DriverMySQL:

		res, err := conn.Exec(query, params...)
		if err != nil {
			return err
		}

		lastID, err := res.LastInsertId()
		if err != nil {
			return err
		}

		if err := LoadByID(ti, lastID); err != nil {
			return err
		}
	}

	return nil
}

// save - Saves the model into the database
func save(ti TableRecordInterface) error {

	query := genSaveQuery(ti)
	fValue := getFieldsValueNoPrimary(ti)
	err := executeSaveUpdateQuery(ti, query, fValue)
	if err != nil {
		return err
	}

	return nil
}

// update - Updates the model into the database
func update(ti TableRecordInterface) error {

	query := genUpdateQuery(ti)
	fValue := getFieldsValueNoPrimary(ti)
	err := executeSaveUpdateQuery(ti, query, append(fValue, GetPrimaryKeyValue(ti)))
	if err != nil {
		return err
	}

	if err := LoadByID(ti, GetPrimaryKeyValue(ti)); err != nil {
		return err
	}

	return nil
}

// All - Returns all models from database, requires a func constructor that return a TableRecordInterface
func All(ntm NewTableModel) ([]TableRecordInterface, error) {

	var result []TableRecordInterface

	pivot := ntm()

	db := pivot.GetTableRecord().db

	query := "SELECT " + AllField(pivot) + " FROM " + pivot.GetTableName()

	rows, err := db.Queryx(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		ti := ntm()

		err = LoadFromRow(rows, ti)
		if err != nil {
			return nil, err
		}

		result = append(result, ti)
	}

	return result, nil
}

// Delete - Deletes the model from database
func Delete(ti TableRecordInterface) (int64, error) {

	db := getTableRecordConnection(ti)

	stmt, err := db.Prepare(genDeleteQuery(ti))
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(GetPrimaryKeyValue(ti))
	if err != nil {
		return 0, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return 0, nil
	}

	return rows, nil
}

// ExecQuery - Execs the query built with QueryBuilder
func ExecQuery(ti TableRecordInterface, ntm NewTableModel) ([]TableRecordInterface, error) {

	t := ti.GetTableRecord()

	stmt, err := t.PrepareStmt(ti.GetTableName())
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Queryx(t.Params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tiList []TableRecordInterface

	for rows.Next() {

		nti := ntm()

		if err := LoadFromRow(rows, nti); err != nil {
			return nil, err
		}

		tiList = append(tiList, nti)
	}

	ti.GetTableRecord().ResetStmt()

	return tiList, nil
}

// FetchSingleRow - Execs the query and retrieves only the first result
func FetchSingleRow(tri TableRecordInterface, query string, params ...interface{}) error {

	db := getTableRecordConnection(tri)

	rows, err := db.Queryx(query, params...)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {

		if err := LoadFromRow(rows, tri); err != nil {
			return err
		}
	}

	return nil
}

// LoadByID -  Loads the model passed by primary key value
func LoadByID(ti TableRecordInterface, id interface{}) error {

	query := "SELECT " + AllField(ti) + " FROM " + ti.GetTableName() + " WHERE " + ti.GetPrimaryKeyName() + " = ?"

	params := []interface{}{interface{}(id)}

	return FetchSingleRow(ti, query, params...)
}

// Save -  Saves the model to the database, if the model is "new" inserts a new one otherwise update the record
func Save(ti TableRecordInterface) error {

	t := ti.GetTableRecord()

	if t.isReadOnly {
		return errors.New("read-only model")
	}

	if t.isNew {

		err := save(ti)
		if err != nil {
			return err
		}

	} else {

		err := update(ti)
		if err != nil {
			return err
		}
	}

	return nil
}
