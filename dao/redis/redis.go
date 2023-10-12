package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go-web/config"
)

var rds *redis.Client

func Init() (err error) {
	ctx := context.Background()

	rds = redis.NewClient(redisConfig())
	err = rds.Ping(ctx).Err()

	return
}

func redisConfig() (redisOp *redis.Options) {
	redisOp = &redis.Options{
		Addr:     config.Cfg.Redis.Addr,
		Password: config.Cfg.Redis.Password,
		DB:       config.Cfg.Redis.DB,
	}

	return
}
