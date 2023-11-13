package main

import "github.com/aixoio/amionline/logger"

func main() {
	logger.Init()

	logger.Info().Println("Hello world!")
}
