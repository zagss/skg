package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/zagss/skg/project"
)

func main() {
	app := &cli.App{
		Name: "skg",
		Commands: []*cli.Command{
			{
				Name:      "create",
				Aliases:   []string{"c"},
				Usage:     "create a new project",
				UsageText: "skg create [options] <app-name>",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "name",
						Aliases: []string{"n"},
						Usage:   "set the project name",
					},
					&cli.StringFlag{
						Name:    "path",
						Aliases: []string{"p"},
						Usage:   "set the project path",
					},
				},
				Action: func(ctx *cli.Context) error {
					path := ctx.String("path")
					name := ctx.String("name")
					return project.Create(path, name)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
