/**
 * Created by shiyi on 2017/10/1.
 * Email: shiyi@fightcoder.com
 */

package controllers

import (
	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	v1 := router.Group("/auth/v1")

	problem := &Problem{}
	problem.Register(v1)

	submit := &Submit{}
	submit.Register(v1)

	user := &User{}
	user.Register(v1)
}
