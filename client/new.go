package client

import log "github.com/sirupsen/logrus"

func NewHandler(args []string) {
	if len(args) > 1 {
		switch args[1] {
		case "game":
			log.Info("Creating new game")
		}
	}
}
