package models

import (
	"testing"
)

func TestSaveCode_Create(t *testing.T) {

	InitAllInTest()

	savecode := &SaveCode{ProblemId: 1, Code: "http://xupt1.fightcoder.com:9001/code/1513783757.txt"}
	if _, err := savecode.Create(savecode); err != nil {
		t.Error("Create() failed. Error:", err)
	}
}

func TestSaveCode_Remove(t *testing.T) {

	InitAllInTest()

	var savecode SaveCode
	if err := savecode.Remove(1); err != nil {
		t.Error("Remove() failed. Error:", err)
	}
}

func TestSaveCode_Update(t *testing.T) {

	InitAllInTest()

	savecode := &SaveCode{ProblemId: 1, Code: "123456"}
	if err := savecode.Update(savecode); err != nil {
		t.Error("Update() failed. Error:", err)
	}
}

func TestLetter_GetById(t *testing.T) {

	InitAllInTest()

	savecode := &SaveCode{ProblemId: 1, Code: "123456"}
	SaveCode{}.Create(savecode)

	getSaveCode, err := SaveCode{}.GetById(savecode.Id)
	if err != nil {
		t.Error("GetById() failed:", err.Error())
	}

	if *getSaveCode != *savecode {
		t.Error("GetById() failed:", "%v != %v", savecode, getSaveCode)
	}
}

func TestLetter_GetBySaveCode(t *testing.T) {

	InitAllInTest()

	savecode := &SaveCode{ProblemId: 1, UserId: 2, Code: "123456"}
	SaveCode{}.Create(savecode)

	getSaveCode, err := SaveCode{}.GetBySaveCode(savecode.ProblemId, savecode.UserId)
	if err != nil {
		t.Error("GetBySaveCode() failed:", err.Error())
	}

	if *getSaveCode != *savecode {
		t.Error("GetBySaveCode() failed:", "%v != %v", savecode, getSaveCode)
	}
}
