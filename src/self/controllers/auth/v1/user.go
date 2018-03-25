/**
 * Created by leeezm on 2017/12/14.
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
	routergrp.PUT("/completeMess", this.httpHandlerCompleteMess)
	routergrp.POST("/uploadImage", this.httpHandlerUploadImage)
	routergrp.GET("/quit", this.httpHandlerQuit)
}

func (this *User) httpHandlerUploadImage(c *gin.Context) {
	picType := this.MustString("picType", c)
	file, _, _ := c.Request.FormFile("upload")
	userId, _ := c.Get("userId")
	if id, ok := userId.(int64); ok {
		flag := managers.UserManager{}.UploadImage(file, id, picType)
		if !flag {
			c.JSON(http.StatusOK, this.Success("上传成功!"))
		}
	} else {
		c.JSON(http.StatusOK, this.Fail("上传失败!"))
	}
}

func (this *User) httpHandlerQuit(c *gin.Context) {
	cookie := http.Cookie{Name: "token", Path: "/", MaxAge: -1}
	http.SetCookie(c.Writer, &cookie)
	c.JSON(http.StatusOK, this.Success("退出成功!"))
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
