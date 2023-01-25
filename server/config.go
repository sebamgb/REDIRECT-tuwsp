package server

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	port        string
	dataBaseUrl string
	driverDb    string
	Debug       bool
}

// NewConfig builder of Config
func NewConfig(DRIVER_DB, SERVER_DB, USER_DB, PASSWORD_DB, PORT_DB, DATABASE, DEBUG string) (*Config, error) {
	query := url.Values{}
	query.Add("database", DATABASE)
	// DEBUG to bool
	debug, err := strconv.ParseBool(DEBUG)
	if err != nil {
		panic(err)
	}
	// making url to db
	u := &url.URL{
		Scheme:   DRIVER_DB,
		User:     url.UserPassword(USER_DB, PASSWORD_DB),
		Host:     fmt.Sprintf("%s:%s", SERVER_DB, PORT_DB),
		RawQuery: query.Encode(),
	}
	// get port from .env
	Port := os.Getenv("PORT_APP")
	// with : to the beggining
	prefix := strings.Index(Port, ":")
	sb := strings.Builder{}
	sb.WriteString(":")
	sb.WriteString(Port)
	// adding : if without
	if prefix != 0 {
		Port = sb.String()
	}
	return &Config{
		dataBaseUrl: u.String(),
		port:        Port,
		driverDb:    DRIVER_DB,
		Debug:       debug,
	}, nil
}
