/**
 * Created by leeezm on 2017/12/13.
 * Email: shiyi@fightcoder.com
 */

package controllers

import (
	"net/http"
	"strconv"
	"time"

	"self/commons/g"
	"self/controllers/baseController"
	"self/managers"

	"github.com/gin-gonic/gin"
)

type Submit struct {
	baseController.Base
}

func (this *Submit) Register(routergrp *gin.RouterGroup) {
	routergrp.POST("/submit", this.httpHandlerAddSubmit)
	routergrp.GET("/submit/id", this.httpHandlerGetSubmitById)
}

func (this *Submit) httpHandlerAddSubmit(c *gin.Context) {
	language := this.MustString("language", c)
	problemId := this.MustInt64("problemId", c)
	var submitTime int64 = time.Now().UnixNano() / 1000000
	code := this.MustString("code", c)

	userId, _ := c.Get("userId")
	if id, ok := userId.(int64); ok {
		flag := managers.SubmitManager{}.AddSubmit(problemId, id, language, submitTime, code)
		if !flag {
			c.JSON(http.StatusOK, this.Success("提交成功!"))
		}
	} else {
		c.JSON(http.StatusOK, this.Fail("提交失败!"))
	}
}

func (this *Submit) httpHandlerGetSubmitById(c *gin.Context) {
	submitId := this.MustInt64("id", c)
	userId, _ := c.Get("userId")
	if id, ok := userId.(int64); ok {
		submit, flag := managers.SubmitManager{}.GetSubmitById(submitId, id)
		if !flag {
			c.JSON(http.StatusOK, this.Success(submit))
		}
	} else {
		c.JSON(http.StatusOK, this.Fail())
	}
}
