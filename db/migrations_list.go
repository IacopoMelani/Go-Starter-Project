package db

import (
	"github.com/IacopoMelani/Go-Starter-Project/db/migrations"
	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/migration"
)

// It define the migrations list, it's important to consider the order because they are executed sequentially
var migrationsList = []migration.Migrable{
	migrations.CreateTasksTable{},
	migrations.CreateCarsTable{},
}

// InitMigrationsList - Initiliaze the migrations list
func InitMigrationsList() {
	migration.InitMigrationsList(migrationsList)
}
