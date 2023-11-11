package main

import (
	"fmt"

	"github.com/aixoio/amionline/server/config"
)

func main() {
	
	config, err := config.LoadConfig("env.json")
	if err != nil {
		fmt.Println("Failed to load config file with error:", err.Error())
		return
	}

}
