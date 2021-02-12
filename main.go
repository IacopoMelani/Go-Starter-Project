package main

import (
	"github.com/IacopoMelani/Go-Starter-Project/command"
	"github.com/IacopoMelani/Go-Starter-Project/pkg/helpers/env"
)

func main() {

	env.LoadEnvFile(".env", true)

	command.InterpretingHumanWord()
}
