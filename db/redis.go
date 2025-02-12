package db

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb *redis.Client

func InitializeRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "host.docker.internal:6379", // Redis server address
		Password: "",                          // No password set
		DB:       0,                           // Use default DB
	})

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
		return
	}
	fmt.Println("Connected to Redis:", pong)
}

func SetValue(key string, value string, ttl time.Duration) error {
	err := rdb.Set(ctx, key, value, ttl).Err()
	if err != nil {
		return fmt.Errorf("error setting value: %v", err)
	}
	return nil
}

func GetValue(key string) (string, error) {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return "", fmt.Errorf("error getting value: %v", err)
	}
	return val, nil
}

func DoesKeyExists(key string) (bool, error) {
	exists, err := rdb.Exists(ctx, key).Result()
	if err != nil {
		return false, fmt.Errorf("error checking if key exists: %v", err)
	}
	return exists == 1, nil
}

func ClearAllKeys() error {
	err := rdb.FlushAll(ctx).Err()
	if err != nil {
		return fmt.Errorf("error clearing all keys: %v", err)
	}
	return nil
}
