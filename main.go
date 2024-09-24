package main

import (
	"fmt"
	"github.com/mobml/ant/cmd/ant"
	"os"
)

func main() {
	rootCmd := ant.RootCommand()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
