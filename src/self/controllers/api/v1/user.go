/**
 * Created by leeezm on 2017/12/30.
 * Email: shiyi@fightcoder.com
 */

package controllers

import (
	"encoding/base64"
	"net/http"
	"self/controllers/baseController"
	"self/managers"

	"github.com/gin-gonic/gin"
)

type User struct {
	baseController.Base
}

func (this *User) Register(routergrp *gin.RouterGroup) {
	routergrp.POST("/login", this.httpHandlerLogin)
	routergrp.POST("/register", this.httpHandlerRegister)
	routergrp.POST("/check", this.httpHandlerCheck)
}

func (this *User) httpHandlerLogin(c *gin.Context) {
	email := this.MustString("email", c)
	password := this.MustString("password", c)
	code, msg := managers.UserManager{}.Login(email, password)

	if code == 1 {
		c.JSON(http.StatusOK, this.Fail(msg))
	} else {
		cookie := &http.Cookie{
			Name:     "token",
			Value:    base64.StdEncoding.EncodeToString([]byte(msg)),
			Path:     "/",
			HttpOnly: true,
		}
		http.SetCookie(c.Writer, cookie)
		c.JSON(http.StatusOK, this.Success("Login successful!"))
	}
}

func (this *User) httpHandlerRegister(c *gin.Context) {
	email := this.MustString("email", c)
	password := this.MustString("password", c)
	flag := managers.UserManager{}.Register(email, password)
	if flag {
		c.JSON(http.StatusOK, this.Fail("注册失败"))
	} else {
		c.JSON(http.StatusOK, this.Success("注册成功"))
	}
}

func (this *User) httpHandlerCheck(c *gin.Context) {
	email := this.MustString("email", c)
	flag := managers.UserManager{}.Check(email)
	if flag == 1 {
		c.JSON(http.StatusOK, this.Fail("try again!"))
	} else if flag == 2 {
		c.JSON(http.StatusOK, this.Fail("false"))
	} else {
		c.JSON(http.StatusOK, this.Success("true"))
	}
}
