package main

import (
	"log/slog"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	log := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &cli.App{
		Name:  "myapp",
		Usage: "make an explosive entrance",
		Action: func(c *cli.Context) error {
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Error("error running app", "error", err)
		os.Exit(1)
	}
}
