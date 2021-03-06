package adapter

import (
	"encoding/json"
	"fmt"

	"bbs_server/application/library/httpcurl"
	"bbs_server/application/model"
	"bbs_server/application/module/message"
)

//TaskAuditRemind 广播任务审核提醒
type TaskAuditRemind struct {
	base
}

func init() {
	Register("taskAuditRemind", func() Tasker {
		return new(TaskAuditRemind)
	})
}

//NewTask 新任务对象
func (T *TaskAuditRemind) NewTask(task model.Queue) error {
	T.base.NewTask(task)
	var action map[string]uint64
	if err := json.Unmarshal([]byte(T.action), &action); err != nil {
		return fmt.Errorf("NewTask action Unmarshal error,taskID:%d,action:%s", T.taskID, T.action)
	}
	T.bbsID = action["bbs_id"]
	if err := T.getBbsInfo(); err != nil {
		return err
	}
	if err := T.getBoardInfo(); err != nil {
		return err
	}
	return nil
}

//CreateRelation 创建feed关系
func (T *TaskAuditRemind) CreateRelation() error {
	return nil
}

//GetPublishScopeUsers 分析发布范围
func (T *TaskAuditRemind) GetPublishScopeUsers() error {
	userIDs, err := new(model.BbsTaskAudit).GetAuditUserIDs(T.bbsID, 0)
	if err != nil {
		return err
	}
	T.userIDs = userIDs
	return T.base.GetPublishScopeUsers()
}

//SendMsg 发送消息
func (T *TaskAuditRemind) SendMsg() error {
	type Signal struct {
		httpcurl.SignalMsg
		DiscussID uint64 `json:"discuss_id"`
		BbsID     uint64 `json:"bbs_id"`
	}
	signalData := Signal{}
	signalData.Action = "audit_remind"
	signalData.BoardID = T.boardID
	signalData.DiscussID = T.boardInfo.DiscussID
	signalData.BbsID = T.bbsID
	data1, err := json.Marshal(signalData)
	if err != nil {
		return err
	}
	msg := message.NewBroadcastMsg(T.bbsInfo.SiteID, message.SIGNAL_MSG)
	msg.PackHead()
	msg.CustomizedMsg(string(data1), "", "", "")
	return msg.Send(T.userIDs)
}
