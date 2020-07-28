package migration

import (
	"os"
	"testing"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/models/table_record/table"

	record "github.com/IacopoMelani/Go-Starter-Project/pkg/models/table_record"

	"github.com/subosito/gotenv"
)

// TestTable - Example
type TestTable struct{}

// GetMigrationName - Example
func (t TestTable) GetMigrationName() string {
	return "create_test_table"
}

// Down - Example
func (t TestTable) Down() string {

	var query string

	switch db.DriverName() {

	case db.DriverMySQL:
		query = "DROP TABLE IF EXISTS test"
	case db.DriverSQLServer:
		query = `
		IF EXISTS (SELECT * FROM sysobjects WHERE name='test' and xtype='U')
		DROP TABLE test`
	}

	return query
}

// Up - Example
func (t TestTable) Up() string {

	var query string

	switch db.DriverName() {

	case db.DriverMySQL:
		query = `CREATE TABLE IF NOT EXISTS test (
		record_id INT AUTO_INCREMENT,
		PRIMARY KEY (record_id)
		)`

	case db.DriverSQLServer:
		query = `
		IF NOT EXISTS (SELECT * FROM sysobjects WHERE name='test' and xtype='U')
		CREATE TABLE test (
			record_id BIGINT IDENTITY(1, 1) NOT NULL PRIMARY KEY
		)`
	}

	return query

}

func TestMigrationManager(t *testing.T) {

	if err := gotenv.Load("./../../../.env"); err != nil {
		t.Fatal("Errore caricamento configurazione")
	}
	db.InitConnection(os.Getenv("SQL_DRIVER"), os.Getenv("STRING_CONNECTION"))

	var migrationsList = []Migrable{
		TestTable{},
	}

	InitMigrationsList(migrationsList)

	var query string

	tx := db.GetConnection()

	switch tx.DriverName() {

	case db.DriverMySQL:
		query = "DROP TABLE IF EXISTS migrations"

	case db.DriverSQLServer:
		query = `
			IF EXISTS (SELECT * FROM sysobjects WHERE name='migrations' and xtype='U')
			DROP TABLE migrations
			`
	}

	_, err := tx.Exec(query)
	if err != nil {
		t.Fatal(err.Error())
	}

	migrator := GetMigratorInstance()

	err = migrator.DoDownMigrations()
	if err != nil {
		t.Fatal(err.Error())
	}

	err = DoUpMigrations()
	if err != nil {
		t.Fatal(err.Error())
	}

	err = DoDownMigrations()
	if err != nil {
		t.Fatal(err.Error())
	}

	err = DoUpMigrations()
	if err != nil {
		t.Fatal(err.Error())
	}

	migration := table.NewMigration(tx)

	err = table.LoadMigrationByName("create_test_table", migration)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = record.Delete(migration)
	if err != nil {
		t.Fatal(err.Error())
	}
}
