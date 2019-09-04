package migrations

// CreateTasksTable - Definisce la strut per la generazione della tabella users
type CreateTasksTable struct{}

// GetMigrationName - Restituisce il nome della migrazione
func (c CreateTasksTable) GetMigrationName() string {
	return "create_tasks_table"
}

// Down - Definisce la query di migrazione down
func (c CreateTasksTable) Down() string { return "DROP TABLE IF EXISTS tasks" }

// Up - Definisce la query di migrazione up
func (c CreateTasksTable) Up() string {
	return `CREATE TABLE IF NOT EXISTS tasks (
    record_id INT AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    start_date DATE,
    status INT NOT NULL,
    PRIMARY KEY (record_id)
	)`
}
