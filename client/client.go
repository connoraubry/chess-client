package client

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	cfg          ClientConfig
	colorHandler color.Attribute
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

func (c *Client) SetupColor() bool {
	switch c.cfg.Color {
	case "red":
		c.colorHandler = color.FgRed
	case "yellow":
		c.colorHandler = color.FgYellow
	case "blue":
		c.colorHandler = color.FgBlue
	case "green":
		c.colorHandler = color.FgGreen
	case "cyan":
		c.colorHandler = color.FgCyan
	case "magenta":
		c.colorHandler = color.FgMagenta
	case "white":
		c.colorHandler = color.FgWhite
	default:
		c.colorHandler = color.FgWhite
		return false
	}
	return true
}
