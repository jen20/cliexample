//agent_command.go
package main

import (
	"github.com/mitchellh/cli"
)

type AgentCommand struct {
	Ui cli.Ui
}

func (c *AgentCommand) Run(_ []string) int {
	c.Ui.Output("Would run an agent here")
	return 0
}

func (c *AgentCommand) Help() string {
	return "Run as an agent (detailed help information here)"
}

func (c *AgentCommand) Synopsis() string {
	return "Run as an agent"
}
