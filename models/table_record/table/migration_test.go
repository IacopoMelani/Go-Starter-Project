package table

import (
	"testing"

	"github.com/IacopoMelani/Go-Starter-Project/db"

	"github.com/subosito/gotenv"
)

func TestMigration(t *testing.T) {

	gotenv.Load("./../../../.env")

	db := db.GetConnection()
	conn, err := db.Begin()
	if err != nil {
		t.Fatal(err.Error())
	}
	defer conn.Rollback()

	mName := "test_migration"

	migration, err := InsertNewMigration(mName, 1)
	if err != nil {
		conn.Rollback()
		t.Fatal(err.Error())
	}

	m := NewMigration()

	err = LoadMigrationByName(mName, m)
	if err != nil {
		conn.Rollback()
		t.Fatal(err.Error())
	}

	if m.Name != migration.Name || m.Status != migration.Status {
		conn.Rollback()
		t.Fatal("Operazione di migrazione errata")
	}
}
