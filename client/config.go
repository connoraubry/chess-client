package client

import (
	"fmt"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type ClientConfig struct {
	Port   int
	Server string
	Token  string
	User   string
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
		default:
			printConfigSetHelp()
		}
	} else {
		printConfigHelp()
	}
	if writeFlag {
		SaveConfig(c.cfg)
	}
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
	c := ClientConfig{Port: 3030, Server: "http://localhost"}
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
