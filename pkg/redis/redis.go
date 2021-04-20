package rds

import (
	"os"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func Init() {
	var (
		REDIS_HOST = os.Getenv("REDIS_HOST")
		REDIS_PORT = os.Getenv("REDIS_PORT")
		REDIS_PASS = os.Getenv("REDIS_PASS")
	)
	redisClient = redis.NewClient(&redis.Options{
		Addr:     REDIS_HOST + ":" + REDIS_PORT,
		Password: REDIS_PASS,
		DB:       0, // use default DB
	})
}

func RedisManager() *redis.Client {
	return redisClient
}
