package main

import (
	"flag"
	"fmt"
	"github.com/mitchellh/cli"
)

type ServerCommand struct {
	HttpPort int
	Ui       cli.Ui
}

func (c *ServerCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("agent", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }

	cmdFlags.IntVar(&c.HttpPort, "http-port", 80, "The port on which to run the HTTP server")
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	c.Ui.Output(fmt.Sprintf("Would run a server here on port %d", c.HttpPort))
	return 0
}

func (c *ServerCommand) Help() string {
	return "Run as a server (detailed help information here)"
}

func (c *ServerCommand) Synopsis() string {
	return "Run as a server"
}
