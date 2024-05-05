package server

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func ConnectRedis() (*redis.Client, error) {

	rd := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	if _, err := rd.Ping(ctx).Result(); err != nil {
		return nil, fmt.Errorf("erro ao conectar ao Redis: %v", err)
	}

	return rd, nil
}
