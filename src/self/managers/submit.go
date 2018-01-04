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

func (this SubmitManager) AddSubmitContest(problemId, userId int64, language string, submitTime int64, code string, contestId int64) {
	path := NewMinioCli().SaveCode(code)
	submitContest := &models.SubmitContest{ProblemId: problemId, UserId: userId, Language: language, SubmitTime: submitTime, Code: path, ContestId: contestId}
	if _, err := (models.SubmitContest{}).Create(submitContest); err != nil {
		panic(err)
	}
}

func (this SubmitManager) GetSubmitContestById(id int64) *models.SubmitContest {
	if submitContest, err := (models.SubmitContest{}).GetById(id); err != nil {
		panic(err)
	} else {
		return submitContest
	}
}

func (this SubmitManager) CountSubmitContest(contestId, problemId, userId int64, language, resultDes string) int {
	if sum, err := (models.SubmitContest{}).Count(contestId, problemId, userId, language, resultDes); err != nil {
		panic(err)
	} else {
		return int(sum)
	}
}

func (this SubmitManager) GetsSubmitContest(contestId, problemId, userId int64, language, resultDes string, size, start int) []*models.SubmitContest {
	if submitContests, err := (models.SubmitContest{}).QueryBySubmitContest(contestId, problemId, userId, language, resultDes, size, start); err != nil {
		panic(err)
	} else {
		return submitContests
	}
}

func (this SubmitManager) AddSubmitUser(problemId, userId int64, language string, submitTime int64, code string) {
	path := NewMinioCli().SaveCode(code)
	submitUser := &models.SubmitUser{ProblemId: problemId, UserId: userId, Language: language, SubmitTime: submitTime, Code: path}
	if _, err := (models.SubmitUser{}).Create(submitUser); err != nil {
		panic(err)
	}
}

func (this SubmitManager) GetSubmitUserById(id int64) *models.SubmitUser {
	if submitUser, err := (models.SubmitUser{}).GetById(id); err != nil {
		panic(err)
	} else {
		return submitUser
	}
}

func (this SubmitManager) CountSubmitUser(problemId, userId int64, language, resultDes string) int {
	if sum, err := (models.SubmitUser{}).Count(problemId, userId, language, resultDes); err != nil {
		panic(err)
	} else {
		return int(sum)
	}
}

func (this SubmitManager) GetsSubmitUser(problemId, userId int64, language, resultDes string, size, start int) []*models.SubmitUser {
	if submitUsers, err := (models.SubmitUser{}).QueryBySubmitUser(problemId, userId, language, resultDes, size, start); err != nil {
		panic(err)
	} else {
		return submitUsers
	}
}
