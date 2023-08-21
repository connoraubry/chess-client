package client

import (
	"fmt"
	"strings"
)

type entry struct {
	command string
	help    string
}

func HelpBuilder(name string, usage string, commands []entry) string {
	var helpString string
	var fmtCmd []string

	for _, entry := range commands {

		fmtCmd = append(fmtCmd, fmt.Sprintf("  %-20s %s", entry.command, entry.help))
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

	var no_commands []entry

	if len(args) > 0 {
		switch args[0] {
		case "config":
			helpConfig(args[1:])
		case "ping":
			fmt.Println(HelpBuilder("ping the server", "ping", no_commands))
		case "info":
			fmt.Println(HelpBuilder("Information on the chess cli", "info", no_commands))
		case "quit":
			fmt.Println(HelpBuilder("Quit the client", "quit", no_commands))
		default:
			fmt.Println("Help message not added yet!")
		}
	} else {
		var helpLine = []entry{
			{"help [command]", "For information on a command"},
			{"config [options]", "For configuring the client"},
			{"new [options]", "Create a new game"},
			{"join [options]", "Join an existing game"},
			{"info", "Information on the client"},
			{"ping", "Ping the server. Test connection"},
			{"quit", "Quit"},
		}
		fmt.Println(HelpBuilder("Chess client", "[command] [options]", helpLine))
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
	var helpConfigHandlerCommands = []entry{
		{"reset", "Reset the configuration back to base"},
		{"set [arg] [value]", "Set a configuration property"},
	}
	fmt.Println(HelpBuilder("Chess config handler", "config [command]", helpConfigHandlerCommands))
}

func printConfigSetHelp() {
	var helpConfigSetHandlerCommands = []entry{
		{"server [val]", "Set server (string)"},
		{"port [val]", "Set port (int)"},
	}
	fmt.Println(HelpBuilder("Set a config value.", "config set [arg] [value]", helpConfigSetHandlerCommands))
}

func printConfigResetHelp() {
	var helpConfigSetHandlerCommands []entry
	fmt.Println(HelpBuilder("Resets the config back to default.", "config reset", helpConfigSetHandlerCommands))
}
