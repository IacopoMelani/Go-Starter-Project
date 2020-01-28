package transactions

import (
	"errors"
	"testing"

	"github.com/subosito/gotenv"

	"github.com/jmoiron/sqlx"

	record "github.com/IacopoMelani/Go-Starter-Project/pkg/models/table_record"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/models/table_record/table"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db"
)

const testMigrationName = "FOFFO"

func TestTransactionx(t *testing.T) {

	if err := gotenv.Load("./../../../../.env"); err != nil {
		t.Fatal("Errore caricamento configurazione")
	}

	WithTransactionx(db.GetConnection().(*sqlx.DB), func(tx db.SQLConnector) error {

		_, err := table.LoadAllMigrations(tx)
		if err != nil {
			return err
		}

		return nil
	})

	err := WithTransactionx(db.GetConnection().(*sqlx.DB), func(tx db.SQLConnector) error {

		newMigration := table.NewMigration(tx)

		newMigration.Name = testMigrationName

		if err := record.Save(newMigration); err != nil {
			return err
		}

		migration := table.NewMigration(tx)
		if err := table.LoadMigrationByName(testMigrationName, migration); err != nil {
			return err
		}

		if migration.Name != testMigrationName {
			t.Error("Fallimento caricamento migrazione")
		}

		return errors.New("Rollback")
	})

	if err == nil {
		t.Error("err dovrebbe essere valorizzato")
	}
}
