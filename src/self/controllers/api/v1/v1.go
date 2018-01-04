/**
 * Created by shiyi on 2017/10/1.
 * Email: shiyi@fightcoder.com
 */

package controllers

import (
	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	problem := &Problem{}
	problem.Register(v1)

	user := &User{}
	user.Register(v1)
}
