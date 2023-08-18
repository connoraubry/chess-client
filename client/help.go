package client

import (
	"fmt"
	"strings"
)

func HelpBuilder(name string, usage string, commands map[string]string) string {
	var helpString string
	var fmtCmd []string

	for cmd, help := range commands {
		fmtCmd = append(fmtCmd, fmt.Sprintf("  %-16s %s", cmd, help))
	}
	combinedString := strings.Join(fmtCmd, "\n")

	if len(commands) == 0 {
		helpString = fmt.Sprintf(`%v

Usage:
  %v`, name, usage)
	} else {
		helpString = fmt.Sprintf(`%v

Usage:
  %v

Available Commands:
%v`, name, usage, combinedString)
	}
	return helpString
}

func (c *Client) commandHelpHandler(args []string) {

	if len(args) > 0 {
		switch args[0] {
		case "config":
			helpConfig(args[1:])
		default:
			fmt.Println("Help message not added yet!")
		}
	} else {
		var helpLine = map[string]string{
			"help [command]":   "For information on a command",
			"config [options]": "For configuring the client",
			"new [options]":    "Create a new game",
			"join [options]":   "Join an existing game",
			"info":             "Information on the client",
			"ping":             "Ping the server. Test connection",
			"quit":             "Quit",
		}
		fmt.Println(HelpBuilder("Chess client", "[command] [options", helpLine))
	}
}

func helpConfig(args []string) {

	if len(args) > 0 {
		switch args[0] {
		case "set":
			printConfigSetHelp()
		case "reset":
			printConfigResetHelp()
		default:
			printConfigHelp()
		}
	} else {
		printConfigHelp()
	}
}

func printConfigHelp() {
	var helpConfigHandlerCommands = map[string]string{
		"reset":             "Reset the configuration back to base",
		"set [arg] [value]": "Set a configuration property",
	}
	fmt.Println(HelpBuilder("Chess config handler", "config [command]", helpConfigHandlerCommands))
}

func printConfigSetHelp() {
	var helpConfigSetHandlerCommands = map[string]string{
		"server [val]": "Set server (string)",
		"port [val]":   "Set port (int)",
	}
	fmt.Println(HelpBuilder("Set a config value.", "config set [arg] [value]", helpConfigSetHandlerCommands))
}
func printConfigResetHelp() {
	var helpConfigSetHandlerCommands = map[string]string{}
	fmt.Println(HelpBuilder("Resets the config back to default.", "config reset", helpConfigSetHandlerCommands))
}

var mainHelpCommands = map[string]string{
	"config": "Manage CLI configuration",
	"join":   "Join a game",
	"new":    "Create a new game",
}

var HelpString string = HelpBuilder("Chess client", "chess [command]", mainHelpCommands)

var configGetHelpCommands = map[string]string{
	"server": "Get configured server",
	"port":   "Get configured port",
	"token":  "Get user token",
	"all":    "Get all configured values",
}
var ConfigGetHelp string = HelpBuilder("Get Configuration Values", "chess config get [option]", configGetHelpCommands)
var ConfigSetHelp string = HelpBuilder("Set Configuration Values", "chess config set [option] [value]", configGetHelpCommands)
