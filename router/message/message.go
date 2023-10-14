package message

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Error(c *gin.Context, code int, err error) {
	c.JSON(code, gin.H{
		"error": err.Error(),
	})
	c.Abort()
}

func Info(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}
