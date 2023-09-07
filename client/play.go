package client

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/connoraubry/chessbot-go/engine"
	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
)

const TERM_HEADER = "chess-client> "

type PlayType int

const (
	STANDARD PlayType = iota
	INGAME
)

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
			c.commandJoinHandler(args)
		case "info":
			log.Debug("info inputted")
			fmt.Printf("%+v\n", c)
		case "ping":
			log.Debug("ping inputted")
			c.commandPingHandler()
		case "game":
			log.Debug("game inputted")
			c.commandGameHandler()
		case "move":
			log.Debug("move inputted")
			c.commandMoveHandler(args)
		}
	}
}

func (c *Client) commandMoveHandler(args []string) {
	if len(args) > 0 {
		move := args[0]
		err := c.TakeMove(move)
		if err != nil {
			log.Errorf("Error taking move: %v", err)
		}
	} else {
		helpMove()
	}

}
func (c *Client) commandJoinHandler(args []string) {
	if len(args) > 0 {
		gameIDstring := args[0]
		gameID, err := strconv.Atoi(gameIDstring)
		if err != nil {
			log.WithField("id", gameIDstring).Error("Unable to parse id into integer")
			return
		}
		err = c.JoinGame(gameID)
		if err != nil {
			log.Errorf("Error joining game: %v", err)
			return
		}
		fmt.Printf("Joined game %v\n", gameID)
	} else {
		helpJoin()
	}
}
func (c *Client) commandGameHandler() {

	game, err := c.GetCurrentGame()
	if err != nil {
		log.Errorf("Error when getting current game: %v", err)
	}
	log.Info(game)
	e := engine.NewEngine(engine.OptFenString(game.Fen))
	e.Print(0, game.LastMove)
}

func (c *Client) commandNewHandler(args []string) {
	if len(args) > 0 {
		switch args[0] {
		case "game":
			log.Debug("new game inputted")
			log.Info("Creating a new game")

			var isDev bool

			if len(args) > 1 && args[1] == "dev" {
				isDev = true
			}

			res, err := c.CreateNewGame(isDev)
			if err != nil {
				log.Errorf("Error creating new game: %v", err)
			} else {
				fields := log.Fields{
					"id":    res.ID,
					"token": res.Token,
				}
				log.WithFields(fields).Info("New game created")
				fmt.Printf("New game created. ID = %v\n", res.ID)
			}
		}
	} else {
		log.Debug("No args provided to new")
		printNewHelp()
	}
}

func (c *Client) commandPingHandler() {
	server := c.GetServer()
	log.WithField("server", server).Info("Sending new request to server")
	c.Ping()
}
