package controllers

import (
	"net/http"
	"strconv"

	"self/commons/g"
	"self/controllers/baseController"
	"self/managers"

	"github.com/gin-gonic/gin"
)

type Submit struct {
	baseController.Base
}

func (this *Submit) Register(routergrp *gin.RouterGroup) {
	routergrp.GET("/submission", this.httpHandlerGetsSubmit)
}

func (this *Submit) httpHandlerGetsSubmit(c *gin.Context) {
	cfg := g.Conf()
	problemId, _ := strconv.ParseInt(c.Query("problemId"), 10, 64)
	//获取用户ID(userId)
	var userId int64 = 12
	language := c.Query("language")
	resultDes := c.Query("resultDes")
	requestPage := this.MustInt("requestPage", c)

	submits := managers.SubmitManager{}.GetsSubmit(problemId, userId, language, resultDes, cfg.Show.PageNum, (requestPage-1)*cfg.Show.PageNum)
	num := managers.SubmitManager{}.CountSubmit(problemId, userId, language, resultDes)
	if num%cfg.Show.PageNum != 0 {
		num = num/(cfg.Show.PageNum) + 1
	} else {
		num = num / (cfg.Show.PageNum)
	}
	resp := baseController.PagingResponse{RequestPage: requestPage, TotalPages: num, Data: submits}
	c.JSON(http.StatusOK, this.Success(resp))
}
