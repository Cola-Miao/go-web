package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-web/pkg"
	"go-web/router/message"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, exist := c.Get("name")
		if exist {
			return
		}

		token, err := c.Cookie("token")
		if err != nil {
			err = errors.New("the user is not logged in")
			message.Error(c, http.StatusUnauthorized, err)
			return
		}
		claim, err := pkg.ValidateToken(token)
		if err != nil {
			message.Error(c, http.StatusBadRequest, err)
			return
		}

		c.Set("name", claim.Name)
		c.Next()
	}
}
