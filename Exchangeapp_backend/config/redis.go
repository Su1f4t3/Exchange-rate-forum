package config

import (
	"exchangeapp/global"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     AppConfig.Redis.Addr,
		DB:       0, // 默认数据库(use default DB)
		Password: "",
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		panic("Failed to connect to Redis")
	}

	global.RedisDB = RedisClient
}
