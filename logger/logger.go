package logger

import (
	"go-web/config"
	"log/slog"
	"os"
	"time"
)

func Init() (err error) {
	var fp *os.File
	if config.Cfg.System.Model == "debug" {
		fp = os.Stdin
	} else {
		if fp, err = makeDir(); err != nil {
			return
		}
	}

	handler := slog.NewJSONHandler(fp, &slog.HandlerOptions{
		AddSource:   config.Cfg.Log.AddSource,
		Level:       slog.Level(config.Cfg.Log.Level),
		ReplaceAttr: nil,
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)

	return
}

func makeDir() (fp *os.File, err error) {
	if err = os.MkdirAll(config.Cfg.Log.Path, os.ModePerm); err != nil {
		return
	}
	var (
		date     = time.Now().Format("20060102")
		fileName = "/" + date + ".log"
	)
	fp, err = os.OpenFile(config.Cfg.Log.Path+fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	return
}
