package config

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// commands config
var cmdCfg = map[string][]*cli.Command{}

// commands
var commands []*cli.Command

// root
var root = "root"

// RegisterCommand register command that registered in function `init`
func RegisterCommand(parent string, command *cli.Command) {
	if parent == "" {
		parent = root
	}
	if _, ok := cmdCfg[parent]; !ok {
		cmdCfg[parent] = []*cli.Command{}
	}
	cmdCfg[parent] = append(cmdCfg[parent], command)
}

// GetRegisteredCommands get registered commands
func GetRegisteredCommands() []*cli.Command {
	ParseRegisteredCommands(root, nil)
	return commands
}

// ParseRegisteredCommands parse all registered commands
func ParseRegisteredCommands(path string, parent *cli.Command) {
	if parent == nil {
		for _, command := range cmdCfg[path] {
			ParseRegisteredCommands(path, command)
			commands = append(commands, command)
		}
	} else {
		path = fmt.Sprintf("%s.%s", path, parent.Name)
		if _, ok := cmdCfg[path]; ok {
			for _, child := range cmdCfg[path] {
				ParseRegisteredCommands(path, child)
				parent.Subcommands = append(parent.Subcommands, child)
			}
		}
	}
}
