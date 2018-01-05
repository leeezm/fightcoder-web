package models

import (
	. "self/commons/store"
)

type SubmitUser struct {
	Id            int64
	ProblemId     int64  //题目ID
	ProblemType   string //题库类型
	UserId        int64  //提交用户ID
	Language      string //提交语言
	SubmitTime    int64  //提交时间
	RunningTime   int    //耗时(ms)
	RunningMemory int    //所占空间
	Result        int    //运行结果
	ResultDes     string //结果描述
	Code          string //提交代码
}

//增加
func (this SubmitUser) Create(submitUser *SubmitUser) (int64, error) {
	_, err := OrmWeb.Insert(submitUser)
	if err != nil {
		return 0, err
	}
	return submitUser.Id, nil
}

//删除
func (this SubmitUser) Remove(id int64) error {
	submitUser := SubmitUser{}
	_, err := OrmWeb.Id(id).Delete(submitUser)
	return err
}

//修改
func (this SubmitUser) Update(submitUser *SubmitUser) error {
	_, err := OrmWeb.AllCols().ID(submitUser.Id).Update(submitUser)
	return err
}

//查询
func (this SubmitUser) GetById(id int64) (*SubmitUser, error) {
	submitUser := new(SubmitUser)
	has, err := OrmWeb.Id(id).Get(submitUser)

	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return submitUser, nil
}

func (this SubmitUser) QueryBySubmitUserWithoutPaging(problemId, userId int64, language, resultDes string) ([]*SubmitUser, error) {
	submitUser := SubmitUser{ProblemId: problemId, UserId: userId, Language: language, ResultDes: resultDes}
	submitUserList := make([]*SubmitUser, 0)

	err := OrmWeb.Find(&submitUserList, submitUser)

	if err != nil {
		return nil, err
	}
	return submitUserList, nil
}

func (this SubmitUser) QueryBySubmitUser(problemId, userId int64, language, resultDes string, size, start int) ([]*SubmitUser, error) {
	submitUser := SubmitUser{ProblemId: problemId, UserId: userId, Language: language, ResultDes: resultDes}
	submitUserList := make([]*SubmitUser, 0)

	err := OrmWeb.Limit(size, start).Find(&submitUserList, submitUser)

	if err != nil {
		return nil, err
	}
	return submitUserList, nil
}

func (this SubmitUser) Count(problemId, userId int64, language, resultDes string) (int64, error) {
	submitUser := SubmitUser{ProblemId: problemId, UserId: userId, Language: language, ResultDes: resultDes}
	sum, err := OrmWeb.Count(submitUser)
	if err != nil {
		return 0, err
	}
	return sum, nil
}
