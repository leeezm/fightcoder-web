package models

import (
	. "self/commons/store"
)

type SaveCode struct {
	Id        int64  `form:"id" json:"id"`
	ProblemId int64  `form:"problemId" json:"problemId"`
	UserId    int64  `form:"userId" json:"userId"` //提交用户Id
	Code      string `form:"code" json:"code"`
}

//增加
func (this SaveCode) Create(savecode *SaveCode) (int64, error) {
	_, err := OrmWeb.Insert(savecode) //第一个参数为影响的行数
	if err != nil {
		return 0, err
	}
	return savecode.Id, nil
}

//删除
func (this SaveCode) Remove(id int64) error {
	savecode := new(SaveCode)
	_, err := OrmWeb.Id(id).Delete(savecode)
	return err
}

//修改
func (this SaveCode) Update(savecode *SaveCode) error {
	_, err := OrmWeb.AllCols().ID(savecode.Id).Update(savecode)
	return err
}

//查询
func (this SaveCode) GetById(id int64) (*SaveCode, error) {
	savecode := new(SaveCode)
	has, err := OrmWeb.Id(id).Get(savecode) //第一个为 bool 类型，表示是否查找到记录

	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return savecode, nil
}

func (this SaveCode) GetBySaveCode(problemId, userId int64) (*SaveCode, error) {
	savecode := new(SaveCode)
	has, err := OrmWeb.Where("problem_id = ? AND user_id = ?", problemId, userId).Get(savecode)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return savecode, nil
}
