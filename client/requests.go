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
	ID int
}

func (c *Client) CreateNewGame() (int, error) {
	url := fmt.Sprintf("%v%v", c.getBaseAddress(), "create")
	var request = []byte(`{}`)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(request))
	if err != nil {
		return -1, err
	}
	log.Info("Successfully sent /create POST request")

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return -1, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return -1, fmt.Errorf("STATUS returned %v. %v", resp.Status, string(body))
	}

	var result createResp
	if err = json.Unmarshal(body, &result); err != nil {
		return -1, err
	}
	log.WithField("id", result.ID).Info("Received create response")

	return result.ID, nil
}
