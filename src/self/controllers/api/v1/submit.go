package controllers

import (
	"net/http"
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
	problemId := c.Query("problemId")
	language := c.Query("language")
	resultDes := c.Query("resultDes")
	requestPage := this.MustInt("requestPage", c)

	submits, err := managers.SubmitManager{}.GetsSubmit(problemId, language, resultDes, cfg.Show.PageNum, (requestPage-1)*cfg.Show.PageNum)
	if err {
		c.JSON(http.StatusOK, this.Fail())
	} else {
		num := managers.SubmitManager{}.CountSubmit(problemId, language, resultDes)
		if num%cfg.Show.PageNum != 0 {
			num = num/(cfg.Show.PageNum) + 1
		} else {
			num = num / (cfg.Show.PageNum)
		}
		resp := baseController.PagingResponse{RequestPage: requestPage, TotalPages: num, Data: submits}
		c.JSON(http.StatusOK, this.Success(resp))
	}
}
