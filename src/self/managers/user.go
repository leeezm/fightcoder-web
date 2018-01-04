/**
 * Created by leeezm on 2017/12/29.
 * Email: shiyi@fightcoder.com
 */

package managers

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"time"

	jwt "self/commons/components"
	"self/commons/g"
	"self/models"
)

type UserManager struct {
}

func (this UserManager) Register(email, password string) {
	account := &models.Account{Email: email, Password: md5Encode(password)}
	accountId, err := models.Account{}.Add(account)
	if err != nil {
		panic(err)
	}
	user := &models.User{AccountId: accountId}
	_, erro := models.User{}.Create(user)
	if erro != nil {
		panic(err)
	}
}

func (this UserManager) Login(email, password string) (int, string) {
	acc := &models.Account{Email: email}
	account, err := models.Account{}.GetByAccount(acc)
	if err != nil {
		panic(err)
	}
	if account == nil {
		return 1, "email is not exist!"
	} else {
		passwd := account.Password
		if passwd == md5Encode(password) {
			if user, err := (models.User{}).GetByAccountId(account.Id); err != nil {
				panic(err)
			} else {
				cfg := g.Conf()
				header := &jwt.HeaderData{cfg.Jwt.EncodeStyle, cfg.Jwt.Type}
				endTime := strconv.FormatInt(time.Now().UnixNano()/1000000+cfg.Jwt.MaxEffectiveTime, 10)
				pay := &jwt.PayLoadData{endTime, user.NickName, strconv.FormatInt(user.Id, 10)}

				str := jwt.GetToken(header, pay)
				return 0, str
			}
		} else {
			return 1, "password is wrong!"
		}
	}
}

func md5Encode(password string) string {
	w := md5.New()
	io.WriteString(w, password) //将password写入到w中
	md5str := string(fmt.Sprintf("%x", w.Sum(nil)))
	return md5str
}
