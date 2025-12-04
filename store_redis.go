package ratelimiter

import (
	"context"
	"errors"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type RedisStoreClientConfig struct {
	Addr     string
	Password string
	DB       int
}

type RedisStoreClient struct {
	cfg         RedisStoreClientConfig
	redisClient *redis.Client
}

func NewRedisStoreClient(config RedisStoreClientConfig) *RedisStoreClient {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})
	return &RedisStoreClient{cfg: config, redisClient: redisClient}
}

func (redisStoreClient *RedisStoreClient) Get(ctx context.Context, key string) (*int32, error) {
	value, err := redisStoreClient.redisClient.Get(ctx, key).Result()

	if err == redis.Nil {
		return nil, nil
	}

	if err != nil {
		return nil, errors.New("not able to get key from Redis")
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return nil, errors.New("unable to parse value")
	}

	int32Value := int32(intValue)
	return &int32Value, nil
}

func (redisStoreClient *RedisStoreClient) Set(ctx context.Context, key string, value int32) error {
	_, err := redisStoreClient.redisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		errSet := redisStoreClient.redisClient.Set(ctx, key, strconv.Itoa(int(value)), 0).Err()
		if errSet != nil {
			return errors.New("unable to set key on Redis")
		}
	}

	if err != nil {
		return errors.New("unable to get key from Redis")
	}
	return nil
}
