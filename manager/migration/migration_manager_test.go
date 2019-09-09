package migration

import (
	"testing"

	record "github.com/IacopoMelani/Go-Starter-Project/models/table_record"

	"github.com/IacopoMelani/Go-Starter-Project/models/table_record/table"

	"github.com/IacopoMelani/Go-Starter-Project/db"
	"github.com/subosito/gotenv"
)

// TestTable - Definisce la strut per la generazione della tabella test
type TestTable struct{}

// GetMigrationName - Restituisce il nome della migrazione
func (t TestTable) GetMigrationName() string {
	return "create_test_table"
}

// Down - Definisce la query di migrazione down
func (t TestTable) Down() string { return "DROP TABLE IF EXISTS test" }

// Up - Definisce la query di migrazione up
func (t TestTable) Up() string {
	return `CREATE TABLE IF NOT EXISTS test (
    record_id INT AUTO_INCREMENT,
    PRIMARY KEY (record_id)
	)`
}

func TestMigrationManager(t *testing.T) {

	gotenv.Load("./../../.env")

	db := db.GetConnection()

	conn, err := db.Begin()
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = db.Exec("DROP TABLE IF EXISTS migrations")
	if err != nil {
		conn.Rollback()
		t.Fatal(err.Error())
	}

	migrationsList = make([]Migrable, 0)

	migrationsList = append(migrationsList, TestTable{})

	migrator := GetMigratorInstance()

	err = migrator.DoDownMigrations()
	if err != nil {
		conn.Rollback()
		t.Fatal(err.Error())
	}

	err = migrator.DoUpMigrations()
	if err != nil {
		conn.Rollback()
		t.Fatal(err.Error())
	}

	err = migrator.DoUpMigrations()
	if err != nil {
		conn.Rollback()
		t.Fatal(err.Error())
	}

	err = migrator.DoDownMigrations()
	if err != nil {
		conn.Rollback()
		t.Fatal(err.Error())
	}

	migration := table.NewMigration()

	err = table.LoadMigrationByName("create_test_table", migration)
	if err != nil {
		conn.Rollback()
		t.Fatal(err.Error())
	}

	_, err = record.Delete(migration)
	if err != nil {
		conn.Rollback()
		t.Fatal(err.Error())
	}

	conn.Commit()
}
