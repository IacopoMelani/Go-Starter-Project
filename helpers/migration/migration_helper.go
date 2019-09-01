package migration

// Migrable - Definisce l'interfaccia per poter avviare una migrazione sul database
type Migrable interface {
	Down()
	Up()
}

// Migrator - Struct che si occupa di effettuare le migrazione sul database
type Migrator struct {
	migrationsList map[string]Migrable
}
