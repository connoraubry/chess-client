package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/connoraubry/chess-client/client"
	log "github.com/sirupsen/logrus"
)

//var (
//	logLevel = flag.String("level", "info", "Level to set logging at. [debug, info, warning, error]")
//)

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
		log.Debug("No arguments provided")
		//fmt.Println(client.HelpString)
	}
	fen, err := getFEN(3)
	if err != nil {
		log.Error(err)
	} else {
		log.WithField("fen", fen).Info("Got FEN")
	}

	//newGameId := createNewGame()
	//fmt.Println(newGameId)
}
func getTest() string {
	resp, err := http.Get("http://localhost:3030/api/v1/test")
	if err != nil {
		log.Error(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(body)

}

func getFEN(id int) (string, error) {
	url := fmt.Sprintf("http://localhost:3030/api/v1/fen/%d", id)
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	} else {
		log.Info("Successfully received request")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {

		return "", fmt.Errorf("STATUS returned %v. %v", resp.Status, string(body))
	}

	var result fenResp
	if err = json.Unmarshal(body, &result); err != nil {
		return "", err
	}
	log.WithField("fen", result.Fen).Info("Received FEN response")

	return result.Fen, nil

}

type fenResp struct {
	Fen string
}

func createNewGame() int {
	resp, err := http.Get("http://localhost:3030/api/v1/create")
	if err != nil {
		log.Error(err)
		return -1
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var result newGameResp
	if err = json.Unmarshal(body, &result); err != nil {
		log.Error("Can not unmarshal JSON")
	}

	log.WithField("id", result.Id).Info("Created new game")

	return result.Id
}

type newGameResp struct {
	Id int
}
