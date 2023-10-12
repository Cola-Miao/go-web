package router

import (
	"github.com/gin-gonic/gin"
	"go-web/config"
)

var engine *gin.Engine

func Init() (err error) {
	if config.Cfg.System.Model != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	engine = gin.New()
	engine.Use(gin.Logger(), gin.Recovery())
	engine.LoadHTMLGlob("template/*.gohtml")

	engine.GET("/health", healthHandler)

	err = engine.Run(config.Cfg.Server.Addr)
	return
}
