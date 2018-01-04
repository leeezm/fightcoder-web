/**
 * Created by leeezm on 2017/12/30.
 * Email: shiyi@fightcoder.com
 */

package controllers

import (
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
			Value:    msg,
			Path:     "/auth/v1",
			HttpOnly: true,
		}
		http.SetCookie(c.Writer, cookie)
		c.JSON(http.StatusOK, this.Success("Login successful!"))
	}
}

func (this *User) httpHandlerRegister(c *gin.Context) {

}

func (this *User) httpHandlerCheck(c *gin.Context) {

}

func (this *User) httpHandlerCompleteMess(c *gin.Context) {

}
