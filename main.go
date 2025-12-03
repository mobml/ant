package main

import (
	"log"

	"github.com/mobml/ant/cmd"
	"github.com/mobml/ant/database"
)

func main() {

	if err := database.InitDB("app.db"); err != nil {
		log.Fatal(err)
	}
	defer database.CloseDB()

	if err := database.Migrate(database.DB()); err != nil {
		log.Fatal(err)
	}

	cmd.Execute()
}
