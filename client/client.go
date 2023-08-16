package client

import (
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

func (c *Client) SetupLogging(logLevel string) {

	//log.SetFormatter(&log.TextFormatter{
	//	FullTimestamp:   true,
	//	TimestampFormat: "2006-01-02 15:04:05",
	//})

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
	log.WithField("level", level).Info("Logging set up")
}
