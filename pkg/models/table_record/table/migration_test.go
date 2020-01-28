package table

import (
	"errors"
	"testing"

	"github.com/jmoiron/sqlx"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db"
	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db/transactions"

	"github.com/subosito/gotenv"
)

func TestMigration(t *testing.T) {

	if err := gotenv.Load("./../../../../.env"); err != nil {
		t.Fatal("Errore caricamento configurazione")
	}

	err := transactions.WithTransactionx(db.GetConnection().(*sqlx.DB), func(tx db.SQLConnector) error {

		mName := "test_migration"

		_, err := InsertNewMigration(tx, "", 1)
		if err == nil {
			t.Error("Dovrebbe essere errore")
		}

		migration, err := InsertNewMigration(tx, mName, 1)
		if err != nil {
			t.Fatal(err.Error())
		}

		if migration.GetPrimaryKeyValue() == 0 {
			t.Error("Chiave primaria non impostata")
		}

		m := NewMigration(tx)

		err = LoadMigrationByName(mName, m)
		if err != nil {
			t.Fatal(err.Error())
		}

		if m.Name != migration.Name || m.Status != migration.Status {
			t.Fatal("Operazione di migrazione errata")
		}

		allMigrations, err := LoadAllMigrations(tx)
		if err != nil {
			t.Fatal(err.Error())
		}

		if len(allMigrations) == 0 {
			t.Fatal("Migrazioni non caricate correttamente")
		}

		return errors.New("Rollback")
	})

	if err.Error() != "Rollback" {
		t.Fatal(err.Error())
	}
}
