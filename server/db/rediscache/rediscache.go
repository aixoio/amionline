package rediscache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func Connect(addr string) {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	ctx := context.Background()
	client.Set(ctx, "foo", "bar", 0)
}
