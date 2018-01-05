package controllers

import (
	"self/commons/g"
	"self/controllers/baseController"

	models "self/models"

	"net/http"
	"self/managers"

	"github.com/gin-gonic/gin"
)

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
	problemId := this.MustInt64("problemId", c)
	code := this.MustString("code", c)
	userId, _ := c.Get("userId")
	if id, ok := userId.(int64); ok {
		managers.ProblemManager{}.SaveCode(problemId, id, code)
		c.JSON(http.StatusOK, this.Success("保存成功"))
	} else {
		c.JSON(http.StatusOK, this.Fail("保存失败"))
	}

}

func (this *Problem) httpHandlerGetCode(c *gin.Context) {
	problemId := this.MustInt64("id", c)
	userId, _ := c.Get("userId")
	if id, ok := userId.(int64); ok {
		code := managers.ProblemManager{}.GetCode(problemId, id)
		c.JSON(http.StatusOK, this.Success(code))
	} else {
		c.JSON(http.StatusOK, this.Fail("保存失败"))
	}
}

func (this *Problem) httpHandlerAddProblem(c *gin.Context) {
	var problem models.ProblemUser
	if err := c.BindJSON(&problem); err != nil {
		panic(err)
	}
	managers.ProblemManager{}.SubmitByUser(&problem)
	c.JSON(http.StatusOK, this.Success())
}

func (this *Problem) httpHandlerDeleteProblem(c *gin.Context) {
	id := this.MustInt64("id", c)
	managers.ProblemManager{}.RemoveProblemUser(id)
	c.JSON(http.StatusOK, this.Success())
}

func (this *Problem) httpHandlerUpdateProblem(c *gin.Context) {
	var problem models.ProblemUser
	if err := c.BindJSON(&problem); err != nil {
		panic(err)
	}

	managers.ProblemManager{}.UpdateByUser(&problem)
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
	resp := baseController.PagingResponse{RequestPage: requestPage, TotalPages: num, List: problems}
	c.JSON(http.StatusOK, this.Success(resp))
}

func (this *Problem) httpHandlerAddProblemCheck(c *gin.Context) {
	problemId := this.MustInt64("id", c)
	managers.ProblemManager{}.AddProblemCheck(problemId)
	c.JSON(http.StatusOK, this.Success())
}
