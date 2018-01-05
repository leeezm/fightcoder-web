package controllers

import (
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
	//routergrp.GET("/submission", this.httpHandlerGetsSubmit)
}

//func (this *Submit) httpHandlerGetsSubmit(c *gin.Context) {
//	cfg := g.Conf()
//	search := c.Query("search")
//	language := c.Query("language")
//	result, _ := strconv.Atoi(c.Query("result"))
//	requestPage := this.MustInt("requestPage", c)
//
//	submits, err := managers.SubmitManager{}.GetsSubmit(search, language, result, cfg.Show.PageNum, (requestPage-1)*cfg.Show.PageNum)
//	if err {
//		c.JSON(http.StatusOK, this.Fail())
//	} else {
//		num := managers.SubmitManager{}.CountSubmit(problemId, language, result)
//		resp := baseController.PagingResponse{RequestPage: requestPage, TotalPages: num, List: submits}
//		c.JSON(http.StatusOK, this.Success(resp))
//	}
//}
