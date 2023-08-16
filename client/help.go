package client

import (
	"fmt"
	"strings"
)

func HelpBuilder(name string, usage string, commands map[string]string) string {
	var fmtCmd []string

	for cmd, help := range commands {
		fmtCmd = append(fmtCmd, fmt.Sprintf("  %-16s %s", cmd, help))
	}
	combinedString := strings.Join(fmtCmd, "\n")
	s := fmt.Sprintf(`%v

Usage:
  %v

Available Commands:
%v

%v`, name, usage, combinedString, FlagString)
	return s
}

var mainHelpCommands = map[string]string{
	"config": "Manage CLI configuration",
	"join":   "Join a game",
	"start":  "Start a new game",
}

var FlagString = `Flags:
  -h, --help			Help for chess
  --loglevel string		Level to log at. One of panic, fatal, error, warning, info, debug, trace (default: "info")
  -s, --server string	Chess server address
  -v, --version			Version for chess`

var HelpString string = HelpBuilder("Chess client", "chess [command]", mainHelpCommands)

var configHelpCommands = map[string]string{
	"get": "Get a current configuration",
	"set": "Set a configuration property",
}
var ConfigHelp string = HelpBuilder("Configure chess app", "chess config [command]", configHelpCommands)

var configGetHelpCommands = map[string]string{
	"server": "Get configured server",
	"port":   "Get configured port",
	"token":  "Get user token",
	"all":    "Get all configured values",
}
var ConfigGetHelp string = HelpBuilder("Get Configuration Values", "chess config get [option]", configGetHelpCommands)
var configSetHelpCommands = map[string]string{
	"server [val]": "Set server (string)",
	"port [val]":   "Set port (int)",
}
var ConfigSetHelp string = HelpBuilder("Set Configuration Values", "chess config set [option] [value]", configGetHelpCommands)
