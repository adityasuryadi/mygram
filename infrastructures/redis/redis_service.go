package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	// config "mygram/infrastructures"
)

var ctx = context.Background()

func NewRedisClient() *redis.Client {
	// host := configuration.Get("REDIS_HOST")
	// password := configuration.Get("REDIS_PASSWORD")
	host := "localhost:6379"
	password := "password"
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0,
	})
	return client
}

func Set(key string, value interface{}, ttl int) error {
	rdb := NewRedisClient()
	tme := time.Duration(ttl) * time.Second
	err := rdb.Set(ctx, key, value, tme).Err()
	if err != nil {
		return err
	}
	return nil
}

func Get(key string) (interface{}, error) {
	rdb := NewRedisClient()
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return val, nil
}
