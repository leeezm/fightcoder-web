package controllers

import (
	"self/commons/g"
	"self/controllers/baseController"

	models "self/models"

	"net/http"
	"self/managers"

	"github.com/gin-gonic/gin"
)

//会先进行登录验证
type Problem struct {
	baseController.Base
}

func (this *Problem) Register(routergrp *gin.RouterGroup) {
	routergrp.POST("/save", this.httpHandlerSaveCode)
	routergrp.POST("/problem", this.httpHandlerAddProblem)
	routergrp.DELETE("/problem", this.httpHandlerDeleteProblem)
	routergrp.PUT("/problem", this.httpHandlerUpdateProblem)

	routergrp.GET("/problem/upload", this.httpHandlerGetsProblemUser)
	routergrp.GET("/problem/apply", this.httpHandlerAddProblemCheck)
	routergrp.GET("/problem/getCode", this.httpHandlerGetCode)
}

func (this *Problem) httpHandlerSaveCode(c *gin.Context) {
	var saveCode models.SaveCode
	if err := c.BindJSON(&saveCode); err != nil {
		panic(err)
	}
	managers.ProblemManager{}.SaveCode(saveCode.ProblemId, saveCode.UserId, saveCode.Code)
	c.JSON(http.StatusOK, this.Success())
}

func (this *Problem) httpHandlerGetCode(c *gin.Context) {
	problemId := this.MustInt64("id", c)
	var userId int64 = 2
	code := managers.ProblemManager{}.GetCode(problemId, userId)
	c.JSON(http.StatusOK, this.Success(code))
}

func (this *Problem) httpHandlerAddProblem(c *gin.Context) {
	var problem ProblemRequest
	if err := c.BindJSON(&problem); err != nil {
		panic(err)
	}
	managers.ProblemManager{}.SubmitByUser(&problem.Data)
	c.JSON(http.StatusOK, this.Success())
}

func (this *Problem) httpHandlerDeleteProblem(c *gin.Context) {
	var problem ProblemRequest
	if err := c.BindJSON(&problem); err != nil {
		panic(err)
	}
	managers.ProblemManager{}.RemoveProblemUser(problem.Data.Id)
	c.JSON(http.StatusOK, this.Success())
}

func (this *Problem) httpHandlerUpdateProblem(c *gin.Context) {
	var problem ProblemRequest
	if err := c.BindJSON(&problem); err != nil {
		panic(err)
	}

	managers.ProblemManager{}.UpdateByUser(&problem.Data)
	c.JSON(http.StatusOK, this.Success())
}

func (this *Problem) httpHandlerGetsProblemUser(c *gin.Context) {
	cfg := g.Conf()
	requestPage := this.MustInt("requestPage", c)
	id := this.MustInt64("id", c)

	problems := managers.ProblemManager{}.GetsProblemUser(id, cfg.Show.PageNum, (requestPage-1)*cfg.Show.PageNum)
	num := managers.ProblemManager{}.CountProblemUser(id)
	if num%cfg.Show.PageNum != 0 {
		num = num/(cfg.Show.PageNum) + 1
	} else {
		num = num / (cfg.Show.PageNum)
	}
	resp := baseController.PagingResponse{RequestPage: requestPage, TotalPages: num, Data: problems}
	c.JSON(http.StatusOK, this.Success(resp))
}

func (this *Problem) httpHandlerAddProblemCheck(c *gin.Context) {
	problemId := this.MustInt64("id", c)
	managers.ProblemManager{}.AddProblemCheck(problemId)
	c.JSON(http.StatusOK, this.Success())
}

type ProblemRequest struct {
	UserId int64              `json:"userId"`
	Token  string             `json:"token"`
	Data   models.ProblemUser `json:"data"`
}
