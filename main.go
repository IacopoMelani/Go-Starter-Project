package main

import (
	"github.com/IacopoMelani/Go-Starter-Project/command"
	"github.com/subosito/gotenv"
)

func main() {

	if err := gotenv.Load(); err != nil {
		panic(err)
	}

	command.InterpretingHumanWord()
}
