package redis_repo

import (
	"context"
	"fmt"
	"shared/config"
	"time"

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
	if err := rdb.Ping(ctx).Err(); err != nil {
		logrus.Warn(err.Error())
	} else {
		logrus.Info("redis is pinged !")
	}

	return &RedisRepo{
		client:      rdb,
		refreshTime: time.Minute * time.Duration(conf.Redis.RefreshM),
		ctx:         ctx,
	}
}
