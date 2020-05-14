package transactions

import (
	"errors"
	"os"
	"testing"

	"github.com/subosito/gotenv"

	"github.com/jmoiron/sqlx"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db"
)

func TestTransactionx(t *testing.T) {

	if err := gotenv.Load("./../../../../.env"); err != nil {
		t.Fatal("Errore caricamento configurazione")
	}
	db.InitConnection(os.Getenv("SQL_DRIVER"), os.Getenv("STRING_CONNECTION"))

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
