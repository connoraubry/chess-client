package client

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func NewHandler(args []string) {
	if len(args) > 1 {
		switch args[1] {
		case "game":
			log.Info("Creating new game")
		}
	}
}

func printNewHelp() {
	var helpLine = []entry{
		{"game", "Create a new game"},
	}
	fmt.Println(HelpBuilder("New game creator", "new [option]", helpLine))
}

func printNewGameHelp() {
	fmt.Println(HelpBuilder("Create a new game", "new game", []entry{}))
}
