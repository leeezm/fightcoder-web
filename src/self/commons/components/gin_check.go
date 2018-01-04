/**
 * Created by leeezm on 2017/12/27.
 * Email: shiyi@fightcoder.com
 */

package components

import (
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

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
	if token != "" {
		nowTime := time.Now().UnixNano() / 1000000
		if endTime, err := strconv.ParseInt(GetPayLoad(token).EndTime, 10, 64); err != nil {
			panic(err)
		} else {
			if nowTime <= endTime {
				if CheckToken(token) {
					return
				}
			}
		}
	}
	loginUrl := "/api/v1/problems?requestPage=1"
	c.Redirect(302, loginUrl)
}
