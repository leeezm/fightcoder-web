package managers

import (
	models "self/models"
)

type SubmitManager struct {
}

func (this SubmitManager) AddSubmit(problemId, userId int64, language string, submitTime int64, code string) {
	path := NewMinioCli().SaveCode(code)
	submit := &models.Submit{ProblemId: problemId, UserId: userId, Language: language, SubmitTime: submitTime, Code: path}
	if Id, err := (models.Submit{}).Create(submit); err != nil {
		panic(err)
	} else {
		Nsq{}.send("trueJudge", &SendMess{"problem", 1, "submit", Id})
	}
}

func (this SubmitManager) GetSubmitById(id int64) *models.Submit {
	if submit, err := (models.Submit{}).GetById(id); err != nil {
		panic(err)
	} else {
		return submit
	}
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