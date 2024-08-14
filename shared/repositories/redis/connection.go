package redis_repo

import (
	"fmt"
	"shared/config"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisRepo struct {
	client      *redis.Client
	refreshTime time.Duration
}

func NewRedis() *RedisRepo {
	conf := config.Getenv()

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.Redis.Host, conf.Redis.Port),
		Password: "",
		DB:       0,
	})

	return &RedisRepo{
		client:      rdb,
		refreshTime: time.Minute * time.Duration(conf.Redis.RefreshM),
	}
}
