package client

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
)

const TERM_HEADER = "chess-client> "

func (c *Client) PlayHandler() {

	reader := bufio.NewReader(os.Stdin)
	for {
		color.Set(c.colorHandler)
		fmt.Print(TERM_HEADER)
		color.Unset()

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
			c.commandNewHandler(args)
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

func (c *Client) commandNewHandler(args []string) {
	if len(args) > 0 {
		switch args[0] {
		case "game":
			log.Debug("new game inputted")
			log.Info("Creating a new game")
			gameID, err := c.CreateNewGame()
			if err != nil {
				log.Errorf("Error creating new game: %v", err)
			} else {
				log.WithField("id", gameID).Info("New game created")
			}
		}
	} else {
		log.Debug("No args provided to new")
	}
}

func (c *Client) commandPingHandler() {
	server := c.GetServer()
	log.WithField("server", server).Info("Sending new request to server")
	c.Ping()
}
