package main

import (
	"github.com/mobml/ant/cmd"
	"github.com/mobml/ant/database"
)

func main() {

	cmd.Execute()
	defer database.DB.Close()
}
