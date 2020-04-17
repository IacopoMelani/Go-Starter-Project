package migration

import (
	"errors"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db"
	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db/transactions"

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
func (t TestTable) Down() string { return "DROP TABLE IF EXISTS test" }

// Up - Example
func (t TestTable) Up() string {
	return `CREATE TABLE IF NOT EXISTS test (
    record_id INT AUTO_INCREMENT,
    PRIMARY KEY (record_id)
	)`
}

func TestMigrationManager(t *testing.T) {

	if err := gotenv.Load("./../../../.env"); err != nil {
		t.Fatal("Errore caricamento configurazione")
	}
	db.InitConnection("mysql", os.Getenv("STRING_CONNECTION"))

	err := transactions.WithTransactionx(db.GetConnection().(*sqlx.DB), func(tx db.SQLConnector) error {

		var migrationsList = []Migrable{
			TestTable{},
		}

		InitMigrationsList(migrationsList)

		_, err := tx.Exec("DROP TABLE IF EXISTS migrations")
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

		return errors.New("Rollback")
	})

	if err.Error() != "Rollback" {
		t.Fatal(err.Error())
	}
}
