package migration

import (
	"os"
	"testing"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db"
	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db/driver"

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
	return driver.GetDropTableQuery(db.DriverName(), "test")
}

// Up - Example
func (t TestTable) Up() string {
	return driver.GetCreateEmptyTableQuery(db.DriverName(), "test")
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

	tx := db.GetConnection()

	query := driver.GetDropTableQuery(tx.DriverName(), "migrations")
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
