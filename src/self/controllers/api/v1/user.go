/**
 * Created by leeezm on 2017/12/30.
 * Email: shiyi@fightcoder.com
 */

package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"self/controllers/baseController"
	"self/managers"
	"strings"

	"github.com/gin-gonic/gin"
)

type User struct {
	baseController.Base
}

type Mess struct {
	Ret             int    `json:"ret"`
	Msg             string `json:"msg"`
	IsLost          int    `json:"is_lost"`
	Nickname        string `json:"nickname"`
	Gender          string `json:"gender"`
	Province        string `json:"province"`
	Year            string `json:"year"`
	Figureurl       string `json:"figureurl"`
	Figureurl1      string `json:"figureurl_1"`
	Figureurl2      string `json:"figureurl_2"`
	FigureurlQq1    string `json:"figureurl_qq_1"`
	FigureurlQq2    string `json:"figureurl_qq_2"`
	Vip             string `json:"vip"`
	YellowVipLevel  string `json:"yellow_vip_level"`
	Level           string `json:"level"`
	IsYellowYearVip string `json:"is_yellow_year_vip"`
}

func (this *User) Register(routergrp *gin.RouterGroup) {
	routergrp.GET("/qq_connect/callback", this.httpHandlerQqConnect)
	routergrp.POST("/login", this.httpHandlerLogin)
	routergrp.POST("/register", this.httpHandlerRegister)
	routergrp.POST("/check", this.httpHandlerCheck)
}

func (this *User) httpHandlerQqConnect(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusOK, this.Fail())
	} else {
		url := "https://graph.qq.com/oauth2.0/token?grant_type=authorization_code&client_id=101466300&client_secret=0104260a8f8faac3900cbf184bae55f5&redirect_uri=http%3a%2f%2fxupt4.fightcoder.com%2f%23%2fproblem&code="
		url += code
		resp, err := http.Get(url)
		if err != nil {
			panic(err.Error())
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err.Error())
		}
		strs := strings.Split(string(body), "&")
		token := strings.Split(strs[0], "=")

		url = "https://graph.qq.com/oauth2.0/me?access_token="
		resp, err = http.Get(url + token[1])
		if err != nil {
			panic(err.Error())
		}

		defer resp.Body.Close()
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err.Error())
		}

		strs = strings.Split(string(body), "\"")

		url = "https://graph.qq.com/user/get_user_info?oauth_consumer_key=101466300&access_token=" + token[1] + "&openid=" + strs[7]
		resp, err = http.Get(url)
		if err != nil {
			panic(err.Error())
		}

		defer resp.Body.Close()
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err.Error())
		}
		mess := &Mess{}
		if err = json.Unmarshal(body, mess); err != nil {
			fmt.Println(err.Error())
		}
		c.JSON(http.StatusOK, this.Success(mess))
	}
}

func (this *User) httpHandlerLogin(c *gin.Context) {
	//根据code获取openid,判断用户是否第一次第三方登录，如果是，调至完善信息==>返回Token
	//否，登录成功，返回token
	//token用于进行登录标识和用户身份的标识
	email := this.MustString("email", c)
	password := this.MustString("password", c)
	code, msg := managers.UserManager{}.Login(email, password)

	if code == 1 {
		c.JSON(http.StatusOK, this.Fail(msg))
	} else {
		cookie := &http.Cookie{
			Name:     "token",
			Value:    base64.StdEncoding.EncodeToString([]byte(msg)),
			Path:     "/",
			HttpOnly: true,
		}
		http.SetCookie(c.Writer, cookie)
		c.JSON(http.StatusOK, this.Success("Login successful!"))
	}
}

func (this *User) httpHandlerRegister(c *gin.Context) {
	email := this.MustString("email", c)
	password := this.MustString("password", c)
	flag := managers.UserManager{}.Register(email, password)
	if flag {
		c.JSON(http.StatusOK, this.Fail("注册失败"))
	} else {
		c.JSON(http.StatusOK, this.Success("注册成功"))
	}
}

func (this *User) httpHandlerCheck(c *gin.Context) {
	email := this.MustString("email", c)
	flag := managers.UserManager{}.Check(email)
	if flag == 1 {
		c.JSON(http.StatusOK, this.Fail("try again!"))
	} else if flag == 2 {
		c.JSON(http.StatusOK, this.Fail("false"))
	} else {
		c.JSON(http.StatusOK, this.Success("true"))
	}
}
