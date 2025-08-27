package cache

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hanzala211/instagram/utils"
	"github.com/redis/go-redis/v9"
)

type RedisRepo struct {
	Client *redis.Client
}

func NewRedisClient() *RedisRepo {
	redisURL := utils.GetEnv("REDIS_URL", "redis://localhost:6379")
	log.Printf("Attempting to connect to Redis at: %s", redisURL)
	
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Printf("Failed to parse Redis URL: %v", err)
		panic(fmt.Sprintf("Invalid Redis URL: %v", err))
	}

	opt.DialTimeout = 10 * time.Second
	opt.ReadTimeout = 30 * time.Second
	opt.WriteTimeout = 30 * time.Second
	opt.PoolTimeout = 30 * time.Second
	opt.MaxRetries = 3

	client := redis.NewClient(opt)
	
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	
	_, err = client.Ping(ctx).Result()
	if err != nil {
		log.Printf("Failed to connect to Redis: %v", err)
		panic(fmt.Sprintf("Cannot connect to Redis: %v", err))
	}
	
	log.Println("Connection to Redis successful")
	return &RedisRepo{Client: client}
}

func (r *RedisRepo) Set(key string, value interface{}, expiration time.Duration) error {
	return r.Client.Set(context.Background(), key, value, expiration).Err()
}

func (r *RedisRepo) Get(key string) (string, error) {
	str, err := r.Client.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return str, nil
}

func (r *RedisRepo) Delete(key string) error {
	return r.Client.Del(context.Background(), key).Err()
}