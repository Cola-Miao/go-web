package router

import (
	"github.com/gin-gonic/gin"
	"go-web/dao/mysql"
	"go-web/logic"
	"go-web/model"
	"go-web/pkg"
	"go-web/router/message"
	"net/http"
)

func healthHandler(c *gin.Context) {
	message.Info(c, "ok")
}

func authHandler(c *gin.Context) {
	name := c.GetString("name")
	message.Info(c, "hello,"+name)
}

func signupHandler(c *gin.Context) {
	var sf model.SignupForm
	var err error

	if err = c.ShouldBind(&sf); err != nil {
		message.Error(c, http.StatusBadRequest, err)
		return
	}
	if err = mysql.CreateUser(&sf); err != nil {
		message.Error(c, http.StatusBadRequest, err)
		return
	}

	message.Info(c, "signup successful")
}

func loginHandler(c *gin.Context) {
	var lf model.LoginForm
	var err error
	if err = c.ShouldBind(&lf); err != nil {
		message.Error(c, http.StatusBadRequest, err)
		return
	}
	if err = logic.LoginAuth(&lf); err != nil {
		message.Error(c, http.StatusBadRequest, err)
		return
	}

	var token string
	if token, err = pkg.GenToken(&lf); err != nil {
		message.Error(c, http.StatusServiceUnavailable, err)
		return
	}
	c.SetCookie("token", token, 0, "/", "127.0.0.1", false, true)
	message.Info(c, "login successful")
}
