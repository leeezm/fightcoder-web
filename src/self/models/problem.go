package models

import (
	. "self/commons/store"
)

type Problem struct {
	Id                 int64  `form:"id" json:"id"`
	UserId             int64  `form:"userId" json:"userId"`                         //题目提供者
	TestData           string `form:"testData" json:"testData"`                     //测试数据
	Title              string `form:"title" json:"title"`                           //题目标题
	Description        string `form:"description" json:"description"`               //题目描述
	InputDes           string `form:"inputDes" json:"inputDes"`                     //输入描述
	OutputDes          string `form:"outputDes" json:"outputDes"`                   //输出描述
	Case               string `form:"Case" json:"Case"`                             //样例输入
	Hint               string `form:"hint" json:"hint"`                             //题目提示(可以为对样例输入输出的解释)
	TimeLimit          int    `form:"timeLimit" json:"timeLimit"`                   //时间限制
	MemoryLimit        int    `form:"memoryLimit" json:"memoryLimit"`               //内存限制
	Tag                string `form:"tag" json:"tag"`                               //题目标签
	IsSpecialJudge     bool   `form:"isSpecialJudge" json:"isSpecialJudge"`         //是否特判
	SpecialJudgeSource string `form:"specialJudgeSource" json:"specialJudgeSource"` //特判程序源代码
	Code               string `form:"code" json:"code"`                             //标准程序
	LanguageLimit      string `form:"languageLimit" json:"languageLimit"`           //语言限制
}

//增加
func (this Problem) Create(problem *Problem) (int64, error) {
	_, err := OrmWeb.Insert(problem) //第一个参数为影响的行数
	if err != nil {
		return 0, err
	}
	return problem.Id, nil
}

//删除
func (this Problem) Remove(id int64) error {
	problem := new(Problem)
	_, err := OrmWeb.Id(id).Delete(problem)
	return err
}

//修改
func (this Problem) Update(problem *Problem) error {
	_, err := OrmWeb.AllCols().ID(problem.Id).Update(problem)
	return err
}

//查询
func (this Problem) GetById(id int64) (*Problem, error) {
	problem := new(Problem)
	has, err := OrmWeb.Id(id).Get(problem) //第一个为 bool 类型，表示是否查找到记录

	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return problem, nil
}

//func (this Problem) QueryByProblem(problem *Problem, size, start int) ([]*Problem, error) {
//	problemList := make([]*Problem, 0)
//	err := OrmWeb.Limit(size, start).Find(&problemList, problem)
//	if err != nil {
//		return nil, err
//	}
//	return problemList, nil
//}
//
//func (this Problem) Count(problem *Problem) (int64, error) {
//	sum, err := OrmWeb.Count(problem)
//	if err != nil {
//		return 0, err
//	}
//	return sum, nil
//}

func (this Problem) QueryBySearch(title string, problem *Problem, size, start int) ([]*Problem, error) {
	problemList := make([]*Problem, 0)
	err := OrmWeb.Where("title like ? or id = ?", "%"+title+"%", title).Limit(size, start).Find(&problemList, problem)
	if err != nil {
		return nil, err
	}
	return problemList, nil
}

func (this Problem) CountBySearch(title string, problem *Problem) (int64, error) {
	sum, err := OrmWeb.Where("title like ? or id = ?", "%"+title+"%", title).Count(&Problem{}, problem)
	if err != nil {
		return 0, err
	}
	return sum, nil
}
