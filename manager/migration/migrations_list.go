package migration

import (
	"github.com/IacopoMelani/Go-Starter-Project/db/migrations"
)

// Definisce la lista in sequenza di tutte le migrazioni
var migrationsList = []Migrable{
	migrations.CreateUserTable{},
}
