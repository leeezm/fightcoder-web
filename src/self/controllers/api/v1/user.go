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
	routergrp.POST("/completeMess", this.httpHandlerCompleteMess)

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
		c.JSON(http.StatusOK, this.Success("false"))
	} else {
		c.JSON(http.StatusOK, this.Success("true"))
	}
}

func (this *User) httpHandlerCompleteMess(c *gin.Context) {
	nickName := this.MustString("nickName", c)
	description := this.MustString("description", c)
	sex := this.MustInt("sex", c)
	birthday := this.MustString("birthday", c)
	dailyAddress := this.MustString("dailyAddress", c)
	recvAddress := this.MustString("recvAddress", c)
	tShirtSize := this.MustString("tShirtSize", c)
	statSchool := this.MustInt("statSchool", c)
	blog := this.MustString("blog", c)
	git := this.MustString("git", c)
	userId, _ := c.Get("userId")
	if id, ok := userId.(int64); ok {
		err := managers.UserManager{}.CompleteMess(id, nickName, description, sex, birthday, dailyAddress, recvAddress, tShirtSize, statSchool, blog, git)
		if err == nil {
			c.JSON(http.StatusOK, this.Success("更新成功!"))
		}
	}
	c.JSON(http.StatusOK, this.Fail("更新失败!"))
}
