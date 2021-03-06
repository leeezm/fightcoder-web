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
	"self/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	baseController.Base
}

func (this *User) Register(routergrp *gin.RouterGroup) {
	routergrp.POST("/login", this.httpHandlerLogin)
	routergrp.POST("/register", this.httpHandlerRegister)
	routergrp.POST("/check", this.httpHandlerCheck)
	routergrp.GET("/usermess", this.httpHandlerUserMess)
}

func (this *User) httpHandlerUserMess(c *gin.Context) {
	userIdString := c.Query("userId")
	if userIdString == "" {
		c.JSON(http.StatusOK, this.Fail("参数不能为空!"))
	}
	userId, err := strconv.ParseInt(userIdString, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, this.Fail("参数有误!"))
	}
	user, err := models.User{}.GetById(userId)
	if err != nil {
		c.JSON(http.StatusOK, this.Fail("请重新尝试!"))
	}
	c.JSON(http.StatusOK, this.Success(user))
}

func (this *User) httpHandlerLogin(c *gin.Context) {
	loginType := this.MustString("type", c)
	var param1, param2 string
	if loginType == "qq" {
		param1 = this.MustString("code", c)
		param2 = this.MustString("state", c)
	} else if loginType == "simple" {
		param1 = this.MustString("email", c)
		param2 = this.MustString("password", c)
	} else {
		c.JSON(http.StatusOK, this.Fail("参数错误!"))
	}
	state, msg, userId := managers.AccountManager{}.Login(param1, param2, loginType)

	if state == managers.EMAIL_NOT_EXIT || state == managers.PASSWORD_IS_WRONG || state == managers.PARAM_IS_WRONG {
		var msg string
		switch state {
		case managers.EMAIL_NOT_EXIT:
			msg = "Email not exit!"
			break
		case managers.PASSWORD_IS_WRONG:
			msg = "Password is wrong!"
			break
		case managers.PARAM_IS_WRONG:
			msg = "Param is wrong!"
			break
		}
		c.JSON(http.StatusOK, this.Fail(msg))
	} else {
		cookie := &http.Cookie{
			Name:     "token",
			Value:    base64.StdEncoding.EncodeToString([]byte(msg)),
			Path:     "/",
			HttpOnly: true,
		}
		http.SetCookie(c.Writer, cookie)
		result := make(map[string]string)
		if state == managers.FIRST_LOGIN {
			result["is_first"] = "true"
			result["user_id"] = strconv.FormatInt(userId, 10)
		} else {
			result["is_first"] = "false"
			result["user_id"] = strconv.FormatInt(userId, 10)
		}
		c.JSON(http.StatusOK, this.Success(result))
	}
}

func (this *User) httpHandlerRegister(c *gin.Context) {
	email := this.MustString("email", c)
	password := this.MustString("password", c)
	userId := managers.AccountManager{}.Register(email, password)
	if userId > 0 {
		c.JSON(http.StatusOK, this.Success(userId))
	} else {
		c.JSON(http.StatusOK, this.Fail())
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
