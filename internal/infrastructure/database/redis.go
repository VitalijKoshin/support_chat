package database

import (
	"context"
	"fmt"

	"supportchat/internal/apperrors"
	"supportchat/internal/config"

	"github.com/go-redis/redis/v8"
)

type RedisDB struct {
	RedisClient *redis.Client
}

func NewRedisDB(cfg config.RedisConfig) (*RedisDB, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Username: cfg.User,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, apperrors.PingRedisError.AppendMessage(err)
	}

	return &RedisDB{
		RedisClient: client,
	}, nil
}
