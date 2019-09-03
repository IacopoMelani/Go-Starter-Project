package table

import (
	"errors"
	"time"

	"github.com/IacopoMelani/Go-Starter-Project/db"

	record "github.com/IacopoMelani/Go-Starter-Project/models/table_record"
)

// Costanti relative alla tabella migrations
const (
	MigrationsColRecordID  = "record_id"
	MigrationsColCreatedAt = "created_at"
	MigrationsColName      = "name"
	MigrationsColStatus    = "status"

	MigrationsTableName = "migrations"

	// MigrationNotRun - Definisce la costante per lo stato di una migrazione che non è stata ancora eseguita
	MigrationNotRun = 0
	// MigrationSuccess - Definisce la costante per lo stato di una migrazione che è stata eseguita con successo
	MigrationSuccess = 1
	// MigrationFailed - Definisce la costante per lo stato di una migrazione che è stata eseguita ma è fallita
	MigrationFailed = 2
)

// Migration - Struct che definisce la tabella migrations
// implementa TableRecordInterface
type Migration struct {
	tr        *record.TableRecord
	CreatedAt time.Time `db:"created_at"`
	Name      string    `db:"name"`
	Status    int       `db:"status"`
}

// InsertNewMigration - Si occupa di inserire un record nella tabella migrations
func InsertNewMigration(name string, status int) (*Migration, error) {

	if name == "" {
		return nil, errors.New("Empty migration's name")
	}

	m := NewMigration()
	m.Name = name
	m.Status = status
	m.CreatedAt = time.Now().UTC()
	err := record.Save(m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// LoadMigrationByName - Si occupa di caricare l'istanza di un record della tabella migrations dato il nome
func LoadMigrationByName(name string, m *Migration) error {

	db := db.GetConnection()

	query := "SELECT " + record.AllField(m) + " FROM " + m.GetTableName() + " WHERE " + MigrationsColName + " = ?"

	rows, err := db.Query(query, name)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {

		_, vField := record.GetFieldMapper(m)

		dest := append([]interface{}{&m.tr.RecordID}, vField...)

		err := rows.Scan(dest)
		if err != nil {
			return err
		}
	}

	return nil
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
	return MigrationsColRecordID
}

// GetTableName - Restituisce il nome della tabella
func (m Migration) GetTableName() string {
	return MigrationsTableName
}

// New - Si occupa di istanziare una nuova struct andando ad istaziare table record e settanto il campo isNew a true
func (m Migration) New() record.TableRecordInterface {
	return NewMigration()
}
