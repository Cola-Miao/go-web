package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go-web/model"
	"log/slog"
)

var Cfg model.Config

func Init() (err error) {
	vp := viper.New()
	vp.SetConfigFile("config.yaml")
	if err = vp.ReadInConfig(); err != nil {
		return
	}
	if err = vp.Unmarshal(&Cfg); err != nil {
		return
	}

	vp.WatchConfig()
	vp.OnConfigChange(func(in fsnotify.Event) {
		slog.Info("config changed", in.Op, in.Name)
		if err = vp.Unmarshal(&Cfg); err != nil {
			slog.Error("load config failed", "error", err.Error())
		}
	})

	return
}
