package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

func main() {

	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	c := cli.NewCLI("cliexample", "0.0.1")
	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"server": func() (cli.Command, error) {
			return &ServerCommand{
				Ui: &cli.ColoredUi{
					Ui:          ui,
					OutputColor: cli.UiColorBlue,
				},
			}, nil
		},
		"agent": func() (cli.Command, error) {
			return &AgentCommand{
				Ui: &cli.ColoredUi{
					Ui:          ui,
					OutputColor: cli.UiColorGreen,
				},
			}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	os.Exit(exitStatus)
}
