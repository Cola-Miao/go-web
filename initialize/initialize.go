package initialize

import (
	"go-web/config"
	"go-web/dao/mysql"
	"go-web/dao/redis"
	"go-web/logger"
	"go-web/router"
)

func Default() (err error) {
	if err = config.Init(); err != nil {
		return
	}
	if err = logger.Init(); err != nil {
		return
	}
	if err = mysql.Init(); err != nil {
		return
	}
	if err = redis.Init(); err != nil {
		return
	}
	if err = router.Init(); err != nil {
		return
	}

	return
}
