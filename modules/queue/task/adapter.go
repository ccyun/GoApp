package task

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/ccyun/GoApp/model"
)

//Tasker 任务接口
type Tasker interface {
	Begin() error
	Rollback() error
	Commit() error
	Run() error
	CreateRelation() error
	CreateUnread() error
	SendMsg() error
}

//queue 队列
type queue struct {
	taskInfo model.Queue
	adapter  Tasker
	model    *model.Queue
}

//适配器
var adapters = make(map[string]Tasker)

//Register 初测适配器
func Register(name string, adapter Tasker) {
	if adapter == nil {
		panic("task: Register adapter is nil")
	}
	if _, ok := adapters[name]; ok {
		panic("task: Register called twice for adapter " + name)
	}
	adapters[name] = adapter
}

//Run 运行
func Run() {
	logs.Info("Start processing tasks")
	q := new(queue)
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
	logs.Info("Successful")
}

//getTaskInfo 读取任务
func (q *queue) getTaskInfo() (model.Queue, bool) {
	taskInfo, err := q.model.GetOneTask()
	if err != nil {
		if err == orm.ErrNoRows {
			logs.Notice("Not found task info.")
		} else {
			logs.Error(err)
		}
		return model.Queue{}, false
	}
	return taskInfo, true
}

//checkTask 检查任务有效性
func (q *queue) checkTask() bool {
	adapter, ok := adapters[q.taskInfo.TaskType]
	if !ok {
		logs.Error("taskInfo.TaskType not in('bbs','taskReply','taskAudit','taskClose').")
		if err := q.model.Update(q.taskInfo.ID); err != nil {
			logs.Error(err)
		}
		return false
	}
	q.adapter = adapter
	return true
}

//runTask 执行任务
func (q *queue) runTask() bool {
	//开启事务
	if err := q.adapter.Begin(); err != nil {
		logs.Error(err)
		return false
	}

	//提交事务
	if err := q.adapter.Commit(); err != nil {
		logs.Error(err)
		return false
	}
	if err := q.model.Delete(q.taskInfo.ID); err != nil {
		logs.Error(err)
		return false
	}
	return true
}
