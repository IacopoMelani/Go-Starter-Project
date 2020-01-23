package transactions

import (
	"github.com/IacopoMelani/Go-Starter-Project/pkg/db"
	"github.com/jmoiron/sqlx"
)

// Refs - https://pseudomuto.com/2018/01/clean-sql-transactions-in-golang/

// TxFn - Funzione eseguita per durante la transaction
type TxFn func(db.SQLConnector) error

// WithTransactionx - Crea una nuova transaction eseguendo l'handle, si una con libreria	"github.com/jmoiron/sqlx"
func WithTransactionx(db *sqlx.DB, fn TxFn) (err error) {

	tx, err := db.Beginx()
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	err = fn(tx)
	return err
}
