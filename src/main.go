package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path"
	"strings"

	"github.com/fatih/color"
	cli "github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

const VERSION = "0.2"

const (
	NO_CMD  = iota + 1
	NO_FILE = iota
)

func main() {
	app := &cli.App{
		Name:  "crit",
		Usage: "Launch a program with injected secrets",
		Commands: []*cli.Command{
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "Shows the current version",
				Action: func(c *cli.Context) error {
					fmt.Println(VERSION)
					return nil
				},
			},
			getStartCmd(),
			getOutCmd(),
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}

// Inject injects secrets into a process
func Inject(c *cli.Context) error {
	app := c.Args()

	if !app.Present() {
		color.Red("Must provide an application to launch")
		cli.ShowAppHelpAndExit(c, NO_CMD)
	}

	var cmd []string
	for i := 0; i < app.Len(); i++ {
		cmd = append(cmd, app.Get(i))
	}

	// They probably used string encapsulation, so we must break it up
	if len(cmd) == 1 {
		cmd = strings.Split(cmd[0], " ")
	}

	executable := exec.Command(cmd[0], cmd[1:]...)

	color.Cyan("Executing: %s\n", executable)

	if err := runCommand(executable); err != nil {
		return err
	}

	return nil
}

func getHome() string {
	usr, err := user.Current()

	if err != nil {
		panic(err)
	}

	return usr.HomeDir
}

func getSecretAt(filepath string) (string, error) {
	if _, err := os.Stat(filepath); err != nil && os.IsNotExist(err) {
		color.Red(fmt.Sprintf("❌ File at path %s does not exist", filepath))
		os.Exit(NO_FILE)
	}

	file, err := os.Open(filepath)

	if err != nil {
		return "", err
	}

	if b, err := ioutil.ReadAll(file); err != nil {
		return "", err
	} else {
		return string(b), nil
	}
}

func getSecretDefault() (string, error) {
	return getSecretAt(path.Join(getHome(), ".secrets"))
}

func getSecrets(text string) ([]Secret, error) {
	file := &SecretFile{}

	err := yaml.Unmarshal([]byte(text), &file)

	if err != nil {
		return []Secret{}, err
	}

	secrets := []Secret{}

	for _, puller := range file.Pullers {
		secret, err := puller.toSecret()

		if err != nil {
			return []Secret{}, err
		}

		secrets = append(secrets, secret)
	}

	return secrets, nil
}

func runCommand(c *exec.Cmd) error {
	contents, err := getSecretDefault()

	if err != nil {
		log.Fatal(err)
	}

	secret, err := getSecrets(contents)
	envs := os.Environ()

	for _, s := range secret {
		envs = append(envs, s.ToKeyValue())
	}

	c.Env = envs
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	res := c.Run()

	return res
}
