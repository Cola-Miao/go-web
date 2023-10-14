package redis

import (
	"github.com/redis/go-redis/v9"
	"go-web/config"
)

var rds *redis.Client

func Init() {
	rds = redis.NewClient(redisConfig())
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
