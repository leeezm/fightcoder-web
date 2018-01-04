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

func (this UserManager) UploadImage(reader io.Reader, userId int64, picType string) bool {
	var flag bool
	path := NewMinioCli().SaveImg(reader, userId, picType)
	user, error := models.User{}.GetById(userId)
	if user == nil || error != nil {
		flag = true
	} else {
		user.Avator = path
		if err := (models.User{}).Update(user); err != nil {
			flag = true
		}
	}
	return flag
}

func (this UserManager) CompleteMess(id int64, nickName, description string, sex int, birthday, dailyAddress, recvAddress, tShirtSize string, statSchool int, blog, git string) error {
	birth, _ := strconv.ParseInt(birthday, 10, 64)
	user := &models.User{Id: id, NickName: nickName, Description: description, Sex: sex, Birthday: birth, DailyAddress: dailyAddress, RecvAddress: recvAddress, TShirtSize: tShirtSize, StatSchool: statSchool, Blog: blog, Git: git}
	err := (models.User{}).Update(user)
	return err
}

func (this UserManager) Check(email string) int {
	account, err := models.Account{}.GetByAccount(&models.Account{Email: email})
	if err != nil {
		return 1
	} else if account == nil {
		return 2
	} else {
		return 3
	}
}

func (this UserManager) Register(email, password string) bool {
	var flag bool
	account := &models.Account{Email: email, Password: md5Encode(password)}
	accountId, err := models.Account{}.Add(account)
	if err != nil {
		flag = true
	}
	user := &models.User{AccountId: accountId}
	_, erro := models.User{}.Create(user)
	if erro != nil {
		flag = true
	}
	return flag
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
