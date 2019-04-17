package models

import (
	"strings"
	"testDB/db"
)

// TableMirrorInterface -
type TableMirrorInterface interface {
	GetTableName() string
	GetPrimaryKeyName() string
	Save() (int, error)
}

// TableMirror - definisce una generica struct relativa a un model
type TableMirror struct {
	isNew  bool
	fields map[string]interface{}
}

func insertNewRow(t *TableMirror, ti TableMirrorInterface) (int, error) {

	db := db.GetConnection()

	queryParams, queryParamsName, queryPreparedParams := prepareQueryParams(*t, true)

	querySQL := "INSERT INTO " + ti.GetTableName() + "(" + strings.Join(queryParamsName, ",") + ") VALUES (" + strings.Join(queryPreparedParams, ",") + ")"

	stmt, err := db.Prepare(querySQL)
	if err != nil {
		return -1, err
	}

	res, err := stmt.Exec(queryParams...)
	if err != nil {
		return -1, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return int(lastID), nil
}

func prepareQueryParams(t TableMirror, insert bool) ([]interface{}, []string, []string) {
	var queryParams []interface{}
	var queryParamsName []string
	var queryPreparedParams []string

	for key, value := range t.fields {
		queryParams = append(queryParams, value)

		if insert {
			queryParamsName = append(queryParamsName, key)
		} else {
			queryParamsName = append(queryParamsName, key+" = ?")
		}

		queryPreparedParams = append(queryPreparedParams, "?")
	}

	return queryParams, queryParamsName, queryPreparedParams
}

func updateRow(t *TableMirror, ti TableMirrorInterface) (int, error) {

	db := db.GetConnection()

	queryParams, queryParamsName, _ := prepareQueryParams(*t, false)

	querySQL := "UPDATE " + ti.GetTableName() + " SET " + strings.Join(queryParamsName, ",") + " WHERE " + ti.GetPrimaryKeyName() + " = ? "

	stmt, err := db.Prepare(querySQL)
	if err != nil {
		return -1, err
	}

	queryParams = append(queryParams, t.fields[ti.GetPrimaryKeyName()])

	res, err := stmt.Exec(queryParams...)
	if err != nil {
		return -1, nil
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return int(lastID), nil
}

// Save -
func Save(t *TableMirror, ti TableMirrorInterface) (int, error) {
	if t.isNew {
		id, err := insertNewRow(t, ti)
		if err != nil {
			return -1, err
		}
		t.fields[ti.GetPrimaryKeyName()] = id
		return id, nil
	}
	id, err := updateRow(t, ti)
	if err != nil {
		return -1, err
	}
	return id, nil
}

// GetField - Restituisce il valore leggendolo da fields
func (m TableMirror) GetField(name string) interface{} {
	return m.fields[name]
}

// SetField - Inserisce il campo all'interno di fields
func (m *TableMirror) SetField(name string, value interface{}) *TableMirror {
	if m.fields == nil {
		m.fields = make(map[string]interface{})
	}
	m.fields[name] = value
	return m
}

// SetIsNew - Imposta se il record è nuovo o già esistente
func (m *TableMirror) SetIsNew(state bool) *TableMirror {
	m.isNew = state
	return m
}
