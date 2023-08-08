package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/connoraubry/chess-client/client"
	log "github.com/sirupsen/logrus"
)

var (
	logLevel = flag.String("level", "info", "Level to set logging at. [debug, info, warning, error]")
)

func main() {
	flag.Parse()
	c := client.New()
	c.SetupLogging(*logLevel)

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "config":
			client.Config()
		default:
			fmt.Println(client.HelpString)
		}
	} else {
		log.Debug("No")
		fmt.Println(client.HelpString)
	}
}
