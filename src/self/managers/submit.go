package managers

import (
	"flag"
	models "self/models"
)

type SubmitManager struct {
}

func (this SubmitManager) AddSubmit(problemId, userId int64, language string, submitTime int64, code string) bool {
	path := NewMinioCli().SaveCode(code)
	submit := &models.Submit{ProblemId: problemId, UserId: userId, Language: language, SubmitTime: submitTime, Code: path}
	ProblemManager{}.SaveCode(problemId, userId, code)
	if Id, err := (models.Submit{}).Create(submit); err != nil {
		return true
	} else {
		Nsq{}.send("trueJudge", &SendMess{"problem", 1, "submit", Id})
	}
	return false
}

func (this SubmitManager) GetSubmitById(id, userId int64) (*models.CompleteSubmitResponse, bool) {
	var flag bool
	res := &models.CompleteSubmitResponse{}
	submit, err := (models.Submit{}).GetById(id)
	if submit == nil || err != nil {
		flag = true
	} else {
		res.SubmitResponse = submit
		if submit.UserId != userId {
			res.IsHide = true
		}
	}
	return res, flag
}

func (this SubmitManager) CountSubmit(problemId, userId int64, language, resultDes string) int {
	if sum, err := (models.Submit{}).Count(problemId, userId, language, resultDes); err != nil {
		panic(err)
	} else {
		return int(sum)
	}
}

func (this SubmitManager) GetsSubmit(problemId, userId int64, language, resultDes string, size, start int) []*models.Submit {
	if submits, err := (models.Submit{}).QueryBySubmit(problemId, userId, language, resultDes, size, start); err != nil {
		panic(err)
	} else {
		return submits
	}
}
