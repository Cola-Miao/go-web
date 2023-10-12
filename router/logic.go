package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
