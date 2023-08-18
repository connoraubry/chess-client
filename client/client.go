package client

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Client struct {
	cfg ClientConfig
}

func New() *Client {

	c := &Client{}
	c.cfg = LoadConfig()
	return c
}

func (c *Client) GetServer() string {
	return fmt.Sprintf("%s:%d", c.cfg.Server, c.cfg.Port)
}

func (c *Client) SetupLogging(logLevel string) {

	var level log.Level

	switch strings.ToLower(logLevel) {
	case "debug":
		level = log.DebugLevel
	case "info":
		level = log.InfoLevel
	case "warning":
		level = log.WarnLevel
	case "error":
		level = log.ErrorLevel
	default:
		level = log.DebugLevel
	}
	log.SetLevel(level)
	log.WithField("level", level).Debug("Logging set up")
}
