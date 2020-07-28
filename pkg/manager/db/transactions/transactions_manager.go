package transactions

import (
	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db"
	"github.com/jmoiron/sqlx"
)

// Refs - https://pseudomuto.com/2018/01/clean-sql-transactions-in-golang/

// TxFn - Func type to executes during transaction
type TxFn func(db.SQLConnector) error

// WithTransactionx - Creates new transaction instance and executes the handler
func WithTransactionx(db *sqlx.DB, fn TxFn) (err error) {

	tx, err := db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	err = fn(tx)
	return err
}
