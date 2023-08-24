package client

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type ClientConfig struct {
	Port      int
	Server    string
	GameID    int
	Token     string
	User      string
	TermColor string
}

const CONFIG_PATH = "./config.yaml"

func (cfg *ClientConfig) Print() {
	fmt.Printf("Server: %v\n", cfg.Server)
	fmt.Printf("Port: %v\n", cfg.Port)
	fmt.Printf("User: %v\n", cfg.User)
	fmt.Printf("Token: %v\n", cfg.Token)
}

func (c *Client) commandConfigHandler(args []string) {
	if len(args) > 0 {

		switch args[0] {
		case "reset":
			c.cfg = ResetConfig()
		case "set":
			c.commandConfigSetHandler(args[1:])
		default:
			printConfigHelp()
		}
	} else {
		c.cfg.Print()
	}
}

func printConfigHelp() {
	var helpConfigHandlerCommands = []entry{
		{"reset", "Reset the configuration back to base"},
		{"set [arg] [value]", "Set a configuration property"},
	}
	fmt.Println(HelpBuilder("Chess config handler", "config [command]", helpConfigHandlerCommands))
}

func printConfigResetHelp() {
	var helpConfigSetHandlerCommands []entry
	fmt.Println(HelpBuilder("Resets the config back to default.", "config reset", helpConfigSetHandlerCommands))
}

func (c *Client) commandConfigSetHandler(args []string) {
	writeFlag := false
	if len(args) > 1 {
		switch args[0] {
		case "server":
			c.cfg.Server = args[1]
			writeFlag = true
		case "port":
			newPort, err := strconv.Atoi(args[1])
			if err != nil {
				log.Fatalf("Invalid integer provided: %v", args[1])
			}
			c.cfg.Port = newPort
			writeFlag = true
		case "color":
			writeFlag = c.configSetColor(args[1])
		default:
			printConfigSetHelp()
		}
	} else {
		printConfigSetHelp()
	}
	if writeFlag {
		SaveConfig(c.cfg)
	}
}

func (c *Client) configSetColor(color string) bool {
	log.Debugf("Requesting color change to %v", color)
	oldColor := c.cfg.TermColor
	c.cfg.TermColor = color
	res := c.SetupColor()

	if !res {
		c.cfg.TermColor = oldColor
		log.Errorf("Invalid color: %v", color)
		printConfigSetColorHelp()
	}
	return res
}

func (c *Client) configSetGame(gameID int, token string) {
	c.cfg.GameID = gameID
	c.cfg.Token = token

	SaveConfig(c.cfg)
}

func printConfigSetHelp() {
	var helpConfigSetHandlerCommands = []entry{
		{"server [val]", "Set server (string)"},
		{"port [val]", "Set port (int)"},
		{"color [val]", "Change color of prompt"},
	}
	fmt.Println(HelpBuilder("Set a config value.", "config set [arg] [value]", helpConfigSetHandlerCommands))
}

func printConfigSetColorHelp() {
	var helpConfigSetColorCommands = []entry{
		{color.RedString("red"), ""},
		{color.YellowString("yellow"), ""},
		{color.BlueString("blue"), ""},
		{color.GreenString("green"), ""},
		{color.CyanString("cyan"), ""},
		{color.MagentaString("magenta"), ""},
		{"white", ""},
	}
	fmt.Println(HelpBuilder("Set the terminal color", "config set color [value]", helpConfigSetColorCommands))
}

func LoadConfig() ClientConfig {

	if _, err := os.Stat(CONFIG_PATH); err != nil {
		log.Info("Config file does not exist. Creating")
		return NewConfig()
	}

	return ReadFile()
}
func ResetConfig() ClientConfig {
	return NewConfig()
}
func NewConfig() ClientConfig {
	c := ClientConfig{Port: -1, Server: "http://localhost", TermColor: "white"}
	SaveConfig(c)
	return c
}

func ReadFile() ClientConfig {

	var cfg ClientConfig

	f, err := os.ReadFile(CONFIG_PATH)
	if err != nil {
		log.Errorf("Error reading config file: %v", err)
	}

	err = yaml.Unmarshal(f, &cfg)
	if err != nil {
		log.Errorf("Parsing yaml file failed: %v", err)
		log.Debug("Error parsing yaml file. Deleting config and starting new")
		cfg = NewConfig()
	}
	return cfg
}

func SaveConfig(cfg ClientConfig) {

	encoding, err := yaml.Marshal(cfg)
	if err != nil {
		log.Fatalf("error marshalling config file: %v", err)
	}

	err = os.WriteFile(CONFIG_PATH, encoding, 0644)
	if err != nil {
		log.Fatalf("error opening/creating file: %v", err)
	}
	log.Info("Successfully updated config file")
}
