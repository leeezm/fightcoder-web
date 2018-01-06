package managers

import (
	"self/commons/g"
	models "self/models"
)

type ProblemManager struct {
}

//代码保存
func (this ProblemManager) SaveCode(problemId, userId int64, code string) {
	path := NewMinioCli().SaveCode(code)
	saveCode, err := models.SaveCode{}.GetBySaveCode(problemId, userId)
	if err != nil {
		panic(err)
	}

	if saveCode == nil {
		saveCode := &models.SaveCode{ProblemId: problemId, UserId: userId, Code: path}
		if _, err := (models.SaveCode{}).Create(saveCode); err != nil {
			panic(err)
		}
	} else {
		saveCode.Code = path
		if err := (models.SaveCode{}).Update(saveCode); err != nil {
			panic(err)
		}
	}

}

func (this ProblemManager) GetCode(problemId, userId int64) string {
	saveCode, err := models.SaveCode{}.GetBySaveCode(problemId, userId)
	if err != nil || saveCode == nil {
		return ""
	} else {
		str := NewMinioCli().GetCode(saveCode.Code)
		return str
	}
}

//题库题目的检索
func (this ProblemManager) GetProblemById(id int64) *models.Problem {
	problem, err := models.Problem{}.GetById(id)
	if err != nil {
		panic(err)
	}
	return problem
}

func (this ProblemManager) GetsProblem(search, tag string, requestPage int) ([]*models.Problem, int) {
	cfg := g.Conf()
	problem := &models.Problem{Tag: tag}
	problems, err := models.Problem{}.QueryBySearch(search, problem, cfg.Show.PageNum, (requestPage-1)*cfg.Show.PageNum)
	if err != nil {
		panic(err)
	}
	count, err := models.Problem{}.CountBySearch(search, problem)
	if err != nil {
		panic(err)
	}
	return problems, int(count)
}

//用户上传题目
func (this ProblemManager) SubmitByUser(user *models.ProblemUser) {
	if user.Code != "" {
		path := NewMinioCli().SaveCode(user.Code)
		user.Code = path
	}
	if _, err := (models.ProblemUser{}).Create(user); err != nil {
		panic(err)
	}
}

//用户申请将题目加入题库
func (this ProblemManager) AddProblemCheck(problemUserId int64) {
	if pro, err := (models.ProblemUser{}).GetById(problemUserId); err != nil {
		panic(err)
	} else {
		if pro == nil {
			return
		}
		problemCheck := &models.ProblemCheck{}
		problemCheck.UserId = pro.UserId
		//problemCheck.TestData = pro.TestData
		//problemCheck.Title = pro.Title
		problemCheck.Description = pro.Description
		problemCheck.InputDes = pro.InputDes
		problemCheck.OutputDes = pro.OutputDes
		//problemCheck.Case = pro.Case
		problemCheck.Hint = pro.Hint
		problemCheck.TimeLimit = pro.TimeLimit
		problemCheck.MemoryLimit = pro.MemoryLimit
		problemCheck.Tag = pro.Tag
		problemCheck.IsSpecialJudge = pro.IsSpecialJudge
		problemCheck.SpecialJudgeSource = pro.SpecialJudgeSource
		problemCheck.Code = pro.Code
		problemCheck.LanguageLimit = pro.LanguageLimit
		problemCheck.CheckStatus = "审核中"
		//problemCheck.ProblemUserId = pro.Id

		if _, err := (models.ProblemCheck{}).Create(problemCheck); err != nil {
			panic(err)
		}
	}
}

//
////管理员查看审核题目
//func (this ProblemManager) GetsProblemCheck(status string, size, start int) []*models.ProblemCheck {
//	problems, err := models.ProblemCheck{}.QueryByStatus(status, size, start)
//	if err != nil {
//		panic(err)
//	}
//	return problems
//}
//
////管理员审核题目
//func (this ProblemManager) IsAddProblem(isAdd bool, checkId int64) {
//	pro, err := models.ProblemCheck{}.GetById(checkId)
//	if err != nil {
//		panic(err)
//	}
//	if isAdd {
//		problem := &models.Problem{}
//		problem.UserId = pro.UserId
//		problem.TestData = pro.TestData
//		problem.Title = pro.Title
//		problem.Description = pro.Description
//		problem.InputDes = pro.InputDes
//		problem.OutputDes = pro.OutputDes
//		problem.Case = pro.Case
//		problem.Case = pro.Case
//		problem.Hint = pro.Hint
//		problem.TimeLimit = pro.TimeLimit
//		problem.MemoryLimit = pro.MemoryLimit
//		problem.Tag = pro.Tag
//		problem.IsSpecialJudge = pro.IsSpecialJudge
//		problem.SpecialJudgeSource = pro.SpecialJudgeSource
//		problem.Code = pro.Code
//		problem.LanguageLimit = pro.LanguageLimit
//
//		if _, err := (models.Problem{}).Create(problem); err != nil {
//			panic(err)
//		}
//		pro.ProblemId = problem.Id
//		pro.CheckStatus = "审核通过"
//	} else {
//		pro.CheckStatus = "审核未通过"
//	}
//	if err := (models.ProblemCheck{}).Update(pro); err != nil {
//		panic(err)
//	}
//}

//用户查看上传题目
func (this ProblemManager) GetsProblemUser(userId int64, size, start int) []*models.ProblemUser {
	problems, err := models.ProblemUser{}.QueryByUserId(userId, size, start)
	if err != nil {
		panic(err)
	}

	return problems
}

func (this ProblemManager) CountProblemUser(userId int64) int {
	problemUser := &models.ProblemUser{UserId: userId}
	count, err := models.ProblemUser{}.Count(problemUser)
	if err != nil {
		panic(err)
	}
	return int(count)
}

//用户更新上传题目
func (this ProblemManager) UpdateByUser(user *models.ProblemUser) {
	if user.Code != "" {
		path := NewMinioCli().SaveCode(user.Code)
		user.Code = path
	}
	if err := (models.ProblemUser{}).Update(user); err != nil {
		panic(err)
	}
}

//用户删除上传题目
func (this ProblemManager) RemoveProblemUser(id int64) {
	if err := (models.ProblemUser{}).Remove(id); err != nil {
		panic(err)
	}

	submitUser := &models.SubmitUser{ProblemId: id}
	submits, err := models.SubmitUser{}.QueryBySubmitUserWithoutPaging(submitUser.ProblemId, submitUser.UserId, submitUser.Language, submitUser.ResultDes)

	if err != nil {
		panic(err)
	}
	if len(submits) == 0 {
		return
	}
	for _, sub := range submits {
		if sub == nil || sub.Id < 1 {
			continue
		}
		//删除对象存储中对应文件
		cli := NewMinioCli()
		name := cli.GetNameByPath(sub.Code)
		cli.RemoveCode(name)

		models.SubmitUser{}.Remove(sub.Id)
	}
}
