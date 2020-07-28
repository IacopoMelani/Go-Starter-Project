package migrations

// CreateCarsTable - Some example
type CreateCarsTable struct{}

// GetMigrationName -
func (c CreateCarsTable) GetMigrationName() string {
	return "create_cars_table"
}

// Down -
func (c CreateCarsTable) Down() string { return "DROP TABLE IF EXISTS cars" }

// Up -
func (c CreateCarsTable) Up() string {
	return `CREATE TABLE IF NOT EXISTS cars (
    record_id INT AUTO_INCREMENT,
    brand VARCHAR(255) NOT NULL,
    model VARCHAR(255) NOT NULL,
    color VARCHAR(255) NOT NULL,
    PRIMARY KEY (record_id)
	)`
}
