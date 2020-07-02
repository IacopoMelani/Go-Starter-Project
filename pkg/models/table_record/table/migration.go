package table

import (
	"errors"
	"time"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db"
	record "github.com/IacopoMelani/Go-Starter-Project/pkg/models/table_record"
)

// Constants related to the migrations table
const (
	MigrationsColRecordID  = "record_id"
	MigrationsColCreatedAt = "created_at"
	MigrationsColName      = "name"
	MigrationsColStatus    = "status"

	MigrationsTableName = "migrations"

	// MigrationNotRun - Defines the constant for the state of a migration that has not yet been performed
	MigrationNotRun = 0
	// MigrationSuccess - Defines the constant for the state of a successful migration
	MigrationSuccess = 1
	// MigrationFailed - Defines the constant for the state of a migration that has been performed but has failed
	MigrationFailed = 2
)

// Migration - Struct that defines the migrations table, implements TableRecordInterface
type Migration struct {
	tr        *record.TableRecord
	RecordID  int64     `db:"record_id"`
	CreatedAt time.Time `db:"created_at"`
	Name      string    `db:"name"`
	Status    int       `db:"status"`
}

var dm = &Migration{}

// InsertNewMigration - It takes care of inserting a record in the migrations table
func InsertNewMigration(db db.SQLConnector, name string, status int) (*Migration, error) {

	if name == "" {
		return nil, errors.New("Empty migration's name")
	}

	m := NewMigration(db)
	m.Name = name
	m.Status = status
	m.CreatedAt = time.Now().UTC()
	if err := record.Save(m); err != nil {
		return nil, err
	}
	m.tr.SetIsNew(false).SetSQLConnection(db)

	return m, nil
}

// LoadAllMigrations - Load all instances of Migration from the database
func LoadAllMigrations(db db.SQLConnector) ([]*Migration, error) {

	query := "SELECT " + record.AllField(dm) + " FROM " + dm.GetTableName()

	rows, err := db.Queryx(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*Migration

	for rows.Next() {

		m := NewMigration(db)

		if err := record.LoadFromRow(rows, m); err != nil {
			return nil, err
		}

		result = append(result, m)
	}

	return result, nil
}

// LoadMigrationByName - It takes care of loading the instance of a record of the migrations table given the name
func LoadMigrationByName(name string, m *Migration) error {

	query := "SELECT " + record.AllField(m) + " FROM " + m.GetTableName() + " WHERE " + MigrationsColName + " = ?"

	return record.FetchSingleRow(m, query, name)
}

// NewMigration - It takes care of instantiating a new Migration object by instantiating its TableRecord and setting it as "new"
// We recommend that you always use this method to create a new Migration instance
func NewMigration(db db.SQLConnector) *Migration {

	m := new(Migration)
	m.tr = record.NewTableRecord(true, false)
	m.tr.SetSQLConnection(db)
	return m
}

// GetTableRecord - Returns the instance of TableRecord
func (m Migration) GetTableRecord() *record.TableRecord {
	return m.tr
}

// GetPrimaryKeyName - Returns the primary key name
func (m Migration) GetPrimaryKeyName() string {
	return MigrationsColRecordID
}

// GetPrimaryKeyValue - Returns the primary key value
func (m Migration) GetPrimaryKeyValue() int64 {
	return m.RecordID
}

// GetTableName - Returns the table name
func (m Migration) GetTableName() string {
	return MigrationsTableName
}
