package command

import (
	"flag"
	"fmt"

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

// InterpretingHumanWord - Si occupa di interpretare i comandi
func InterpretingHumanWord() {

	start := flag.String(fire, "", getDefaultMessage())

	flag.Parse()

	if *start == "" {
		flag.PrintDefaults()
	}

	switch *start {
	case startServer:
		boot.InitServer()
		break
	case showConfig:
		config.GetInstance()
		fmt.Println(config.Config)
		break
	default:
		flag.PrintDefaults()
	}
}
