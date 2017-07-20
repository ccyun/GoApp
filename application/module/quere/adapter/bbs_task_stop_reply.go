package adapter

import (
	"encoding/json"
	"fmt"
	"strconv"

	"bbs_server/application/function"
	"bbs_server/application/library/httpcurl"
	"bbs_server/application/model"
	"bbs_server/application/module/feed"
)

//TaskStopReply 广播任务提醒反馈
type TaskStopReply struct {
	base
}

func init() {
	Register("taskStopReply", func() Tasker {
		return new(TaskStopReply)
	})
}

//NewTask 新任务对象
func (T *TaskStopReply) NewTask(task model.Queue) error {
	T.base.NewTask(task)
	bbsID, err := strconv.Atoi(T.action)
	if err != nil {
		return fmt.Errorf("NewTask strconv.Atoi error,taskID:%d,action:%s", T.taskID, T.action)
	}
	T.bbsID = uint64(bbsID)
	if err := T.getBbsInfo(); err != nil {
		return err
	}
	if err := T.getBoardInfo(); err != nil {
		return err
	}
	if err := T.getBbsTaskInfo(); err != nil {
		return err
	}
	T.feedType = feed.FeedTypeTaskClose
	return nil
}

//GetPublishScopeUsers 分析发布范围
func (T *TaskStopReply) GetPublishScopeUsers() error {
	T.userIDs = new(model.Msg).GetUserIDs(T.siteID, T.boardID, T.bbsID, -1)
	T.userIDs = append(T.userIDs, T.boardInfo.EditorIDs...)
	userIDs, err := new(model.BbsTaskAudit).GetAuditUserIDs(T.bbsID, 0)
	if err != nil {
		return err
	}
	T.userIDs = append(T.userIDs, userIDs...)
	T.userIDs = function.SliceUnique(T.userIDs).Uint64()
	return T.base.GetPublishScopeUsers()
}

//SendMsg 发送消息
func (T *TaskStopReply) SendMsg() error {
	if len(T.PublishScopeuserLoginNames) == 0 {
		return nil
	}
	type Signal struct {
		httpcurl.SignalMsg
		DiscussID uint64 `json:"discuss_id"`
		BbsID     uint64 `json:"bbs_id"`
	}
	signalData := Signal{}
	signalData.Action = "task_update"
	signalData.BoardID = T.boardID
	signalData.DiscussID = T.boardInfo.DiscussID
	signalData.BbsID = T.bbsID
	uc := new(httpcurl.UC)
	data := httpcurl.CustomizedSender{
		SiteID:     strconv.FormatUint(T.siteID, 10),
		ToUsers:    T.PublishScopeuserLoginNames,
		ToPartyIds: T.bbsInfo.PublishScope.GroupIDs,
	}

	data1, err := json.Marshal(signalData)
	if err != nil {
		return err
	}
	data.Data1 = string(data1)
	return uc.CustomizedSend(data)
}
