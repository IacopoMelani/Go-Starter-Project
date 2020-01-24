package builder

import (
	"database/sql"
	"testing"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db"
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

	if err := gotenv.Load("./../../../../.env"); err != nil {
		t.Fatal("Errore caricamento configurazione")
	}

	db := db.GetConnection()

	ts := new(testStruct)

	_, err := db.Exec("INSERT INTO users(name, lastname, gender) VALUES(?, ?, ?)", []interface{}{"Mario", "Rossi", nil}...)
	if err != nil {
		t.Fatalf("Errore durante salvataggio record da ricercare, errore %s", err.Error())
	}

	ts.SelectField("*").WhereEqual("name", "Mario").WhereEqual("lastname", "Rossi").WhereNull("gender", true).WhereOperator("record_id", ">=", 1).WhereOperator("record_id", "<", "99999").OrderByAsc("record_id").OrderByDesc("name")

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

	query = ts.BuildQuery("users")
	stmt, err = db.Prepare(query)
	if err != nil {
		t.Fatalf("Errore durante la costruzione della query: %s", query)
	}
	defer stmt.Close()

	ts.ResetStmt()

	ts.SelectField("name").SelectField("lastname").GroupByField("name").GroupByField("lastname")

	query = ts.BuildQuery("users")

	stmt, err = db.Prepare(query)
	if err != nil {
		t.Fatalf("Errore durante la query %s parametri %v", query, ts.Params)
	}
	defer stmt.Close()

	ts.ResetStmt()

	ts.WhereOperator("name", "=", "Mario").WhereNull("gender", false)

	query = ts.BuildQuery("users")

	stmt, err = db.Prepare(query)
	if err != nil {
		t.Fatalf("Errore durante la query %s parametri %v", query, ts.Params)
	}
	defer stmt.Close()

	ts.ResetStmt()

	ts.WhereNull("gender", false)

	query = ts.BuildQuery("users")

	stmt, err = db.Prepare(query)
	if err != nil {
		t.Fatalf("Errore durante la query %s parametri %v", query, ts.Params)
	}
	defer stmt.Close()

	ts.ResetStmt()

	ts.WhereNull("gender", true)

	query = ts.BuildQuery("users")

	stmt, err = db.Prepare(query)
	if err != nil {
		t.Fatalf("Errore durante la query %s parametri %v", query, ts.Params)
	}
	defer stmt.Close()
}
