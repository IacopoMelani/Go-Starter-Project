package pipelines

import (
	"database/sql"
	"sync"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db"
	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db/transactions"
	"github.com/jmoiron/sqlx"
)

// Pipeline - Defines generics interface to implemtns a pipeline
type Pipeline interface {
	Exec(conn db.SQLConnector) (sql.Result, error)
}

// RunPipelines - Takes a slice of Pipeline and execs all them
func RunPipelines(conn db.SQLConnector, pipelines ...Pipeline) error {

	for _, ps := range pipelines {

		_, err := ps.Exec(conn)
		if err != nil {
			return err
		}
	}

	return nil
}

// RunpipelinesWithTransactionx - Takes a slice of Pipeline and execs all them under transaction
func RunpipelinesWithTransactionx(conn db.SQLConnector, pipelines ...Pipeline) error {
	return transactions.WithTransactionx(conn.(*sqlx.DB), func(tx db.SQLConnector) error {
		return RunPipelines(tx, pipelines...)
	})
}

// PipelineManager - Defines a manager for executing pipeline
type PipelineManager struct {
	conn      db.SQLConnector
	mu        sync.Mutex
	pipelines []Pipeline
}

// NewPipelineManager - Returns a new instance ptr of PipelineManager
func NewPipelineManager(conn db.SQLConnector) *PipelineManager {
	pm := new(PipelineManager)
	pm.conn = conn
	return pm
}

// AddPipe - Appends the single pipe to the queue
func (pm *PipelineManager) AddPipe(p Pipeline) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	pm.pipelines = append(pm.pipelines, p)
}

// RunPipelines - Runs all added Pipeline interfaces
func (pm *PipelineManager) RunPipelines() error {
	pm.mu.Lock()
	defer func() {
		pm.pipelines = pm.pipelines[:0]
		pm.mu.Unlock()
	}()
	return RunPipelines(pm.conn, pm.pipelines...)
}

// RunPipelinesWithTransactionx - Runs all added Pipeline interfaces under transaction
func (pm *PipelineManager) RunPipelinesWithTransactionx() error {
	return transactions.WithTransactionx(pm.conn.(*sqlx.DB), func(tx db.SQLConnector) error {
		return pm.RunPipelines()
	})
}

// PipelineStmt - Defines a wrapper for sql stmt
type PipelineStmt struct {
	query string
	args  []interface{}
}

// NewPipelineStmt - Returns instance ptr of PipelineStmt
func NewPipelineStmt(query string, args ...interface{}) *PipelineStmt {
	return &PipelineStmt{query, args}
}

// RunPipelineStmts - Runs a slice of *PipelineStmt
func RunPipelineStmts(conn db.SQLConnector, stmts ...*PipelineStmt) error {

	var pipelines []Pipeline
	for _, ps := range stmts {
		pipelines = append(pipelines, ps)
	}

	return RunPipelines(conn, pipelines...)
}

// RunPipelineStmtsWithTransactionx - Runs a slice of *PipelineStmt under transaction
func RunPipelineStmtsWithTransactionx(conn db.SQLConnector, stmts ...*PipelineStmt) error {
	return transactions.WithTransactionx(conn.(*sqlx.DB), func(tx db.SQLConnector) error {
		return RunPipelineStmts(tx, stmts...)
	})
}

// Exec - Exec a stmt, necessary to implemnt Pipeline
func (ps *PipelineStmt) Exec(conn db.SQLConnector) (sql.Result, error) {
	return conn.Exec(ps.query, ps.args...)
}
