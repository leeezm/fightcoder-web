/**
 * Created by shiyi on 2017/10/1.
 * Email: shiyi@fightcoder.com
 */

package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"self/commons"
	components "self/commons/components"
	"self/commons/g"
	"self/controllers"

	"github.com/TV4/graceful"
	"github.com/gin-gonic/gin"
)

func main() {
	//接收命令行参数
	version := flag.Bool("v", false, "show version")
	cfgfile := flag.String("c", "cfg/cfg.toml.debug", "set config file")
	flag.Parse()

	if *version {
		fmt.Println("version:", g.GitVer)
		fmt.Println("build time:", g.BuildTime)
		os.Exit(0)
	}

	// 初始化框架组件
	commons.InitAll(*cfgfile)

	gin.SetMode(g.Conf().Run.Mode)

	//初始化路由
	router := gin.Default()

	router.Use(Cors())

	router.Use(components.Check())

	controllers.Register(router)

	//优雅退出
	graceful.LogListenAndServe(&http.Server{
		Addr:    ":8888",
		Handler: router,
	})

	//关闭框架组件
	commons.CloseAll()
}

//解决跨域问题(待过滤器)
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}