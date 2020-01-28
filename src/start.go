package main

import "github.com/urfave/cli/v2"

func getStartCmd() *cli.Command {
	cmd := &cli.Command{
		Name:    "start",
		Aliases: []string{"s"},
		Usage:   "Start a program with secrets injected",
		Action:  Inject,
	}

	return cmd
}
