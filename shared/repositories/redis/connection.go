package redis_repo

import (
	"context"
	"fmt"
	"shared/config"
	"time"

	"github.com/pingcap/log"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type RedisRepo struct {
	client      *redis.Client
	refreshTime time.Duration
	ctx         context.Context
}

func NewRedis(ctx context.Context) *RedisRepo {
	conf := config.Getenv()

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Redis.Host, conf.Redis.Port),
		Password: "",
		DB:       0,
	})
	if rdb == nil {
		log.Error("COULD NOT CONNECT TO REDIS")

	}

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	logrus.Warn("key", val)

	return &RedisRepo{
		client:      rdb,
		refreshTime: time.Minute * time.Duration(conf.Redis.RefreshM),
		ctx:         ctx,
	}
}
