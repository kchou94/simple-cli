package main

import (
	"fmt"
	"os"

	"github.com/kchou94/simple-cli/tasks"
)

func main() {
	if err := tasks.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	tasks.AddCommands()
}
