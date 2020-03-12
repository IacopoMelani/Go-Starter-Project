package migrations

// CreateTasksTable - Some example
type CreateTasksTable struct{}

// GetMigrationName -
func (c CreateTasksTable) GetMigrationName() string {
	return "create_tasks_table"
}

// Down -
func (c CreateTasksTable) Down() string { return "DROP TABLE IF EXISTS tasks" }

// Up -
func (c CreateTasksTable) Up() string {
	return `CREATE TABLE IF NOT EXISTS tasks (
    record_id INT AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    start_date DATE,
    status INT NOT NULL,
    PRIMARY KEY (record_id)
	)`
}
