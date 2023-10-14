package router

import (
	"github.com/gin-gonic/gin"
	"go-web/config"
	"go-web/router/middleware"
)

var engine *gin.Engine

func Init() (err error) {
	if config.Cfg.System.Model != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	engine = gin.New()
	engine.Use(gin.Logger(), gin.Recovery())
	engine.LoadHTMLGlob("template/*.gohtml")

	public := engine.Group("/")
	{
		public.GET("/health", healthHandler)
		public.POST("/signup", signupHandler)
		public.POST("/login", loginHandler)
	}

	private := engine.Group("/", middleware.Auth())
	{
		private.GET("/auth", authHandler)
	}

	err = engine.Run(config.Cfg.Server.Addr)
	return
}
