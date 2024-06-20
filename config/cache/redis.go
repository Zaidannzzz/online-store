package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

var (
	rdb *redis.Client
	ctx = context.Background()
)

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0, // use default DB
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis: ", err)
	}

	log.Println("Connected to Redis")
}

func Set(key string, value interface{}, expiration time.Duration) error {
	val, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal value: %v", err)
	}

	err = rdb.Set(ctx, key, val, expiration).Err()
	if err != nil {
		return fmt.Errorf("failed to set key %s in Redis: %v", key, err)
	}

	return nil
}

func Get(key string, dest interface{}) error {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return fmt.Errorf("failed to get key %s from Redis: %v", key, err)
	}

	err = json.Unmarshal([]byte(val), &dest)
	if err != nil {
		return fmt.Errorf("failed to unmarshal value: %v", err)
	}

	return nil
}
