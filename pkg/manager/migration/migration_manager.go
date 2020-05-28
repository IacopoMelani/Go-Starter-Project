package migration

import (
	"sync"

	"github.com/jmoiron/sqlx"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db"
	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db/transactions"
	"github.com/IacopoMelani/Go-Starter-Project/pkg/models/table_record/table"

	record "github.com/IacopoMelani/Go-Starter-Project/pkg/models/table_record"
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
	migrator           *Migrator
	onceMigrator       sync.Once
	migrationsList     []Migrable
	onceMigrationsList sync.Once
)

// createMigrationsTable - Si occupa di creare la tabella delle migrazioni
func createMigrationsTable(conn db.SQLConnector) error {

	var query string

	switch conn.DriverName() {

	case db.DriverSQLServer:
		query = `
		IF NOT EXISTS (SELECT * FROM sysobjects WHERE name='migrations' and xtype='U')
		CREATE TABLE migrations (
		record_id BIGINT IDENTITY(1, 1) NOT NULL PRIMARY KEY,
		created_at DATETIME2(1) NOT NULL,
		name NVARCHAR(255) NOT NULL,
		status INT NOT NULL
		)`

	case db.DriverMySQL:
		query = `CREATE TABLE IF NOT EXISTS migrations (
		record_id INT AUTO_INCREMENT,
		created_at DATETIME NOT NULL,
		name VARCHAR(255) NOT NULL,
		status INT NOT NULL,
		PRIMARY KEY (record_id)
		)`
	}


	_, err := conn.Exec(query)

	return err
}

// DoUpMigrations - Esegue la migrazione del DB
func DoUpMigrations() error {

	migrationManager := GetMigratorInstance()
	err := migrationManager.DoUpMigrations()
	if err != nil {
		return err
	}

	return nil
}

// DoDownMigrations - Esegue il rollbacl del DB
func DoDownMigrations() error {

	migrationManager := GetMigratorInstance()
	err := migrationManager.DoDownMigrations()
	if err != nil {
		return err
	}

	return nil
}

// GetMigratorInstance - Restituisce l'unica istanza di migrator
func GetMigratorInstance() *Migrator {

	onceMigrator.Do(func() {
		migrator = new(Migrator)
		migrator.migrationsList = migrationsList
	})

	return migrator
}

// InitMigrationsList - Imposta la lista delle migrazioni
func InitMigrationsList(ml []Migrable) {
	onceMigrationsList.Do(func() {
		migrationsList = ml
	})
}

// execDownMigrations - Esegue fisicamente il rollback
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

// execUpMigrations - Si occupa di eseguire fisicamente la migrazione del database
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

// DoDownMigrations - Si occupa di fare il rollback delle tabelle definite in migrations_list
func (m *Migrator) DoDownMigrations() error {

	return transactions.WithTransactionx(db.GetConnection().(*sqlx.DB), func(tx db.SQLConnector) error {

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

// DoUpMigrations - Si occupa di migrare le tabelle definite in migrations_list
func (m *Migrator) DoUpMigrations() error {

	return transactions.WithTransactionx(db.GetConnection().(*sqlx.DB), func(tx db.SQLConnector) error {

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
