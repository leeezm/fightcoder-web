package models

import (
	. "self/commons/store"
)

type Submit struct {
	Id            int64  `form:"id" json:"id"`
	ProblemId     int64  `form:"problemId" json:"problemId"`         //题目ID
	UserId        int64  `form:"userId" json:"userId"`               //提交用户ID
	Language      string `form:"language" json:"language"`           //提交语言
	SubmitTime    int64  `form:"submitTime" json:"submitTime"`       //提交时间
	RunningTime   int    `form:"runningTime" json:"runningTime"`     //耗时(ms)
	RunningMemory int    `form:"runningMemory" json:"runningMemory"` //所占空间
	Result        int    `form:"result" json:"result"`               //运行结果
	ResultDes     string `form:"resultDes" json:"resultDes"`         //结果描述
	Code          string `form:"code" json:"code"`                   //提交代码
}

//增加
func (this Submit) Create(submit *Submit) (int64, error) {
	_, err := OrmWeb.Insert(submit)
	if err != nil {
		return 0, err
	}
	return submit.Id, nil
}

//删除
func (this Submit) Remove(id int64) error {
	submit := Submit{}
	_, err := OrmWeb.Id(id).Delete(submit)
	return err
}

//修改
func (this Submit) Update(submit *Submit) error {
	_, err := OrmWeb.AllCols().ID(submit.Id).Update(submit)
	return err
}

//查询
func (this Submit) GetById(id int64) (*Submit, error) {
	submit := new(Submit)
	has, err := OrmWeb.Id(id).Get(submit)

	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return submit, nil
}

func (this Submit) QueryBySubmit(problemId, userId int64, language, resultDes string, size, start int) ([]*Submit, error) {
	submit := Submit{ProblemId: problemId, UserId: userId, Language: language, ResultDes: resultDes}
	submitList := make([]*Submit, 0)

	err := OrmWeb.Limit(size, start).Find(&submitList, submit)

	if err != nil {
		return nil, err
	}
	return submitList, nil
}

func (this Submit) Count(problemId, userId int64, language, resultDes string) (int64, error) {
	submit := Submit{ProblemId: problemId, UserId: userId, Language: language, ResultDes: resultDes}
	sum, err := OrmWeb.Count(submit)
	if err != nil {
		return 0, err
	}
	return sum, nil
}
