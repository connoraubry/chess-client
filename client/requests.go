package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

const API_PATH = "api/v1/"

func (c *Client) getBaseAddress() string {
	baseIP, _ := strings.CutSuffix(c.cfg.Server, "/")

	addr := fmt.Sprintf("%v:%v/%v", baseIP, c.cfg.Port, API_PATH)
	return addr
}

func (c *Client) Ping() {
	url := c.getBaseAddress()
	resp, err := http.Get(url)

	if err != nil {
		log.Errorf("Ping request unsuccessful. %v", err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Error reading ping response. %v", err)
	}

	log.WithField("resp", string(body)).Info("Ping responded")
}

func getTest(dest string) string {
	resp, err := http.Get(dest)
	if err != nil {
		log.Error(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(body)
}

type fenResp struct {
	Fen string
}

func (c *Client) Get(id int) (string, error) {
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

type createResp struct {
	ID    int
	Token string
}

func (c *Client) CreateNewGame() (createResp, error) {
	var result createResp
	url := fmt.Sprintf("%v%v", c.getBaseAddress(), "create")
	var request = []byte(`{}`)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(request))
	if err != nil {
		return result, err
	}
	log.Info("Successfully sent /create POST request")

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return result, fmt.Errorf("STATUS returned %v. %v", resp.Status, string(body))
	}

	if err = json.Unmarshal(body, &result); err != nil {
		return result, err
	}

	c.configSetGame(result.ID, result.Token)
	return result, nil
}

type gameResp struct {
	ID      int
	Fen     string
	Done    bool
	PgnPath string
}

func (c *Client) GetCurrentGame() (gameResp, error) {

	var result gameResp
	url := fmt.Sprintf("%v%v/%v", c.getBaseAddress(), "game", c.cfg.GameID)
	resp, err := http.Get(url)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return result, fmt.Errorf("STATUS returned %v. %v", resp.Status, string(body))
	}
	err = json.Unmarshal(body, &result)
	return result, err
}

func (c *Client) TakeMove(move string) error {
	url := fmt.Sprintf("%v%v", c.getBaseAddress(), "move")

	type takeMoveRequest struct {
		Move  string
		ID    int
		Token string
	}

	req := takeMoveRequest{Move: move, ID: c.cfg.GameID, Token: c.cfg.Token}
	reqByte, err := json.Marshal(req)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqByte))
	if err != nil {
		return err
	}
	log.Info("Successfully sent /move POST request")

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("STATUS returned %v. %v", resp.Status, string(body))
	}
	return nil
}
