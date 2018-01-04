/**
 * Created by leeezm on 2017/12/14.
 * Email: shiyi@fightcoder.com
 */

package controllers

import (
	"net/http"
	"self/controllers/baseController"
	"self/models"

	"github.com/gin-gonic/gin"
)

type User struct {
	baseController.Base
}

func (this *User) Register(routergrp *gin.RouterGroup) {
	routergrp.GET("/currentUser", this.httpHandlerCurrentUser)
	routergrp.GET("/quit", this.httpHandlerQuit)
}

func (this *User) httpHandlerCurrentUser(c *gin.Context) {
	userId, _ := c.Get("userId")
	var user *models.User
	var err error
	if id, ok := userId.(int64); ok {
		user, err = models.User{}.GetById(id)
	} else {
		c.JSON(http.StatusOK, this.Fail("获取失败!"))
	}
	if err != nil {
		c.JSON(http.StatusOK, this.Fail("获取失败!"))
	} else {
		c.JSON(http.StatusOK, this.Success(user))
	}
}

func (this *User) httpHandlerQuit(c *gin.Context) {
	cookie := http.Cookie{Name: "token", Path: "/", MaxAge: -1}
	http.SetCookie(c.Writer, &cookie)
	c.JSON(http.StatusOK, this.Success("退出成功!"))
}

//func (this *User) httpHandlerUpload(c *gin.Context) {
//	_, _, err := c.Request.FormFile("upload")
//	if err != nil {
//		panic(err)
//	}
//
//	//managers.BaseManager{}.SaveImage(file)
//
//	c.JSON(http.StatusOK, this.Success())
//}
