package main

import (
	"github.com/aixoio/amionline/server/config"
	"github.com/aixoio/amionline/server/db/mysql"
	"github.com/aixoio/amionline/server/db/rediscache"
	"github.com/aixoio/amionline/logger"
	"github.com/aixoio/amionline/server/router"
)

func main() {

	logger.Init()
	
	config, err := config.LoadConfig("env.json")
	if err != nil {
		logger.Error().Println("Failed to load config file with error:", err.Error())
		return
	}

	db_connecter, err := mysql.Connect(config.Mysqldb)
	if err != nil {
		logger.Error().Println("Failed to connect to database with error:", err.Error())
	}
	defer mysql.Disconnect(db_connecter)

	redis_client := rediscache.Connect(config.Redisurl)

	router.Start(db_connecter, redis_client, config)
}
