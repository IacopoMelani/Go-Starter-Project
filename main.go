package main

import (
	"github.com/IacopoMelani/Go-Starter-Project/boot"

	"github.com/subosito/gotenv"
)

func main() {

	gotenv.Load()

	boot.InitServer()
}
