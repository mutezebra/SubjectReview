package client

import (
	"context"

	redis "github.com/redis/go-redis/v9"

	"github.com/mutezebra/subject-review/config"
	"github.com/mutezebra/subject-review/pkg/logger"
)

var RedisClient *redis.Client

func initCache() {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Host + ":" + config.Redis.Port,
		Password: config.Redis.Password,
		Network:  config.Redis.Network,
		DB:       config.Redis.Database,
	})

	err := client.Ping(context.Background()).Err()
	if err != nil {
		logger.Fatal(err)
	}
	RedisClient = client
}

func NewRedisClient() *redis.Client {
	if RedisClient == nil {
		initCache()
	}
	return RedisClient
}
