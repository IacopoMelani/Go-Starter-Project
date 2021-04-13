package db

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	// Carica il driver sql per la connessione al db
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"

	"sync"
)

// SQLConnector - Generalizes a sql connection
type SQLConnector interface {
	DriverName() string
	Exec(string, ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Prepare(string) (*sql.Stmt, error)
	Preparex(string) (*sqlx.Stmt, error)
	Query(string, ...interface{}) (*sql.Rows, error)
	Queryx(string, ...interface{}) (*sqlx.Rows, error)
	QueryRow(string, ...interface{}) *sql.Row
	QueryRowx(string, ...interface{}) *sqlx.Row
	Select(dest interface{}, query string, args ...interface{}) error
}

// InvalidConnectionKeyError - Defines the error for an invalid key provided to retrive or use a connection instance
type InvalidConnectionKeyError struct {
	msg string
}

// NewInvalidConnectionKeyError - Returns a new InvalidConnectionKeyError instance
func NewInvalidConnectionKeyError(msg string) *InvalidConnectionKeyError {
	return &InvalidConnectionKeyError{msg}
}

// Error - Implements error interface
func (e *InvalidConnectionKeyError) Error() string {
	return e.msg
}

// DefaultConnectionName - default connection key
const DefaultConnectionName = "default"

// Defines all possible sql drivers
const (
	DriverSQLServer = "mssql"
	DriverMySQL     = "mysql"
)

// connectionsPool - Defines a pool of connection
type connectionsPool struct {
	ok          map[string]chan bool
	once        map[string]*sync.Once
	connections map[string]*dbContainer
	mu          sync.Mutex
}

var (
	pool *connectionsPool
)

// getSimpleSelectQueryForTable - Returns a simple select "LIMIT 1" query string for a specific connection key and table
func getSimpleSelectQueryForTable(driver string, table string) string {

	query := ""

	switch driver {
	case DriverMySQL:
		query = "SELECT * FROM " + table + " LIMIT 1"
	case DriverSQLServer:
		query = "SELECT TOP 1 * FROM " + table
	}

	return query
}

// init - Initialize db package
func init() {

	pool = new(connectionsPool)

	pool.connections = make(map[string]*dbContainer)
	pool.ok = make(map[string]chan bool)
	pool.once = make(map[string]*sync.Once)
}

// initChanForKey - Initializes the the chan bool for provided key
func initChanForKey(key string) {

	pool.mu.Lock()
	defer pool.mu.Unlock()

	if _, ok := pool.ok[key]; !ok {
		pool.ok[key] = make(chan bool, 1)
	}
}

// initConnection - Intializes the connection for specific key and values
func initConnection(key string, drvName string, connection string) {

	pool.connections[key] = newDbContainer(drvName, connection)
	pool.connections[key].initConnection()

	pool.ok[key] <- true
	close(pool.ok[key])
}

// initSyncOnceForKey - Initializes sync.once for provided key
func initSyncOnceForKey(key string) {

	pool.mu.Lock()
	defer pool.mu.Unlock()

	if _, ok := pool.once[key]; !ok {
		pool.once[key] = &sync.Once{}

	}
}

// DriverName - Returns default driver name
func DriverName() string {
	driver, err := DriverNameByKey(DefaultConnectionName)
	if err != nil {
		panic(err)
	}
	return driver
}

// DriverNameByKey - Returns driver name for the connection key provided
func DriverNameByKey(key string) (string, error) {

	conn, err := GetConnectionWithKey(key)
	if err != nil {
		return "", err
	}

	return conn.DriverName(), nil
}

// GetConnection - Returns the default instance of db connection
func GetConnection() SQLConnector {

	conn, err := GetConnectionWithKey(DefaultConnectionName)
	if err != nil {
		panic(err)
	}

	return conn
}

// GetConnectionWithKey - Returns an istance of the db connection by connection key (for multiple connections)
func GetConnectionWithKey(key string) (SQLConnector, error) {

	pool.mu.Lock()
	defer pool.mu.Unlock()

	if _, ok := pool.connections[key]; !ok {
		return nil, NewInvalidConnectionKeyError(fmt.Sprintf("Invalid connection key %s", key))
	}

	<-pool.ok[key]

	return pool.connections[key].getConnection(), nil
}

// GetSQLXFromSQLConnector - Returns a ptr of sqlx.DB from a SQLConnector
func GetSQLXFromSQLConnector(db SQLConnector) *sqlx.DB {
	return db.(*sqlx.DB)
}

// InitConnection - Initialize the connection with driver and connection string
func InitConnection(drvName string, connection string) {
	InitConnectionWithKey(DefaultConnectionName, drvName, connection)
}

// InitConnectionWithKey - Initialize the connection for specific key with driver and connection string
func InitConnectionWithKey(key string, drvName string, connection string) {

	initSyncOnceForKey(key)
	initChanForKey(key)

	pool.once[key].Do(func() {
		initConnection(key, drvName, connection)
	})
}

// Query - Executes the query and return a *Rows instance
func Query(query string, args ...interface{}) (*sqlx.Rows, error) {
	return QueryWithKey(DefaultConnectionName, query, args...)
}

// QueryWithKey - Executes the query for a specific connection key and return a *Rows instance
func QueryWithKey(key string, query string, args ...interface{}) (*sqlx.Rows, error) {

	db, err := GetConnectionWithKey(key)
	if err != nil {
		return nil, err
	}

	rows, err := db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

// QueryOrPanic - Executes the query and return a *Rows instance, panics if error occurs
func QueryOrPanic(query string, args ...interface{}) *sqlx.Rows {
	return QueryOrPanicWithKey(DefaultConnectionName, query, args...)
}

// QueryOrPanicWithKey - Executes the query for a specific connection key and return a *Rows instance, panics if error occurs
func QueryOrPanicWithKey(key string, query string, args ...interface{}) *sqlx.Rows {

	rows, err := QueryWithKey(key, query, args...)
	if err != nil {
		panic(err)
	}

	return rows
}

// TableExists - Returns true if table exists otherwise false
func TableExists(tableName string) bool {
	return TableExistsWithKey(DefaultConnectionName, tableName)
}

// TableExistsWithKey - Returns true if table exists otherwise false for a specific connection key
func TableExistsWithKey(key string, tableName string) bool {

	db, err := GetConnectionWithKey(key)
	if err != nil {
		panic(err)
	}

	query := getSimpleSelectQueryForTable(db.DriverName(), tableName)
	rows, err := db.Queryx(query)
	if err != nil {
		return false
	}
	defer rows.Close()

	return true
}
