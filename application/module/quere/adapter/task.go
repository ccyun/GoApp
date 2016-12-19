package adapter

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/ccyun/GoApp/application/model"
)

//Tasker 任务接口
type Tasker interface {
	Begin() error
	Rollback() error
	Commit() error
	NewTask(model.Queue) error
	CreateFeed() error
	CreateRelation() error
	CreateUnread() error
	GetPublishScopeUsers() error
	SendMsg() error
}

//queue 队列
type queue struct {
	taskInfo model.Queue
	mode     Tasker
	model    *model.Queue
}

//RequestID 请求ID
var RequestID string

//适配器
var modes = make(map[string]Tasker)

//Register 初测适配器
func Register(name string, mode Tasker) {
	if mode == nil {
		panic("task: Register mode is nil")
	}
	if _, ok := modes[name]; ok {
		panic("task: Register called twice for adapter " + name)
	}
	modes[name] = mode
}

//Run 运行
func Run(option map[string]string) {
	q := new(queue)
	if option["requestID"] != "" {
		RequestID = option["requestID"]
	}
	q.run()
}

func (q *queue) run() {

	logs.Info(L("Start processing tasks"))
	q.model = new(model.Queue)

	taskInfo, ok := q.getTaskInfo()
	if ok == false {
		return
	}
	q.taskInfo = taskInfo
	if q.checkTask() == false {
		return
	}
	if q.runTask() == false {
		return
	}
	logs.Info(L("Successful"))
}

//getTaskInfo 读取任务
func (q *queue) getTaskInfo() (model.Queue, bool) {
	new(model.Bbs).GetOne(13)
	//model.NewCache("bbs", "").ClearCache("bbs")

	taskInfo, err := q.model.GetOneTask()
	if err != nil {
		if err == orm.ErrNoRows {
			logs.Notice(L("Not found task info."))
		} else {
			logs.Error(L("getTaskInfo GetOneTask"), err)
		}
		return model.Queue{}, false
	}
	return taskInfo, true
}

//checkTask 检查任务有效性
func (q *queue) checkTask() bool {
	mode, ok := modes[q.taskInfo.TaskType]
	if !ok {
		logs.Error(L("taskInfo.TaskType not in('bbs','taskReply','taskAudit','taskClose')."))
		if err := q.model.Update(q.taskInfo.ID); err != nil {
			logs.Error(L("checkTask Update error"), err)
		}
		return false
	}
	q.mode = mode
	return true
}

//runTask 执行任务
func (q *queue) runTask() bool {
	//开启事务
	if err := q.mode.Begin(); err != nil {
		logs.Error(L("runTask Begin error"), err)
		q.mode.Rollback()
		return false
	}
	//新的任务
	if err := q.mode.NewTask(q.taskInfo); err != nil {
		logs.Error(L("runTask NewTask error"), err)
		q.mode.Rollback()
		return false
	}
	//分析发布范围

	//创建feed

	//创建关系

	//写入未读计数

	//修改广播状态

	//提交事务
	if err := q.mode.Commit(); err != nil {
		logs.Error(L("runTask Commit error"), err)
		return false
	}

	// if err := q.model.Delete(q.taskInfo.ID); err != nil {
	// 	logs.Error(q.requestID, err)
	// 	return false
	// }
	return true
}

//L 语言log
func L(l string) string {
	return RequestID + "  " + l
}
