package models

import (
	"testing"
)

func TestSubmitContestCreate(t *testing.T) {
	InitAllInTest()

	submitContest := &SubmitContest{ProblemId: 1, UserId: 2, Language: "Java", SubmitTime: 1000, RunningTime: 10, RunningMemory: 100, Result: 1, ResultDes: "haha", Code: "sskajka"}
	if _, err := submitContest.Create(submitContest); err != nil {
		t.Error("Create() failed. Error:", err)
	}
}
func TestSubmitContestUpdate(t *testing.T) {
	InitAllInTest()

	submitContest := &SubmitContest{Id: 1, ResultDes: "haha"}
	if err := submitContest.Update(submitContest); err != nil {
		t.Error("Update() failed. Error:", err)
	}
}
func TestSubmitContestRemove(t *testing.T) {
	InitAllInTest()

	var submitContest SubmitContest
	if err := submitContest.Remove(1); err != nil {
		t.Error("Remove() failed. Error:", err)
	}
}
func TestSubmitContestGetById(t *testing.T) {
	InitAllInTest()

	submitContest := &SubmitContest{ProblemId: 1, UserId: 2, Language: "Java", SubmitTime: 1000, RunningTime: 10, RunningMemory: 100, Result: 1, ResultDes: "haha",Code: "sskajka"}
	submitContest.Create(submitContest)

	getSubmitContest, err := submitContest.GetById(submitContest.Id)
	if err != nil {
		t.Error("GetById() failed:", err.Error())
	}

	if *getSubmitContest != *submitContest {
		t.Error("GetById() failed:", "%v != %v", getSubmitContest, submitContest)
	}
}
func TestSubmitContestQueryBySubmitContest(t *testing.T) {
	InitAllInTest()

	var submitContest SubmitContest
	if _, err := submitContest.QueryBySubmitContest(2, 2, 1, "2", "1", 3, 0); err != nil {
		t.Error("QueryBySubmitContest() failed:", err)
	}
}
