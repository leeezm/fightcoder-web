package models

import (
	"fmt"
	"testing"
)

func TestSubmitCreate(t *testing.T) {
	InitAllInTest()

	submit := &Submit{ProblemId: 1, UserId: 2, Language: "Java", SubmitTime: 1000, RunningTime: 10, RunningMemory: 100, Result: 1, ResultDes: "haha", Code: "sskajka"}
	if _, err := submit.Create(submit); err != nil {
		t.Error("Create() failed. Error:", err)
	}
}
func TestSubmitUpdate(t *testing.T) {
	InitAllInTest()

	submit := &Submit{Id: 1, ResultDes: "haha"}
	if err := submit.Update(submit); err != nil {
		t.Error("Update() failed. Error:", err)
	}
}
func TestSubmitRemove(t *testing.T) {
	InitAllInTest()

	var submit Submit
	if err := submit.Remove(1); err != nil {
		t.Error("Remove() failed. Error:", err)
	}
}
func TestSubmitGetById(t *testing.T) {
	InitAllInTest()

	submit := &Submit{ProblemId: 1, UserId: 2, Language: "Java", SubmitTime: 1000, RunningTime: 10, RunningMemory: 100, Result: 1, ResultDes: "haha", Code: "sskajka"}
	submit.Create(submit)

	getSubmit, _ := submit.GetById(submit.Id)
	fmt.Println(getSubmit.Submit, getSubmit.Title)
	//if err != nil {
	//	t.Error("GetById() failed:", err.Error())
	//}
	//
	//if *getSubmit != *submit {
	//	t.Error("GetById() failed:", "%v != %v", submit, getSubmit)
	//}
}
func TestSubmitQueryBySubmit(t *testing.T) {
	InitAllInTest()

	submits, _ := Submit{}.QueryBySubmit(1, 123, "c_cpp", 1, 3, 0)

	fmt.Println(len(submits))
}
