package main

import (
	"fmt"
	"log"
	"testDB/models"
)

func main() {

	user := models.User{}

	user.SetIsNew(false)

	user.SetName("max")

	id, err := user.Save()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(id)

}
