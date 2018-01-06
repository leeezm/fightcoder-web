package controllers

import (
	"net/http"
	"self/commons/g"
	"self/managers"
	"strconv"

	//"net/http"
	//"self/commons/g"
	"self/controllers/baseController"
	//"self/managers"
	//"strconv"

	"github.com/gin-gonic/gin"
)

type Submit struct {
	baseController.Base
}

func (this *Submit) Register(routergrp *gin.RouterGroup) {
	routergrp.GET("/submission", this.httpHandlerGetsSubmit)
	routergrp.GET("/submit/id", this.httpHandlerGetSubmitById)
}

func (this *Submit) httpHandlerGetsSubmit(c *gin.Context) {
	cfg := g.Conf()
	problemId := c.Query("problemId")
	language := c.Query("language")
	result, _ := strconv.Atoi(c.Query("result"))
	requestPage := this.MustInt("requestPage", c)

	submits, err := managers.SubmitManager{}.GetsSubmit(problemId, language, result, cfg.Show.PageNum, (requestPage-1)*cfg.Show.PageNum)
	if err {
		c.JSON(http.StatusOK, this.Fail())
	} else {
		num := managers.SubmitManager{}.CountSubmit(problemId, language, result)
		resp := baseController.PagingResponse{RequestPage: requestPage, TotalPages: num, List: submits}
		c.JSON(http.StatusOK, this.Success(resp))
	}
}

func (this *Submit) httpHandlerGetSubmitById(c *gin.Context) {
	submitId := this.MustInt64("id", c)
	var userid int64 = 0
	userId, _ := c.Get("userId")
	if userId != nil {
		if id, ok := userId.(int64); ok {
			userid = id
		}
	}
	submit, flag := managers.SubmitManager{}.GetSubmitById(submitId, userid)
	if !flag {
		c.JSON(http.StatusOK, this.Success(submit))
	} else {
		c.JSON(http.StatusOK, this.Fail())
	}
}
