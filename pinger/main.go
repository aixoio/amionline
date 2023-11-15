package main

import (
	"os"
	"os/signal"

	"github.com/aixoio/amionline/logger"
	"github.com/aixoio/amionline/pinger/config"
	"github.com/aixoio/amionline/pinger/events"
)

func main() {
	logger.Init()

	con, err := config.LoadConfig("env.json")
	if err != nil {
		logger.Error().Println("Failed to load config file with error:", err.Error())
		return
	}

	go events.Start(con)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<- c
		os.Exit(0)
	}()

	open := make(chan bool)

	<- open

}
