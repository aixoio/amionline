package main

import (
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

	events.Start(con)

}
