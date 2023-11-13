package main

import (
	"github.com/aixoio/amionline/logger"
	"github.com/aixoio/amionline/pinger/config"
)

func main() {
	logger.Init()

	con, err := config.LoadConfig("env.json")
	if err != nil {
		logger.Error().Println("Failed to load config file with error:", err.Error())
		return
	}

	logger.Info().Println(con)
	
	logger.Info().Println("Hello world!")
}
