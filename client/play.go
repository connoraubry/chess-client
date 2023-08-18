package client

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

const TERM_HEADER = "chess-client> "

func (c *Client) PlayHandler() {

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(TERM_HEADER)

		cmd, _ := reader.ReadString('\n')

		if runtime.GOOS == "windows" {
			cmd = strings.TrimRight(cmd, "\r\n")
		} else {
			cmd = strings.TrimSuffix(cmd, "\n")
		}

		log.Debugf("command received: %v", cmd)
		c.CommandHandler(cmd)
		if strings.Compare(cmd, "quit") == 0 {
			fmt.Println("Bye!")
			return
		}
	}
}

func (c *Client) CommandHandler(command string) {

	spaceSplit := strings.Split(command, " ")

	/*
		help [help]
		new [user / game]
		config
		join [ID]
	*/
	if len(spaceSplit) > 0 {
		cmd := spaceSplit[0]
		args := spaceSplit[1:]
		switch cmd {
		case "help":
			log.Debug("help inputted")
			c.commandHelpHandler(args)
		case "config":
			log.Debug("config inputted")
			c.commandConfigHandler(args)
		case "new":
			log.Debug("new inputted")
		case "join":
			log.Debug("join inputted")
		case "info":
			log.Debug("info inputted")
			fmt.Printf("%+v\n", c)
		case "ping":
			log.Debug("ping inputted")
			c.commandPingHandler()
		}
	}
}
func (c *Client) commandPingHandler() {
	server := c.GetServer()
	log.WithField("server", server).Info("Sending new request to server")
	c.Ping()
}
func (c *Client) commandNewHandler(args []string) {
	server := c.GetServer()
	log.WithField("server", server).Info("Sending new request to server")
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
