package main

import (
	"log"
	"os"

	_ "github.com/jormin/downloader/commands"
	"github.com/jormin/downloader/config"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "downloader",
		Usage:       "This is a tool to download video from third-paty video sites such as bilibili, aiyiqi, youku etc. Only support public free sources, no cracking of vip resources.",
		Version:     "v0.0.1",
		Description: "A simple tool to manage your todo list",
		Commands:    config.GetRegisteredCommands(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "d",
				Usage:       "the directory to save video",
				Required:    false,
				DefaultText: "",
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
