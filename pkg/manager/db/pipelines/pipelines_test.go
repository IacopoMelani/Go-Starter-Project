package pipelines

import (
	"os"
	"testing"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db"

	"github.com/subosito/gotenv"
)

func getDb() db.SQLConnector {
	loadEnv()
	db.InitConnection(os.Getenv("SQL_DRIVER"), os.Getenv("STRING_CONNECTION"))
	return db.GetConnection()
}

func getInsertParams() []interface{} {
	return []interface{}{"pipe", "line", "P"}
}

func getInsertQuery() string {
	return `
		INSERT INTO users(name, lastname, gender)
		VALUES(?, ?, ?)
	`
}

func loadEnv() {
	if err := gotenv.Load("./../../../../.env"); err != nil {
		panic("Errore caricamento configurazione")
	}
}

func TestPipeline(t *testing.T) {

	conn := getDb()

	if !db.TableExists("users") {
		t.Error("Error: table users not exists")
	}

	var pipelineStmt []*PipelineStmt

	for range make([]int, 10) {
		pipelineStmt = append(pipelineStmt, NewPipelineStmt(getInsertQuery(), getInsertParams()...))
	}

	if err := RunPipelineStmtsWithTransactionx(conn, pipelineStmt...); err != nil {
		t.Error(err)
	}

	pm := NewPipelineManager(conn)

	for _, p := range pipelineStmt {
		pm.AddPipe(p)
	}

	if err := pm.RunPipelinesWithTransactionx(); err != nil {
		t.Error(err)
	}

	var pipelines []Pipeline
	for _, p := range pipelineStmt {
		pipelines = append(pipelines, p)
	}

	if err := RunpipelinesWithTransactionx(conn, pipelines...); err != nil {
		t.Error(err)
	}

	pipelineStmtErr := NewPipelineStmt("ERR SQL", 1)

	if err := RunpipelinesWithTransactionx(conn, pipelineStmtErr); err == nil {
		t.Error("Error has to occur")
	}
}
