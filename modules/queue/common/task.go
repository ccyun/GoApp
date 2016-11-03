package common

import (
	"log"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/ccyun/daemon/example/app/model"
)

//Tasker 任务接口
type Tasker interface {
	Run() error
}

//适配器
var controllers = make(map[string]Tasker)

//Register 初测适配器
func Register(name string, controller Tasker) {
	log.Println(name)
	if controller == nil {
		panic("task: Register adapter is nil")
	}
	if _, ok := controllers[name]; ok {
		panic("task: Register called twice for adapter " + name)
	}
	controllers[name] = controller
}

//Run 运行
func Run() {
	logs.Info("Start processing tasks")
	queue := new(model.Queue)
	taskInfo, err := queue.GetOneTask()
	if err != nil {
		if err == orm.ErrNoRows {
			logs.Notice("Not found task info.")
		} else {
			logs.Info(err)
		}
		return
	}
	controller, ok := controllers[taskInfo.TaskType]
	log.Println(taskInfo)
	if !ok {
		logs.Error("taskInfo.TaskType not in('bbs','taskReply','taskAudit','taskClose').")
		if err := queue.Update(taskInfo.ID); err != nil {
			logs.Error(err)
		}
		return
	}
	if err = controller.Run(); err != nil {
		logs.Info("")
	}
	if err := queue.Delete(taskInfo.ID); err != nil {
		logs.Error(err)
	}
	log.Println("Successful")
}
