package db

import (
	"log"
	"sync"

	"github.com/jmoiron/sqlx"
)

// dbContainer - Defines a container for a connection
type dbContainer struct {
	ok         chan bool
	driver     string
	connection string
	once       sync.Once
	conn       *sqlx.DB
}

// newDbContainer - Returns a new dbContainer instance
func newDbContainer(driver string, connection string) *dbContainer {

	conn := new(dbContainer)

	conn.ok = make(chan bool, 1)
	conn.driver = driver
	conn.connection = connection

	return conn
}

// getConnection - Returns the connection resource
func (d *dbContainer) getConnection() SQLConnector {
	<-d.ok
	return d.conn
}

// initConnection - Initializes the connection
func (d *dbContainer) initConnection() {

	d.once.Do(func() {

		db, err := sqlx.Open(d.driver, d.connection)
		if err != nil {
			log.Panic(err)
		}

		if err := db.Ping(); err != nil {
			log.Panic(err)
		}

		d.conn = db

		d.ok <- true
		close(d.ok)
	})
}
