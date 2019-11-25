package command

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	cacheconf "github.com/IacopoMelani/Go-Starter-Project/pkg/cache_config"

	"github.com/IacopoMelani/Go-Starter-Project/db"

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

func migrateCommand() {

	db.InitMigrationsList()

	err := migration.DoUpMigrations()
	if err != nil {
		panic("Error during migrating, error: " + err.Error())
	}
	fmt.Println("Gotcha!")
}

func migrateStatusCommand() {

	migrations, err := table.LoadAllMigrations()
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

func rollbackCommand() {

	db.InitMigrationsList()

	err := migration.DoDownMigrations()
	if err != nil {
		panic("Error during rollback, error: " + err.Error())
	}
	fmt.Println("Bye")
}

// InterpretingHumanWord - Si occupa di interpretare i comandi
func InterpretingHumanWord() {

	start := flag.String(fire, "", getDefaultMessage())

	flag.Parse()

	switch *start {

	case startServer:

		boot.InitServer()
		break

	case showConfig:

		config.GetInstance()
		fmt.Println(cacheconf.GetCurrentConfig())
		break

	case migrate:

		migrateCommand()
		break

	case rollback:

		rollbackCommand()
		break

	case migrateStatus:

		migrateStatusCommand()
		break

	default:
		flag.PrintDefaults()
	}

}
