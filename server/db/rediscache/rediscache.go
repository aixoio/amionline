package rediscache

import (
	"github.com/redis/go-redis/v9"
)

func Connect(addr string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	return client
}
