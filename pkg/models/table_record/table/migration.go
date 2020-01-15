package table

import (
	"errors"
	"time"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/db"
	record "github.com/IacopoMelani/Go-Starter-Project/pkg/models/table_record"
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

var dm = &Migration{}

// InsertNewMigration - Si occupa di inserire un record nella tabella migrations
func InsertNewMigration(name string, status int) (*Migration, error) {

	if name == "" {
		return nil, errors.New("Empty migration's name")
	}

	m := NewMigration()
	m.Name = name
	m.Status = status
	m.CreatedAt = time.Now().UTC()
	if err := record.Save(m); err != nil {
		return nil, err
	}
	m.tr.SetIsNew(false)

	return m, nil
}

// LoadAllMigrations - Carica tutte le istanze di Migration dal database
func LoadAllMigrations() ([]*Migration, error) {

	query := "SELECT " + record.AllField(dm) + " FROM " + dm.GetTableName()

	rows := db.QueryOrPanic(query)
	defer rows.Close()

	var result []*Migration

	for rows.Next() {

		m := NewMigration()

		if err := record.LoadFromRow(rows, m); err != nil {
			return nil, err
		}

		result = append(result, m)
	}

	return result, nil
}

// LoadMigrationByName - Si occupa di caricare l'istanza di un record della tabella migrations dato il nome
func LoadMigrationByName(name string, m *Migration) error {

	query := "SELECT " + record.AllField(m) + " FROM " + m.GetTableName() + " WHERE " + MigrationsColName + " = ?"

	rows := db.QueryOrPanic(query, name)
	defer rows.Close()

	if rows.Next() {

		if err := record.LoadFromRow(rows, m); err != nil {
			return err
		}
	}

	return nil
}

// NewMigration - Si occupa di istanziare un nuovo oggetto Migration istanziando il relativo TableRecord e impostandolo come "nuovo"
// È consigliato utilizzare sempre questo metodo per creare una nuova istanza di Migration
func NewMigration() *Migration {

	m := new(Migration)
	m.tr = record.NewTableRecord(true, false)

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
