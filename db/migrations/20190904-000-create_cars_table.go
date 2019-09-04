package migrations

// CreateCarsTable - Definisce la strut per la generazione della tabella cars
type CreateCarsTable struct{}

// GetMigrationName - Restituisce il nome della migrazione
func (c CreateCarsTable) GetMigrationName() string {
	return "create_cars_table"
}

// Down - Definisce la query di migrazione down
func (c CreateCarsTable) Down() string { return "DROP TABLE IF EXISTS cars" }

// Up - Definisce la query di migrazione up
func (c CreateCarsTable) Up() string {
	return `CREATE TABLE IF NOT EXISTS cars (
    record_id INT AUTO_INCREMENT,
    brand VARCHAR(255) NOT NULL,
    model VARCHAR(255) NOT NULL,
    color VARCHAR(255) NOT NULL,
    PRIMARY KEY (record_id)
	)`
}
