package rediscache

import (
	"github.com/aixoio/amionline/logger"
	"github.com/redis/go-redis/v9"
)

func Connect(addr string) *redis.Client {
	logger.Info().Println("Connecting to Redis")
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	return client
}
