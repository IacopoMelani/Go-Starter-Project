package migration

import (
	"database/sql"
	"sync"

	record "github.com/IacopoMelani/Go-Starter-Project/models/table_record"

	"github.com/IacopoMelani/Go-Starter-Project/models/table_record/table"

	"github.com/IacopoMelani/Go-Starter-Project/db"
)

// Migrable - Definisce l'interfaccia per poter avviare una migrazione sul database
type Migrable interface {
	Down() string
	GetMigrationName() string
	Up() string
}

// Migrator - Struct che si occupa di effettuare le migrazione sul database
type Migrator struct {
	migrationsList []Migrable
}

var (
	migrator     *Migrator
	onceMigrator sync.Once
)

// createMigrationsTable - Si occupa di creare la tabella delle migrazioni
func createMigrationsTable(conn *sql.Tx) error {

	query := `CREATE TABLE IF NOT EXISTS migrations (
    record_id INT AUTO_INCREMENT,
    created_at DATETIME NOT NULL,
    name VARCHAR(255) NOT NULL,
    status INT NOT NULL,
	PRIMARY KEY (record_id)
	)`

	_, err := conn.Exec(query)

	return err
}

// GetMigratorInstance - Restituisce l'unica istanza di migrator
func GetMigratorInstance() *Migrator {

	onceMigrator.Do(func() {
		migrator = new(Migrator)
		migrator.migrationsList = migrationsList
	})

	return migrator
}

// DoUpMigrations -
func (m *Migrator) DoUpMigrations() error {

	conn, _ := db.GetConnection().Begin()

	exist, err := db.TableExists(table.MigrationsTableName)
	if err != nil {
		conn.Rollback()
		return err
	}

	if !exist {

		if err = createMigrationsTable(conn); err != nil {
			conn.Rollback()
			return err
		}
	}

	for _, mi := range m.migrationsList {

		migration := table.NewMigration()

		err = table.LoadMigrationByName(mi.GetMigrationName(), migration)
		if err != nil {
			conn.Rollback()
			return err
		}

		if migration.GetTableRecord().RecordID != 0 && migration.Status == 1 {
			continue
		}

		migration, err := table.InsertNewMigration(mi.GetMigrationName(), 0)
		if err != nil {
			conn.Rollback()
			return err
		}

		_, err = conn.Exec(mi.Up())
		if err != nil {
			conn.Rollback()
			return err
		}

		migration.Status = 1
		err = record.Save(migration)
		if err != nil {
			conn.Rollback()
			return nil
		}
	}

	conn.Commit()

	return nil
}
