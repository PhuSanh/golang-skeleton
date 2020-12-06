package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"golang-skeleton/config"
	"log"
	"sync"
	"time"
)

type IService interface {
	Get(ctx context.Context, key string) (*redis.StringCmd, error)
	Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error
}

type service struct {
	client *redis.Client
}

var (
	mutex   sync.RWMutex
	rClient *redis.Client
	once    sync.Once
)

func InitRedis(cfg *config.Config) {
	mutex.Lock()
	defer mutex.Unlock()
	rClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password, // no password set
		DB:       0,                  // use default DB
		PoolSize: cfg.Redis.Pool,
	})
	pong, err := rClient.Ping(context.Background()).Result()
	if err != nil {
		log.Printf("Redis connect err: %v\n", err)
		return
	}
	log.Printf("[Success] Redis [%s:%s] connected with poolSize %d: %v\n", cfg.Redis.Host, cfg.Redis.Port, cfg.Redis.Pool, pong)
}

func NewRedisService(cfg *config.Config) IService {
	once.Do(func() {
		InitRedis(cfg)
	})
	return &service{client: rClient}
}

func (s *service) Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error {
	return s.client.Set(ctx, key, val, expiration).Err()
}

func (s *service) Get(ctx context.Context, key string) (*redis.StringCmd, error) {
	existed, err := s.client.Exists(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	if existed == 0 {
		return nil, nil
	}
	return s.client.Get(ctx, key), nil
}
