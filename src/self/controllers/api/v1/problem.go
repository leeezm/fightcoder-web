package controllers

import (
	"net/http"
	baseController "self/controllers/baseController"
	managers "self/managers"

	"github.com/gin-gonic/gin"
)

type Problem struct {
	baseController.Base
}

func (this *Problem) Register(routergrp *gin.RouterGroup) {
	routergrp.GET("/problem/id", this.httpHandlerGetProblemByID)
	routergrp.GET("/problems", this.httpHandlerGetProblems)

}

func (this *Problem) httpHandlerGetProblemByID(c *gin.Context) {
	problemId := this.MustInt64("id", c)
	problem := managers.ProblemManager{}.GetProblemById(problemId)

	c.JSON(http.StatusOK, this.Success(problem))
}

func (this *Problem) httpHandlerGetProblems(c *gin.Context) {
	search := c.Query("search")
	tag := c.Query("tag")
	requestPage := this.MustInt("requestPage", c)
	problems, num := managers.ProblemManager{}.GetsProblem(search, tag, requestPage)
	resp := baseController.PagingResponse{RequestPage: requestPage, TotalPages: num, List: problems}
	c.JSON(http.StatusOK, this.Success(resp))
}
