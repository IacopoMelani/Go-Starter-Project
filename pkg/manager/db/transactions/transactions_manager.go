package transactions

import (
	"database/sql"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db"
	"github.com/jmoiron/sqlx"
)

// Refs - https://pseudomuto.com/2018/01/clean-sql-transactions-in-golang/

// TxFn - Funzione eseguita per durante la transaction
type TxFn func(db.SQLConnector) error

// WithTransactionx - Crea una nuova transaction eseguendo l'handle, si interfaccia con la libreria "github.com/jmoiron/sqlx"
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

// PipelineStmt - Defines a wrapper for sql stmt
type PipelineStmt struct {
	query string
	args  []interface{}
}

func NewPipelineStmt(query string, args ...interface{}) *PipelineStmt {
	return &PipelineStmt{query, args}
}

func (ps *PipelineStmt) Exec(conn db.SQLConnector) (sql.Result, error) {
	return conn.Exec(ps.query, ps.args...)
}

func RunPipelineStmts(conn db.SQLConnector, stmts ...*PipelineStmt) error {

	for _, ps := range stmts {

		_, err := ps.Exec(conn)
		if err != nil {
			return err
		}
	}

	return nil
}

func RunPipelineStmtsWithTransactionx(conn db.SQLConnector, stmts ...*PipelineStmt) error {
	return WithTransactionx(conn.(*sqlx.DB), func(tx db.SQLConnector) error {
		return RunPipelineStmts(tx, stmts...)
	})
}

type Pipeline interface {

}