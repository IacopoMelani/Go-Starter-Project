package main

import (
	"github.com/IacopoMelani/Go-Starter-Project/command"
	"github.com/subosito/gotenv"
)

func main() {

	gotenv.Load()

	command.InterpretingHumanWord()
}
