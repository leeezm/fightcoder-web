/**
 * Created by leeezm on 2017/12/13.
 * Email: shiyi@fightcoder.com
 */

package controllers

import (
	"net/http"
	"self/controllers/baseController"
	"self/managers"
	"time"

	"github.com/gin-gonic/gin"
)

type Submit struct {
	baseController.Base
}

func (this *Submit) Register(routergrp *gin.RouterGroup) {
	routergrp.POST("/submit", this.httpHandlerAddSubmit)
}

func (this *Submit) httpHandlerAddSubmit(c *gin.Context) {
	language := this.MustString("language", c)
	problemId := this.MustInt64("problemId", c)
	var submitTime int64 = time.Now().Unix()
	code := this.MustString("code", c)

	userId, _ := c.Get("userId")
	if id, ok := userId.(int64); ok {
		submit, flag := managers.SubmitManager{}.AddSubmit(problemId, id, language, submitTime, code)
		if !flag {
			c.JSON(http.StatusOK, this.Success(submit.Id))
		}
	} else {
		c.JSON(http.StatusOK, this.Fail("提交失败!"))
	}
}
