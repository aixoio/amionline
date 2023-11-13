package main

import "github.com/aixoio/amionline/pinger/logger"

func main() {
	logger.Init()

	logger.Info().Println("Hello world!")
}
