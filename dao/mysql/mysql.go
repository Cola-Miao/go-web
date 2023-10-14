package mysql

import (
	"fmt"
	"go-web/config"
	"go-web/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mySQL *gorm.DB
var migrateList = []any{&model.User{}}

func Init() (err error) {
	dsn := formatDsn()
	if mySQL, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		return
	}
	err = mySQL.AutoMigrate(migrateList...)

	return
}

func formatDsn() (dsn string) {
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Cfg.MySQL.User,
		config.Cfg.MySQL.Pass,
		config.Cfg.MySQL.IP,
		config.Cfg.MySQL.Port,
		config.Cfg.MySQL.DBName,
	)

	return
}
