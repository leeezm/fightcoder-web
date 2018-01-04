/**
 * Created by leeezm on 2017/12/27.
 * Email: shiyi@fightcoder.com
 */

package components

import (
	"encoding/base64"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCurrentId(token string) int64 {
	userId, err := strconv.ParseInt(GetPayLoad(token).Id, 10, 64)
	if err != nil {
		panic(err)
	}
	return userId
}

func Check() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/auth") {
			checkLogin(c)
		}
		c.Next()
	}
}

func checkLogin(c *gin.Context) {
	token, _ := c.Cookie("token")
	data, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		panic(err)
	}

	token = string(data)

	if token != "" {
		nowTime := time.Now().UnixNano() / 1000000
		if endTime, err := strconv.ParseInt(GetPayLoad(token).EndTime, 10, 64); err != nil {
			panic(err)
		} else {
			if nowTime <= endTime {
				if CheckToken(token) {
					c.Set("userId", GetCurrentId(token))
					return
				}
			}
		}
	}
	c.JSON(http.StatusOK, "验证失败，请重新登录!")
}
