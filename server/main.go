package main

import (
	"fmt"

	"github.com/aixoio/amionline/server/config"
	"github.com/aixoio/amionline/server/db"
	"github.com/aixoio/amionline/server/router"
)

func main() {
	
	config, err := config.LoadConfig("env.json")
	if err != nil {
		fmt.Println("Failed to load config file with error:", err.Error())
		return
	}

	db_connecter, err := db.Connect(config.Db)
	if err != nil {
		fmt.Println("Failed to connect to database with error:", err.Error())
	}
	defer db.Disconnect(db_connecter)

	router.Start(db_connecter)
}
