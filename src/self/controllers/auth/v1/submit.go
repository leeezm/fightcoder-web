/**
 * Created by leeezm on 2017/12/13.
 * Email: shiyi@fightcoder.com
 */

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
	routergrp.POST("/submit", this.httpHandlerAddSubmit)
	routergrp.GET("/submit/id", this.httpHandlerGetSubmitById)
	routergrp.GET("/submit/mess", this.httpHandlerGetsSubmit)

	routergrp.POST("/submitContest", this.httpHandlerAddSubmitContest)
	routergrp.GET("/submitContest/id", this.httpHandlerGetSubmitContestById)
	routergrp.GET("/submitContest/mess", this.httpHandlerGetsSubmitContest)

	routergrp.POST("/submitUser", this.httpHandlerAddSubmitUser)
	routergrp.GET("/submitUser/id", this.httpHandlerGetSubmitUserById)
	routergrp.GET("/submitUser/mess", this.httpHandlerGetsSubmitUser)
}

func (this *Submit) httpHandlerAddSubmit(c *gin.Context) {
	problemId := this.MustInt64("problemId", c)
	//鉴权等待实现
	var userId int64 = 12
	language := this.MustString("language", c)
	var submitTime int64 = 1234
	code := this.MustString("code", c)

	managers.SubmitManager{}.AddSubmit(problemId, userId, language, submitTime, code)
	managers.ProblemManager{}.SaveCode(problemId, userId, code)
	c.JSON(http.StatusOK, this.Success())
}

func (this *Submit) httpHandlerGetSubmitById(c *gin.Context) {
	id := this.MustInt64("Id", c)
	submit := managers.SubmitManager{}.GetSubmitById(id)
	c.JSON(http.StatusOK, this.Success(submit))
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

func (this *Submit) httpHandlerAddSubmitContest(c *gin.Context) {
	problemId := this.MustInt64("problemId", c)
	//鉴权等待实现
	var userId int64 = 12
	language := this.MustString("language", c)
	submitTime := this.MustInt64("submitTime", c)
	code := this.MustString("code", c)
	contestId := this.MustInt64("contestId", c)
	managers.SubmitManager{}.AddSubmitContest(problemId, userId, language, submitTime, code, contestId)
	c.JSON(http.StatusOK, this.Success())
}

func (this *Submit) httpHandlerGetSubmitContestById(c *gin.Context) {
	id := this.MustInt64("Id", c)
	submitContest := managers.SubmitManager{}.GetSubmitContestById(id)
	c.JSON(http.StatusOK, this.Success(submitContest))
}

func (this *Submit) httpHandlerGetsSubmitContest(c *gin.Context) {
	cfg := g.Conf()
	contestId := this.MustInt64("contestId", c)
	problemId := this.MustInt64("problemId", c)
	//获取用户ID(userId)
	userId := this.MustInt64("id", c)
	language := this.MustString("language", c)
	resultDes := this.MustString("resultDes", c)
	requestPage := this.MustInt("requestPage", c)

	submitContests := managers.SubmitManager{}.GetsSubmitContest(contestId, problemId, userId, language, resultDes, cfg.Show.PageNum, (requestPage-1)*cfg.Show.PageNum)
	num := managers.SubmitManager{}.CountSubmitContest(contestId, problemId, userId, language, resultDes)
	if num%cfg.Show.PageNum != 0 {
		num = num/(cfg.Show.PageNum) + 1
	} else {
		num = num / (cfg.Show.PageNum)
	}
	resp := baseController.PagingResponse{RequestPage: requestPage, TotalPages: num, Data: submitContests}
	c.JSON(http.StatusOK, this.Success(resp))
}

func (this *Submit) httpHandlerAddSubmitUser(c *gin.Context) {
	problemId := this.MustInt64("problemId", c)
	//鉴权等待实现
	var userId int64 = 12
	language := this.MustString("language", c)
	submitTime := this.MustInt64("submitTime", c)
	code := this.MustString("code", c)

	managers.SubmitManager{}.AddSubmitUser(problemId, userId, language, submitTime, code)
	c.JSON(http.StatusOK, this.Success())
}

func (this *Submit) httpHandlerGetSubmitUserById(c *gin.Context) {
	id := this.MustInt64("Id", c)
	submitUser := managers.SubmitManager{}.GetSubmitUserById(id)
	c.JSON(http.StatusOK, this.Success(submitUser))
}

func (this *Submit) httpHandlerGetsSubmitUser(c *gin.Context) {
	cfg := g.Conf()
	problemId := this.MustInt64("problemId", c)
	//获取用户ID(userId)
	userId := this.MustInt64("id", c)
	language := this.MustString("language", c)
	resultDes := this.MustString("resultDes", c)
	requestPage := this.MustInt("requestPage", c)

	submitUsers := managers.SubmitManager{}.GetsSubmitUser(problemId, userId, language, resultDes, cfg.Show.PageNum, (requestPage-1)*cfg.Show.PageNum)
	num := managers.SubmitManager{}.CountSubmitUser(problemId, userId, language, resultDes)
	if num%cfg.Show.PageNum != 0 {
		num = num/(cfg.Show.PageNum) + 1
	} else {
		num = num / (cfg.Show.PageNum)
	}
	resp := baseController.PagingResponse{RequestPage: requestPage, TotalPages: num, Data: submitUsers}
	c.JSON(http.StatusOK, this.Success(resp))
}
