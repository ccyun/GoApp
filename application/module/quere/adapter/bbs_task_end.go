package adapter

import (
	"encoding/json"
	"fmt"

	"bbs_server/application/model"
)

//TaskEnd 广播任务提醒反馈
type TaskEnd struct {
	base
}

func init() {
	Register("taskEnd", func() Tasker {
		return new(TaskEnd)
	})
}

//NewTask 新任务对象
func (T *TaskEnd) NewTask(task model.Queue) error {
	T.base.NewTask(task)
	var action map[string]uint64
	if err := json.Unmarshal([]byte(T.action), &action); err != nil {
		return fmt.Errorf("NewTask action Unmarshal error,taskID:%d,action:%s", T.taskID, T.action)
	}
	T.bbsID = action["bbs_id"]
	return nil
}

//GetPublishScopeUsers 分析发布范围
func (T *TaskEnd) GetPublishScopeUsers() error {
	return nil
}

//CreateRelation 创建接收者关系
func (T *TaskEnd) CreateRelation() error {
	return nil
}

//UpdateStatus 更新状态
func (T *TaskEnd) UpdateStatus() error {
	return new(model.Msg).UpdateTaskStatus(T.bbsID)
}
