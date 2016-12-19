package model

import (
	"log"
	"testing"

	"github.com/ccyun/GoApp/application/function"
)

func Test_publishScopeHandle(t *testing.T) {
	a := new(Bbs)
	s := `{"discuss_ids":["50032726"],"group_ids":["54299","54342"],"user_ids":["62073932"]}`
	v, err := a.publishScopeHandle(s)
	if err != nil {
		t.Error("model->bbs.publishScopeHandle err", err)
	}
	if v.DiscussIDs[0] != 50032726 || v.GroupIDs[0] != 54299 || v.GroupIDs[1] != 54342 || v.UserIDs[0] != 62073932 {
		t.Error("model->bbs.publishScopeHandle err", s, v)
	}
}

func Test_cache(t *testing.T) {

	log.Println(function.Md5("aaaaa"))

	//err := json.Unmarshal([]byte(value.(string)), data)
}
