package command

import (
	"flag"
	"fmt"

	"github.com/IacopoMelani/Go-Starter-Project/manager/migration"

	"github.com/IacopoMelani/Go-Starter-Project/config"

	"github.com/IacopoMelani/Go-Starter-Project/boot"
)

const (
	fire        = "fire"
	startServer = "go!"
	migrate     = "go-migrate"
	rollback    = "go-rollback"
	showConfig  = "go-config"
)

func getDefaultMessage() string {
	return `
		commands:
			-fire go!         -> start the Server 
			-fire go-migrate  -> migrate database 
			-fire go-rollback -> rollback database 
			-fire go-config   -> show the current environment 
	`
}

func migrateCommand() {

	err := migration.DoUpMigrations()
	if err != nil {
		panic("Error during migrating, error: " + err.Error())
	}
	fmt.Println("Gotcha!")
}

func rollbackCommand() {

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
		fmt.Println(config.Config)
		break

	case migrate:

		migrateCommand()
		break

	case rollback:

		break

	default:
		flag.PrintDefaults()
	}

}
