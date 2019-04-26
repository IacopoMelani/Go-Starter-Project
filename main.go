package main

import (
	"Go-Starter-Project/boot"

	"github.com/subosito/gotenv"
)

func main() {

	gotenv.Load()

	boot.InitServer()
}
