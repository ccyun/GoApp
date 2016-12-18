package adapter

import (
	"strconv"

	"github.com/astaxie/beego/logs"
	"github.com/ccyun/GoApp/application/model"
)

//Bbs 图文广播
type Bbs struct {
	base
}

func init() {
	Register("bbs", new(Bbs))
}

//NewTask 新任务对象
func (B *Bbs) NewTask(taskInfo model.Queue) error {
	B.base.NewTask(taskInfo)
	bbsID, err := strconv.Atoi(B.action)
	if err != nil {
		logs.Error(L("NewTask strconv.Atoi err"), "taskid: ", B.taskID, "action error, action:", B.action, err)
		return err
	}
	B.bbsID = uint64(bbsID)
	if err := B.getBbsInfo(); err != nil {
		logs.Error(L("getBbsInfo error"), err)
		return err
	}
	if err := B.getBoardInfo(); err != nil {
		logs.Error(L("getBoardInfo error"), err)
		return err
	}
	return nil
}
