package mode

import (
	"strconv"

	"github.com/astaxie/beego/logs"
	"github.com/ccyun/GoApp/model"
)

//Bbs 图文广播
type Bbs struct {
	base
}

//NewTask 新任务对象
func (b *Bbs) NewTask(taskInfo model.Queue) error {
	b.base.NewTask(taskInfo)
	bbsID, err := strconv.Atoi(b.action)
	if err != nil {
		logs.Error(b.requestID, "taskid: ", b.taskID, "action error, action:", b.action, err)
		return err
	}
	b.bbsID = uint64(bbsID)

	b.getBbsInfo()

	return nil
}

func init() {
	Register("bbs", new(Bbs))
}
