package migration

import (
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

// GetMigratorInstance - Restituisce l'unica istanza di migrator
func GetMigratorInstance() *Migrator {

	onceMigrator.Do(func() {

		migrator.migrationsList = migrationsList
	})

	return migrator
}

// DoUpMigrations -
func (m *Migrator) DoUpMigrations() error {

	db, err := db.GetConnection().Begin()
	if err != nil {
		return err
	}

	for _, mi := range m.migrationsList {

		migration := table.NewMigration()

		err = table.LoadMigrationByName(mi.GetMigrationName(), migration)
		if err != nil {
			db.Rollback()
			return err
		}

		if migration.GetTableRecord().RecordID != 0 && migration.Status == 1 {
			continue
		}

		migration.Name = mi.GetMigrationName()
		migration.Status = 0
		err = record.Save(migration)
		if err != nil {
			db.Rollback()
			return err
		}

		_, err = db.Exec(mi.Up())
		if err != nil {
			db.Rollback()
			return err
		}

		migration.Status = 1
		err = record.Save(migration)
		if err != nil {
			db.Rollback()
			return nil
		}
	}

	db.Commit()

	return nil
}
