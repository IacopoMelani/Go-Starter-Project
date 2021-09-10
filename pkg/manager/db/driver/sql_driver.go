package driver

import "strings"

// Defines all possible sql drivers
const (
	DriverSQLServer = "mssql"
	DriverMySQL     = "mysql"
)

// GetCreateEmptyTableQuery - Generate a sql query to create an empty table based on a sql driver
func GetCreateEmptyTableQuery(drv string, table string) string {

	query := ""

	switch drv {
	case DriverMySQL:
		query = `CREATE TABLE IF NOT EXISTS ` + table + ` (
			record_id INT AUTO_INCREMENT,
			PRIMARY KEY (record_id)
			)`
	case DriverSQLServer:
		query = `IF NOT EXISTS (SELECT * FROM sysobjects WHERE name='testTable' and xtype='U')
		CREATE TABLE ` + table + ` (
			record_id BIGINT IDENTITY(1, 1) NOT NULL PRIMARY KEY
		)`
	}

	return query
}

// GetCreateMigrationsTableQuery - Generate a sql query to create the migrations table based on a sql driver
func GetCreateMigrationsTableQuery(drv string) string {

	query := ""

	switch drv {
	case DriverSQLServer:
		query = `
		IF NOT EXISTS (SELECT * FROM sysobjects WHERE name='migrations' and xtype='U')
		CREATE TABLE migrations (
		record_id BIGINT IDENTITY(1, 1) NOT NULL PRIMARY KEY,
		created_at DATETIME2(1) NOT NULL,
		name NVARCHAR(255) NOT NULL,
		status INT NOT NULL
		)`
	case DriverMySQL:
		query = `CREATE TABLE IF NOT EXISTS migrations (
		record_id INT AUTO_INCREMENT,
		created_at DATETIME NOT NULL,
		name VARCHAR(255) NOT NULL,
		status INT NOT NULL,
		PRIMARY KEY (record_id)
		)`
	}

	return query
}

// GetDropTableQuery - Generate a sql query to drop a table based on a sql driver
func GetDropTableQuery(drv string, table string) string {

	query := ""

	switch drv {
	case DriverMySQL:
		query = "DROP TABLE IF EXISTS " + table
	case DriverSQLServer:
		query = `
		IF EXISTS (SELECT * FROM sysobjects WHERE name='` + table + `' and xtype='U')
		DROP TABLE ` + table
	}

	return query
}

// GetInsertQuery - Generate a sql query to insert a record into a table based on a sql driver
func GetInsertQuery(drv string, table string, columns []string, values []string) string {

	query := ""

	switch drv {
	case DriverSQLServer:
		query = "INSERT INTO " + table + " (" + strings.Join(columns, ", ") + ") OUTPUT INSERTED.* VALUES ( " + strings.Join(values, ", ") + " )"
	case DriverMySQL:
		query = "INSERT INTO " + table + " (" + strings.Join(columns, ", ") + ") VALUES ( " + strings.Join(values, ", ") + " )"
	}

	return query
}

// GetSelectOneRowQuery - Generate a sql query to select one row from a table based on a sql driver
func GetSelectOneRowQuery(drv string, table string) string {

	query := ""

	switch drv {
	case DriverMySQL:
		query = "SELECT * FROM " + table + " LIMIT 1"
	case DriverSQLServer:
		query = "SELECT TOP 1 * FROM " + table
	}

	return query
}

// GetUpdateQuery - Generate a sql query to update a record in a table based on a sql driver
func GetUpdateQuery(drv string, table string, primaryKeyName string, columns []string) string {

	query := ""

	switch drv {
	case DriverSQLServer:
		query = "UPDATE  " + table + " SET " + strings.Join(columns, ", ") + " OUTPUT INSERTED.* WHERE " + primaryKeyName + " = ?"
	case DriverMySQL:
		query = "UPDATE  " + table + " SET " + strings.Join(columns, ", ") + " WHERE " + primaryKeyName + " = ?"
	}

	return query
}
