package main

import (
	"fmt"
	"go-web/config"
	"go-web/initialize"
	"log/slog"
)

func main() {
	var err error
	if err = initialize.Default(); err != nil {
		slog.Error("initialize failed", "error", err.Error())
		panic(err)
	}

	fmt.Println(config.Cfg)
}
