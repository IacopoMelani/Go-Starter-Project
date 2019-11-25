package db

import (
	"github.com/IacopoMelani/Go-Starter-Project/db/migrations"
	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/migration"
)

// Definisce la lista in sequenza di tutte le migrazioni
var migrationsList = []migration.Migrable{
	migrations.CreateTasksTable{},
	migrations.CreateCarsTable{},
}

// InitMigrationsList - Si occupa di inizializzare la lista delle migrazioni
func InitMigrationsList() {
	migration.InitMigrationsList(migrationsList)
}
