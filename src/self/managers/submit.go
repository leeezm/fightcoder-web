package managers

import (
	models "self/models"
	"strconv"
)

type SubmitManager struct {
}

func (this SubmitManager) GetAcRate(problemId int64) (string, bool) {
	ac, err := models.Submit{}.Count(problemId, 0, "", 4)
	count, err1 := models.Submit{}.Count(problemId, 0, "", 0)
	if count == 0 || err != nil || err1 != nil {
		return "", true
	} else {
		ac1 := float64(ac)
		count1 := float64(count)
		res := strconv.FormatFloat(ac1/count1*100, 'f', 2, 32)
		return res, false
	}
}

func (this SubmitManager) AddSubmit(problemId, userId int64, language string, submitTime int64, code string) (*models.Submit, bool) {
	path := NewMinioCli().SaveCode(code)
	submit := &models.Submit{ProblemId: problemId, UserId: userId, ProblemType: "real", Language: language, SubmitTime: submitTime, Code: path}
	ProblemManager{}.SaveCode(problemId, userId, code)
	if Id, err := (models.Submit{}).Create(submit); err != nil {
		return nil, true
	} else {
		Nsq{}.send("realJudge", &SendMess{"submit", Id})
	}
	return submit, false
}

func (this SubmitManager) GetSubmitById(id, userId int64) (*models.CompleteSubmitResponse, bool) {
	var flag bool
	res := &models.CompleteSubmitResponse{}
	submit, err := (models.Submit{}).GetById(id)
	rate, err1 := this.GetAcRate(submit.ProblemId)
	if submit == nil || err != nil || err1 {
		flag = true
	} else {
		code := NewMinioCli().GetCode(submit.Code)
		res.SubmitResponse = submit
		res.Code = code
		res.AcRate = rate
		if submit.UserId != userId {
			res.IsHide = true
		}
	}
	return res, flag
}

func (this SubmitManager) CountSubmit(problemId, language string, result int) int {
	id, _ := strconv.ParseInt(problemId, 10, 64)
	if sum, err := (models.Submit{}).Count(id, 0, language, result); err != nil {
		panic(err)
	} else {
		return int(sum)
	}
}

func (this SubmitManager) GetsSubmit(problemId, language string, result int, size, start int) ([]*models.SubmitResponse, bool) {
	var response []*models.SubmitResponse
	var flag bool
	id, _ := strconv.ParseInt(problemId, 10, 64)
	submits, err := (models.Submit{}).QueryBySubmit(id, 0, language, result, size, start)
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
