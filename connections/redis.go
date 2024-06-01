package connections

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func ConnectToRedis() {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ctx := context.Background()

	err := RedisClient.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		fmt.Println("ERROR: Failed to set value: 'err := RedisClient.Set(ctx, key, value, 0).Err()'")
		panic(err)
	}

	_, err2 := RedisClient.Get(ctx, "key").Result()
	if err2 != nil {
		fmt.Println("ERROR: Failed to get value: 'val, err := RedisClient.Get(ctx, key).Result()'")
		panic(err2)
	}
	// fmt.Println("key", val)

	fmt.Println("Connected to Redis!")
}

func init() {
	ConnectToRedis()
}
