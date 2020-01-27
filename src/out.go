package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

const FileUsage string = "Prints out the file instead of the true values, WARNING this will potentially print secrets to your terminal."
const OutUsage string = "Outputs the contents of the .secrets file, WARNING this will potentially print secrets to your terminal. Can be ran with `export $(crit out) to populate current environment`"

func getOutCmd() *cli.Command {
	cmd := &cli.Command{
		Name:    "out",
		Aliases: []string{"o"},
		Usage:   OutUsage,
		Action:  handleInput,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "file",
				Value: false,
				Usage: FileUsage,
			},
		},
	}

	return cmd
}

func handleInput(c *cli.Context) error {
	path := defaultPath()

	if c.String("path") != "" {
		path = c.String("path")
	}

	contents, err := getSecretAt(path)

	if err != nil {
		return err
	}

	if c.Bool("file") {
		fmt.Print(contents)
	} else {
		secrets, err := getSecrets(contents)

		if err != nil {
			return err
		}

		for _, s := range secrets {
			fmt.Println(s.ToKeyValue())
		}
	}

	return nil
}
