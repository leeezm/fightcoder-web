/**
 * Created by leeezm on 2017/12/14.
 * Email: shiyi@fightcoder.com
 */

package controllers

import (
	"net/http"
	"self/controllers/baseController"

	"github.com/gin-gonic/gin"
)

type User struct {
	baseController.Base
}

func (this *User) Register(routergrp *gin.RouterGroup) {
	routergrp.POST("/upload", this.httpHandlerUpload)
}

func (this *User) httpHandlerUpload(c *gin.Context) {
	_, _, err := c.Request.FormFile("upload")
	if err != nil {
		panic(err)
	}

	//managers.BaseManager{}.SaveImage(file)

	c.JSON(http.StatusOK, this.Success())
}
