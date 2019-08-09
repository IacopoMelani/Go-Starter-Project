package builder

import (
	"Go-Starter-Project/db"
	"database/sql"
	"testing"

	"github.com/subosito/gotenv"
)

type testStruct struct {
	recordID int64
	name     sql.NullString
	lastname sql.NullString
	gender   sql.NullBool
	Builder
}

func TestQueryBuilder(t *testing.T) {

	gotenv.Load("./../../.env")
	db := db.GetConnection()

	ts := new(testStruct)

	_, err := db.Exec("INSERT INTO users(name, lastname, gender) VALUES(?, ?, ?)", []interface{}{"Mario", "Rossi", nil}...)
	if err != nil {
		t.Fatalf("Errore durante salvataggio record da ricercare, errore %s", err.Error())
	}

	ts.SelectField("*").WhereEqual("name", "Mario").WhereNull("gender", true).WhereOperator("record_id", ">=", 1).OrderByAsc("record_id").OrderByDesc("name")

	query := ts.BuildQuery("users")

	stmt, err := db.Prepare(query)
	if err != nil {
		t.Fatalf("Errore durante la costruzione della query: %s", query)
	}
	defer stmt.Close()

	rows, err := stmt.Query(ts.Params...)
	if err != nil {
		t.Fatalf("Errore durante l'esecuzione della query: %s parametri %v", query, ts.Params)
	}
	defer rows.Close()

	if rows.Next() {

		err := rows.Scan(&ts.recordID, &ts.name, &ts.lastname, &ts.gender)
		if err != nil {
			t.Fatalf("Errore durante prelievo valori, errore restituito: %s", err.Error())
		}
	}

	if ts.name.String != "Mario" {
		t.Fatal("Errore: prelevata riga errata")
	}

	ts.ResetStmt()
}
