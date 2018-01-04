package models

import (
	. "self/commons/store"
)

type Account struct {
	Id       int64  `form:"id" json:"id"`
	Email    string `form:"email" json:"email"`       //邮箱
	Password string `form:"password" json:"password"` //密码
	Phone    string `form:"phone" json:"phone"`       //手机号
	QqNumber string `form:"qqNumber" json:"qqNumber"` //QQ号
	QqId     int    `form:"qqId" json:"qqId"`         //用于第三方登录
}

//增加
func (this Account) Add(account *Account) (int64, error) {
	_, err := OrmWeb.Insert(account)
	if err != nil {
		return 0, err
	}
	return account.Id, nil
}

//删除
func (this Account) Remove(id int64) error {
	account := &Account{}
	_, err := OrmWeb.Id(id).Delete(account)
	return err
}

//修改
func (this Account) Update(account *Account) error {
	_, err := OrmWeb.AllCols().ID(account.Id).Update(account)
	return err
}

//查询
func (this Account) GetById(id int64) (*Account, error) {
	account := new(Account)

	has, err := OrmWeb.Id(id).Get(account)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return account, nil
}

func (this Account) GetByAccount(account *Account) (*Account, error) {
	has, err := OrmWeb.Get(account)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return account, nil
}
