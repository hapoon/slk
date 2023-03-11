package main

import (
	"fmt"
	"os"

	"github.com/hapoon/slk/action"
	"github.com/urfave/cli/v2"
)

const (
	name    = "slk"
	version = "0.1.0"
)

func main() {
	app := &cli.App{
		Version:     version,
		Description: "A cli application for slack",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "V", Usage: "vervose"},
		},
		Commands: []*cli.Command{
			{
				Name:   "init",
				Usage:  "Initialize setting",
				Action: action.ActInit,
			},
			{
				Name:   "wh",
				Usage:  "messaging by Incoming WebHooks",
				Action: action.ActWebHooks,
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
