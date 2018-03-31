/**
 * Created by leeezm on 2017/12/29.
 * Email: shiyi@fightcoder.com
 */

package managers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	jwt "self/commons/components"
	"self/models"
	"strconv"
	"strings"
	"time"
)

const (
	EMAIL_NOT_EXIT    = 0
	PASSWORD_IS_WRONG = 1
	PARAM_IS_WRONG    = 2
	FIRST_LOGIN       = 3
	LOGIN             = 4
)

type AccountManager struct {
}

type Mess struct {
	Ret          int    `json:"ret"`
	Msg          string `json:"msg"`
	Nickname     string `json:"nickname"`
	Gender       string `json:"gender"`
	Province     string `json:"province"`
	City         string `json:"city"`
	Year         string `json:"year"`
	Figureurl    string `json:"figureurl"`
	Figureurl1   string `json:"figureurl_1"`
	Figureurl2   string `json:"figureurl_2"`
	FigureurlQq1 string `json:"figureurl_qq_1"`
	FigureurlQq2 string `json:"figureurl_qq_2"`
}

func (this AccountManager) Register(email, password string) int64 {
	account := &models.Account{Email: email, Password: md5Encode(password)}
	accountId, err := models.Account{}.Add(account)
	if err != nil {
		panic(err)
	}
	user := &models.User{AccountId: accountId}
	userId, erro := models.User{}.Create(user)
	if erro != nil {
		panic(err)
	}
	return userId
}

func (this AccountManager) getGithubOpenId(code string) string {
	if code == "" {
		return "-1"
	} else {
		params := "client_id=080191e49e855122ea33&client_secret=34b9a36397b171f01e83fc3c5b676177b29df79e&code="
		params += code
		resp, err := http.Post("https://github.com/login/oauth/access_token",
			"application/x-www-form-urlencoded",
			strings.NewReader(params))
		if err != nil {
			fmt.Println(err)
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		strs := strings.Split(string(body), "&")
		token := strings.Split(strs[0], "=")

		url := "https://api.github.com/user?access_token="
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
		return strs[3]
	}
}

func (this AccountManager) getQQOpenId(code string) string {
	if code == "" {
		return "-1"
	} else {
		url := "https://graph.qq.com/oauth2.0/token?grant_type=authorization_code&client_id=101466300&client_secret=0104260a8f8faac3900cbf184bae55f5&redirect_uri=http%3a%2f%2fxupt4.fightcoder.com%2f%23%2fuser%2flogin&code="
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

		return strs[7]

		//url = "https://graph.qq.com/user/get_user_info?oauth_consumer_key=101466300&access_token=" + token[1] + "&openid=" + strs[7]
		//resp, err = http.Get(url)
		//if err != nil {
		//	panic(err.Error())
		//}
		//
		//defer resp.Body.Close()
		//body, err = ioutil.ReadAll(resp.Body)
		//if err != nil {
		//	panic(err.Error())
		//}
		//mess := &Mess{}
		//if err = json.Unmarshal(body, mess); err != nil {
		//	fmt.Println(err.Error())
		//}
		//c.JSON(http.StatusOK, this.Success(mess))
	}
}

func (this AccountManager) Login(param1, param2, loginType string) (int, string, int64) {
	var accountId int64
	isFirstLogin := false

	if loginType == "simple" {
		acc := &models.Account{Email: param1}
		account, err := models.Account{}.GetByAccount(acc)
		if err != nil {
			panic(err)
		}

		if account == nil {
			return EMAIL_NOT_EXIT, "", 0
		} else {
			passwd := account.Password
			if passwd != md5Encode(param2) {
				return PASSWORD_IS_WRONG, "", 0
			}
		}

		accountId = account.Id
	} else if loginType == "qq" {
		openId := this.getQQOpenId(param1)
		account := &models.Account{QqId: openId}
		acc, _ := account.GetByAccount(account)
		if acc == nil {
			id, _ := account.Add(account)

			user := &models.User{AccountId: id, NickName: strconv.FormatInt(time.Now().UnixNano(), 10)}
			user.Create(user)
			accountId = id
			isFirstLogin = true
		} else {
			accountId = acc.Id
		}

	} else {
		return PARAM_IS_WRONG, "", 0
	}

	user, err := models.User{}.GetByAccountId(accountId)
	if err != nil {
		fmt.Println(err.Error())
	}
	token := jwt.GetToken(user)
	if isFirstLogin {
		return FIRST_LOGIN, token, user.Id
	} else {
		return LOGIN, token, user.Id
	}
}
