package migrations

// CreateUserTable - Definisce la strut per la generazione della tabella users
type CreateUserTable struct{}

// GetMigrationName - Restituisce il nome della migrazione
func (c CreateUserTable) GetMigrationName() string { return "create_user_table" }

// Down - Definisce la query di migrazione down
func (c CreateUserTable) Down() string { return "" }

// Up - Definisce la query di migrazione up
func (c CreateUserTable) Up() string { return "" }
