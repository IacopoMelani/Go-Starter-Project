package transactions

import (
	"errors"
	"testing"

	"github.com/subosito/gotenv"

	"github.com/jmoiron/sqlx"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db"
)

const testMigrationName = "FOFFO"

func TestTransactionx(t *testing.T) {

	if err := gotenv.Load("./../../../../.env"); err != nil {
		t.Fatal("Errore caricamento configurazione")
	}

	err := WithTransactionx(db.GetConnection().(*sqlx.DB), func(tx db.SQLConnector) error {
		return errors.New("Rollback")
	})
	if err == nil {
		t.Error("err dovrebbe essere valorizzato")
	}

	err = WithTransactionx(db.GetConnection().(*sqlx.DB), func(tx db.SQLConnector) error {
		return nil
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	err = WithTransactionx(db.GetConnection().(*sqlx.DB), func(tx db.SQLConnector) error {
		panic("panic")
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}
