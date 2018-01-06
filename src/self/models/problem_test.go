/**
 * Created by shiyi on 2017/11/22.
 * Email: shiyi@fightcoder.com
 */

package models

import (
	"fmt"
	"testing"
)

func TestProblemCreate(t *testing.T) {
	InitAllInTest()

	problem := &Problem{Title: "sadas", Description: "1111", TimeLimit: 1000, MemoryLimit: 128000}
	if _, err := problem.Create(problem); err != nil {
		t.Error("Create() failed. Error:", err)
	}
}
func TestProblemRemove(t *testing.T) {
	InitAllInTest()

	var problem Problem
	if err := problem.Remove(6); err != nil {
		t.Error("Remove() failed. Error:", err)
	}
}
func TestProblemUpdate(t *testing.T) {
	InitAllInTest()

	problem := &Problem{Title: "sadas", Description: "asdasdasd"}
	if err := problem.Update(problem); err != nil {
		t.Error("Update() failed. Error:", err)
	}
}
func TestProblemGetById(t *testing.T) {
	InitAllInTest()

	problem := &Problem{Title: "sadas", Description: "fffff"}
	Problem{}.Create(problem)

	getProblem, err := Problem{}.GetById(problem.Id)
	if err != nil {
		t.Error("GetById() failed:", err.Error())
	}

	if *getProblem != *problem {
		t.Error("GetById() failed:", "%v != %v", problem, getProblem)
	}
}
func TestProblemQueryByTitle(t *testing.T) {
	InitAllInTest()

	problem1 := &Problem{Title: "测试"}
	problem2 := &Problem{Title: "测试"}

	Problem{}.Create(problem1)
	Problem{}.Create(problem2)

	//problemList, _ := Problem{}.QueryBySearch("测试", &Problem{UserId: 0}, 2, 0)
	//
	//for _, v := range problemList {
	//	fmt.Println(v.Id)
	//}

	//if err != nil {
	//	t.Error("QueryByTitile() failed:", err)
	//}
	//if len(problemList) != 2 {
	//	t.Error("QueryByTitile() failed:", "count is wrong!")
	//}
}
func TestProblemQueryByUserId(t *testing.T) {
	InitAllInTest()

	problem := &Problem{UserId: 20}
	problem.Create(problem)

	problem1 := &Problem{UserId: 20}
	problem.Create(problem1)

	//getProblem, err := problem.QueryByProblem(&Problem{UserId: 20}, 2, 0)
	//if err != nil {
	//	t.Error("QueryByUserId() failed:", err)
	//}
	//if len(getProblem) != 2 {
	//	t.Error("QueryByUserId() failed:", "count is wrong!")
	//}
}

func TestProblem_Count(t *testing.T) {
	InitAllInTest()

	sum, _ := Problem{}.CountBySearch("测试", &Problem{UserId: 0})
	fmt.Println(sum)
}
