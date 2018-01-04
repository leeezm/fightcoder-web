package managers

import (
	models "self/models"
	"strconv"
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

func (this SubmitManager) CountSubmit(problemId, language, resultDes string) int {
	id, _ := strconv.ParseInt(problemId, 10, 64)
	if sum, err := (models.Submit{}).Count(id, 0, language, resultDes); err != nil {
		panic(err)
	} else {
		return int(sum)
	}
}

func (this SubmitManager) GetsSubmit(problemId, language, resultDes string, size, start int) ([]*models.SubmitResponse, bool) {
	var response []*models.SubmitResponse
	var flag bool
	id, _ := strconv.ParseInt(problemId, 10, 64)
	submits, err := (models.Submit{}).QueryBySubmit(id, 0, language, resultDes, size, start)
	if submits == nil || err != nil {
		flag = true
	} else {
		length := len(submits)
		response = make([]*models.SubmitResponse, length)
		for i, v := range submits {
			res := &models.SubmitResponse{Submit: v}
			problem, err := models.Problem{}.GetById(res.ProblemId)
			if err != nil || problem == nil {
				flag = true
				break
			}
			res.Title = problem.Title
			response[i] = res
		}
	}
	return response, flag
}
