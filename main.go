package main

import (
	"github.com/Go-Starter-Project/boot"

	"github.com/subosito/gotenv"
)

func main() {

	gotenv.Load()

	boot.InitServer()
}
