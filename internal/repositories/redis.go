package repositories

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

type RedisConfig struct {
	Host     string
	Port     string
	Username string
	Password string
}

func NewRedisDB(ctx context.Context) (*redis.Client, error) {
	rOptions, err := newRedisConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load redis option: %w", err)
	}

	rdb := redis.NewClient(rOptions)

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}
	return rdb, nil
}

func newRedisConfig() (*redis.Options, error) {
	err := godotenv.Load(".env.local")
	if err != nil {
		return nil, fmt.Errorf("failed to load env file: %w", err)
	}

	host := os.Getenv("REDIS_HOST")
	if host == "" {
		return nil, errors.New("empty REDIS_HOST env")
	}

	port := os.Getenv("REDIS_PORT")
	if port == "" {
		return nil, errors.New("empty REDIS_PORT")
	}

	user := os.Getenv("REDIS_USER")
	if user == "" {
		return nil, errors.New("empty REDIS_USER")
	}

	passwd := os.Getenv("REDIS_USER_PASSWORD")
	if passwd == "" {
		return nil, errors.New("empty REDIS_USER_PASSWORD")
	}

	rOptions := &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Username: user,
		Password: passwd,
	}

	return rOptions, nil

}
