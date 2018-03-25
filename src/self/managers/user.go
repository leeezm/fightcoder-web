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

func md5Encode(password string) string {
	w := md5.New()
	io.WriteString(w, password) //将password写入到w中
	md5str := string(fmt.Sprintf("%x", w.Sum(nil)))
	return md5str
}
