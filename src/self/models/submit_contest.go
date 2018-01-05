package models

import (
	. "self/commons/store"
)

type SubmitContest struct {
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
	ContestId     int64  //比赛Id
}

//增加
func (this SubmitContest) Create(submitContest *SubmitContest) (int64, error) {
	_, err := OrmWeb.Insert(submitContest)
	if err != nil {
		return 0, err
	}
	return submitContest.Id, nil
}

//删除
func (this SubmitContest) Remove(id int64) error {
	submitContest := SubmitContest{}
	_, err := OrmWeb.Id(id).Delete(submitContest)
	return err
}

//修改
func (this SubmitContest) Update(submitContest *SubmitContest) error {
	_, err := OrmWeb.AllCols().ID(submitContest.Id).Update(submitContest)
	return err
}

//查询
func (this SubmitContest) GetById(id int64) (*SubmitContest, error) {
	submitContest := new(SubmitContest)
	has, err := OrmWeb.Id(id).Get(submitContest)

	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return submitContest, nil
}

func (this SubmitContest) QueryBySubmitContest(contestId, problemId, userId int64, language, resultDes string, size, start int) ([]*SubmitContest, error) {
	submitContest := SubmitContest{ProblemId: problemId, UserId: userId, Language: language, ResultDes: resultDes, ContestId: contestId}
	submitContestList := make([]*SubmitContest, 0)

	err := OrmWeb.Limit(size, start).Find(&submitContestList, submitContest)

	if err != nil {
		return nil, err
	}
	return submitContestList, nil
}

func (this SubmitContest) Count(contestId, problemId, userId int64, language, resultDes string) (int64, error) {
	submitContest := SubmitContest{ProblemId: problemId, UserId: userId, Language: language, ResultDes: resultDes, ContestId: contestId}
	sum, err := OrmWeb.Count(submitContest)
	if err != nil {
		return 0, err
	}
	return sum, nil
}
