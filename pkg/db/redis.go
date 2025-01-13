package db

import (
	redis "github.com/redis/go-redis/v9"
)

type RedisStruct struct {
	client *redis.Client
}

func NewRedisClient() *RedisStruct {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // No password set
		DB:       0,  // Use default DB
		Protocol: 2,  // Connection protocol
	})

	return &RedisStruct{
		client: client,
	}
}
