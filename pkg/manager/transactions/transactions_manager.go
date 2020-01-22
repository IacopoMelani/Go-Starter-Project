package transactions

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// Refs - https://pseudomuto.com/2018/01/clean-sql-transactions-in-golang/

// Transaction - Interfaccia per gestire operazioni sotto transaction
type Transaction interface {
	Exec(string, ...interface{}) (sql.Result, error)
	Prepare(string) (*sql.Stmt, error)
	Query(string, ...interface{}) (*sql.Rows, error)
	QueryRow(string, ...interface{}) *sql.Row
}

// TxFn - Funzione eseguita per durante la transaction
type TxFn func(Transaction) error

// WithTransaction - Crea una nuova transaction eseguendo l'handle, si una con libreria "database/sql"
func WithTransaction(db *sql.DB, fn TxFn) (err error) {

	tx, err := db.Begin()
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

// WithTransactionx - Crea una nuova transaction eseguendo l'handle, si una con libreria	"github.com/jmoiron/sqlx"
func WithTransactionx(db *sqlx.DB, fn TxFn) (err error) {

	tx, err := db.Begin()
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
