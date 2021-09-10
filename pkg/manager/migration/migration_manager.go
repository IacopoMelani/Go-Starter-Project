package migration

import (
	"sync"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db"
	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db/driver"
	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db/transactions"
	"github.com/IacopoMelani/Go-Starter-Project/pkg/models/table_record/table"

	record "github.com/IacopoMelani/Go-Starter-Project/pkg/models/table_record"
)

// Migrable - Defines the generic interface to manage a migration
type Migrable interface {
	Down() string
	GetMigrationName() string
	Up() string
}

// Migrator - Manages the migration
type Migrator struct {
	migrationsList []Migrable
}

var (
	migrator           *Migrator
	onceMigrator       sync.Once
	migrationsList     []Migrable
	onceMigrationsList sync.Once
)

// createMigrationsTable - Creates the migrations table
// createMigrationsTable - Si occupa di creare la tabella delle migrazioni
func createMigrationsTable(conn db.SQLConnector) error {
	query := driver.GetCreateMigrationsTableQuery(conn.DriverName())
	_, err := conn.Exec(query)
	return err
}

// DoUpMigrations - Start the migrations
func DoUpMigrations() error {

	migrationManager := GetMigratorInstance()
	err := migrationManager.DoUpMigrations()
	if err != nil {
		return err
	}

	return nil
}

// DoDownMigrations - Start the rollbacks
func DoDownMigrations() error {

	migrationManager := GetMigratorInstance()
	err := migrationManager.DoDownMigrations()
	if err != nil {
		return err
	}

	return nil
}

// GetMigratorInstance - Returns the Migrator's instance
// GetMigratorInstance - Restituisce l'unica istanza di migrator
func GetMigratorInstance() *Migrator {

	onceMigrator.Do(func() {
		migrator = new(Migrator)
		migrator.migrationsList = migrationsList
	})

	return migrator
}

// InitMigrationsList - Sets the migrations list
func InitMigrationsList(ml []Migrable) {
	onceMigrationsList.Do(func() {
		migrationsList = ml
	})
}

// execDownMigrations - Executes the rollbacks at low level
func (m *Migrator) execDownMigrations(db db.SQLConnector) error {

	for i := len(m.migrationsList) - 1; i >= 0; i-- {

		migration := table.NewMigration(db)

		err := table.LoadMigrationByName(m.migrationsList[i].GetMigrationName(), migration)
		if err != nil {
			return err
		}

		if migration.GetTableRecord().IsLoaded() && migration.Status == 1 {

			_, err = db.Exec(m.migrationsList[i].Down())
			if err != nil {
				return err
			}

			migration.Status = 0
			err = record.Save(migration)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// execUpMigrations - Executes the migrations at low level
func (m *Migrator) execUpMigrations(db db.SQLConnector) error {

	for _, mi := range m.migrationsList {

		migration := table.NewMigration(db)

		err := table.LoadMigrationByName(mi.GetMigrationName(), migration)
		if err != nil {
			return err
		}

		if migration.GetTableRecord().IsLoaded() && migration.Status == 1 {
			continue
		}

		if migration.GetTableRecord().IsNew() {

			migration, err = table.InsertNewMigration(db, mi.GetMigrationName(), 0)
			if err != nil {
				return err
			}
		}

		_, err = db.Exec(mi.Up())
		if err != nil {
			return err
		}

		migration.Status = 1
		err = record.Save(migration)
		if err != nil {
			return err
		}
	}

	return nil
}

// DoDownMigrations - Executes the rollbacks defined in the migrations list
func (m *Migrator) DoDownMigrations() error {

	return transactions.WithTransactionx(db.GetSQLXFromSQLConnector(db.GetConnection()), func(tx db.SQLConnector) error {

		exist := db.TableExists(table.MigrationsTableName)

		if !exist {
			return nil
		}

		if err := m.execDownMigrations(tx); err != nil {
			return err
		}

		return nil
	})
}

// DoUpMigrations - Executes the migrations
func (m *Migrator) DoUpMigrations() error {

	return transactions.WithTransactionx(db.GetSQLXFromSQLConnector(db.GetConnection()), func(tx db.SQLConnector) error {

		exist := db.TableExists(table.MigrationsTableName)

		if !exist {

			if err := createMigrationsTable(tx); err != nil {
				return err
			}
		}

		if err := m.execUpMigrations(tx); err != nil {
			return err
		}

		return nil
	})
}
