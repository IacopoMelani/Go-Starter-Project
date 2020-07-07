package record

import (
	"errors"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db"
)

// executeSaveUpdateQuery - Si occupa di eseguire fisicamente la query, in caso di successo restituisce l'Id appena inserito
func executeSaveUpdateQuery(ti TableRecordInterface, query string, params []interface{}) error {

	conn := getTableRecordConnection(ti)

	switch ti.GetTableRecord().DriverName() {

	case db.DriverSQLServer:

		rows, err := conn.Queryx(query, params...)
		if err != nil {
			return err
		}

		if rows.Next() {

			if err := LoadFromRow(rows, ti); err != nil {
				return err
			}
		}

	case db.DriverMySQL:

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

// save - Si occupa di inserire un nuovo record nella tabella
func save(ti TableRecordInterface) error {

	query := genSaveQuery(ti)
	fValue := getFieldsValueNoPrimary(ti)
	err := executeSaveUpdateQuery(ti, query, fValue)
	if err != nil {
		return err
	}

	return nil
}

// update - Si occupa di aggiornare il record nel database
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

// All - Restituisce tutti i risultati per il costruttore del table record passato
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

// Delete - Si occupa di cancellare un record sul database
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

// ExecQuery - Esegue la query costruita con QueryBuilder
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

// FetchSingleRow -
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

// LoadByID - Carica l'istanza passata con i valori della sua tabella ricercando per chiave primaria
func LoadByID(ti TableRecordInterface, id interface{}) error {

	query := "SELECT " + AllField(ti) + " FROM " + ti.GetTableName() + " WHERE " + ti.GetPrimaryKeyName() + " = ?"

	params := []interface{}{interface{}(id)}

	return FetchSingleRow(ti, query, params...)
}

// Save - Si occupa di eseguire il salvataggio della TableRecord eseguendo un inserimento se TableRecord::isNew risulta false, altrimenti ne aggiorna il valore
func Save(ti TableRecordInterface) error {

	t := ti.GetTableRecord()

	if t.isReadOnly {
		return errors.New("Read-only model")
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
