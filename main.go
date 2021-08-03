package main

import (
	"log"
	"os"

	_ "github.com/jormin/download/commands"
	"github.com/jormin/download/config"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "todo",
		Usage:       "A simple tool to manage your todo list",
		Version:     "v1.0.0",
		Description: "A simple tool to manage your todo list",
		Commands:    config.GetRegisteredCommands(),
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
