package rediscache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func Connect() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	ctx := context.Background()
	client.Set(ctx, "foo", "bar", 0)
}
