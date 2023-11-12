package main

import (
	"fmt"

	"github.com/aixoio/amionline/server/config"
	"github.com/aixoio/amionline/server/db/mysql"
	"github.com/aixoio/amionline/server/db/rediscache"
	"github.com/aixoio/amionline/server/router"
)

func main() {
	
	config, err := config.LoadConfig("env.json")
	if err != nil {
		fmt.Println("Failed to load config file with error:", err.Error())
		return
	}

	db_connecter, err := mysql.Connect(config.Mysqldb)
	if err != nil {
		fmt.Println("Failed to connect to database with error:", err.Error())
	}
	defer mysql.Disconnect(db_connecter)

	redis_client := rediscache.Connect(config.Redisurl)

	router.Start(db_connecter, redis_client, config)
}
