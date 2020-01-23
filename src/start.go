package main

import "github.com/urfave/cli/v2"

func getStartCmd() *cli.Command {
	cmd := &cli.Command{
		Name:    "start",
		Aliases: []string{"s"},
		Usage:   "Start a program with secrets injected",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "path",
				Aliases: []string{"p"},
				Usage:   "Loads the secret file from `PATH`",
			},
		},
		Action: Inject,
	}

	return cmd
}
