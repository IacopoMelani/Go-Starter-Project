package table

import (
	"time"

	record "github.com/IacopoMelani/Go-Starter-Project/models/table_record"
)

const (
	migrationNotRun  = 0
	migrationSuccess = 1
	migrationFailed  = 2
)

// Migration - Struct che definisce la tabella migrations
// implementa TableRecordInterface
type Migration struct {
	tr        *record.TableRecord
	CreatedAt time.Time `db:"created_at"`
	Name      string    `db:"name"`
	Status    int       `db:"status"`
}

// NewMigration - Si occupa di istanziare un nuovo oggetto Migration istanziando il relativo TableRecord e impostandolo come "nuovo"
// È consigliato utilizzare sempre questo metodo per creare una nuova istanza di Migration
func NewMigration() *Migration {

	m := new(Migration)
	m.tr = new(record.TableRecord)
	m.tr.SetIsNew(true)

	return m
}

// GetTableRecord - Restituisce l'istanza di TableRecord
func (m Migration) GetTableRecord() *record.TableRecord {
	return m.tr
}

// GetPrimaryKeyName - Restituisce il nome della chiave primaria
func (m Migration) GetPrimaryKeyName() string {
	return "record_id"
}

// GetTableName - Restituisce il nome della tabella
func (m Migration) GetTableName() string {
	return "migrations"
}

// New - Si occupa di istanziare una nuova struct andando ad istaziare table record e settanto il campo isNew a true
func (m Migration) New() *Migration {
	return NewMigration()
}
