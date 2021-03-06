package adapter

import (
	"bbs_server/application/model"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
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
	UpdateStatus() error
	SendMsg() error
}

//newTaskFunc
type newTasker func() Tasker

//queue 队列
type queue struct {
	task      model.Queue
	mode      Tasker
	model     *model.Queue
	RequestID string
}

//适配器
var modes = make(map[string]newTasker)

//Register 初测适配器
func Register(name string, mode newTasker) {
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
		q.RequestID = option["requestID"]
	}
	q.run()
}

func (q *queue) run() {
	logs.Info(q.L("Start processing tasks"))
	q.model = new(model.Queue)
	if q.getTask() == false {
		return
	}
	logs.Info(q.L("Start checkTask"))
	if q.checkTask() == false {
		return
	}
	logs.Info(q.L("Start runTask"))
	if q.runTask() == false { //任务失败
		logs.Critical(q.L("runTask error"))
		q.model.Fail(q.task.ID)
		return
	}
	logs.Info(q.L("Successful"))
}

//getTask 读取任务
func (q *queue) getTask() bool {
	var err error
	if err = q.model.TimeOut(); err != nil {
		logs.Critical(q.L("getTask TimeOut error:"), err)
	}
	if q.task, err = q.model.Pull(); err == nil {
		logs.Info(q.L("task info :"), q.task)
		return true
	}
	if err == orm.ErrNoRows {
		logs.Notice(q.L("Not found task info."))
	} else {
		logs.Critical(q.L("getTaskInfo Pull"), err)
	}
	return false
}

//checkTask 检查任务有效性
func (q *queue) checkTask() bool {
	mode, ok := modes[q.task.TaskType]
	if !ok {
		logs.Error(q.L("taskInfo.TaskType not in('bbs','delete','taskReply','taskAudit','taskClose','taskAuditRemind')."))
		if q.model.Fail(q.task.ID) == false {
			logs.Error(q.L("checkTask Fail error"))
		}
		return false
	}
	q.mode = mode()
	return true
}

//runTask 执行任务
func (q *queue) runTask() bool {
	//新的任务
	logs.Info(q.L("Start NewTask"))
	if err := q.mode.NewTask(q.task); err != nil {
		logs.Error(q.L("runTask NewTask error"), err)
		return false
	}
	//分析发布范围
	logs.Info(q.L("Start GetPublishScopeUsers"))
	if err := q.mode.GetPublishScopeUsers(); err != nil {
		logs.Error(q.L("runTask GetPublishScopeUsers error"), err)
		return false
	}
	//开启事务
	logs.Info(q.L("Start Begin"))
	if err := q.mode.Begin(); err != nil {
		logs.Error(q.L("runTask Begin error"), err)
		return false
	}
	//创建feed
	logs.Info(q.L("Start CreateFeed"))
	if err := q.mode.CreateFeed(); err != nil {
		logs.Error(q.L("runTask CreateFeed error"), err)
		q.mode.Rollback()
		return false
	}
	//写入未读计数
	logs.Info(q.L("Start CreateUnread"))
	if err := q.mode.CreateUnread(); err != nil {
		logs.Error(q.L("runTask CreateUnread error"), err)
		q.mode.Rollback()
		return false
	}
	//修改广播状态
	logs.Info(q.L("Start UpdateStatus"))
	if err := q.mode.UpdateStatus(); err != nil {
		logs.Error(q.L("runTask UpdateStatus error"), err)
		q.mode.Rollback()
		return false
	}
	//提交事务
	logs.Info(q.L("Start Commit"))
	if err := q.mode.Commit(); err != nil {
		logs.Error(q.L("runTask Commit error"), err)
		return false
	}
	//创建关系
	logs.Info(q.L("Start CreateRelation"))
	if err := q.mode.CreateRelation(); err != nil {
		logs.Error(q.L("runTask CreateRelation error"), err)
		return false
	}
	//关闭任务
	logs.Info(q.L("Start Delete task"))
	if q.model.Delete(q.task.ID) == false {
		logs.Error(q.L("runTask Delete error"))
		return false
	}

	model.AfterUpdate("bbs", q.task.SiteID)
	model.AfterUpdate("feed", q.task.SiteID)
	//发送消息
	logs.Info(q.L("Start SendMsg"))
	if err := q.mode.SendMsg(); err != nil {
		logs.Error(q.L("runTask SendMsg error"), err)
		return false
	}
	return true
}

//L 语言log
func (q *queue) L(l string) string {
	return q.RequestID + "  " + l
}
