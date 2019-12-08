package redis

import (
	"github.com/go-redis/redis"
)


func Connect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})


	return client
}
