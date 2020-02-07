package command

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	cacheconf "github.com/IacopoMelani/Go-Starter-Project/pkg/cache_config"

	"github.com/IacopoMelani/Go-Starter-Project/db"
	dbm "github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/migration"
	"github.com/IacopoMelani/Go-Starter-Project/pkg/models/table_record/table"
	"github.com/olekukonko/tablewriter"

	"github.com/IacopoMelani/Go-Starter-Project/config"

	"github.com/IacopoMelani/Go-Starter-Project/boot"
)

const (
	fire          = "fire"
	startServer   = "go!"
	migrate       = "go-migrate"
	rollback      = "go-rollback"
	migrateStatus = "go-migrate-status"
	showConfig    = "go-config"
)

// getDefaultMessage - Restituisce il messaggio default
func getDefaultMessage() string {
	return `
		commands:
			-fire go!               -> start the Server 
			-fire go-migrate        -> migrate database 
			-fire go-rollback       -> rollback database 
			-fire go-migrate-status -> show migrations status  
			-fire go-config         -> show the current environment 
	`
}

// migrateCommand - Si occupa di eseguire la migrazione del database
func migrateCommand() {

	db.InitMigrationsList()

	err := migration.DoUpMigrations()
	if err != nil {
		panic("Error during migrating, error: " + err.Error())
	}
	fmt.Println("Gotcha!")
	migrateStatusCommand()
}

// migrateStatusCommand - Si occupa di recuperare lo stato delle migrazioni
func migrateStatusCommand() {

	migrations, err := table.LoadAllMigrations(dbm.GetConnection())
	if err != nil {
		panic("Error loading migrations table, error: " + err.Error())
	}

	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"Date", "Name", "status"})
	table.SetBorder(false)

	for _, m := range migrations {

		data := []string{m.CreatedAt.String(), m.Name, strconv.FormatInt(int64(m.Status), 10)}
		table.Append(data)
	}

	table.Render()
}

// rollbackCommand - Si occupa di richiamare il rollback del database
func rollbackCommand() {

	db.InitMigrationsList()

	err := migration.DoDownMigrations()
	if err != nil {
		panic("Error during rollback, error: " + err.Error())
	}
	fmt.Println("Bye")
	migrateStatusCommand()
}

// InterpretingHumanWord - Si occupa di interpretare i comandi
func InterpretingHumanWord() {

	start := flag.String(fire, "", getDefaultMessage())

	flag.Parse()
	config := config.GetInstance()
	dbm.InitConnection("mysql", config.StringConnection)

	switch *start {

	case startServer:

		boot.InitServer()

	case showConfig:

		fmt.Println(cacheconf.GetCurrentConfig())

	case migrate:
		migrateCommand()

	case rollback:
		rollbackCommand()

	case migrateStatus:
		migrateStatusCommand()

	default:
		flag.PrintDefaults()
	}

}
