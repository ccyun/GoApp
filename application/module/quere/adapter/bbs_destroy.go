package adapter

import (
	"encoding/json"
	"fmt"
	"strconv"

	"bbs_server/application/function"
	"bbs_server/application/library/httpcurl"
	"bbs_server/application/model"
	"bbs_server/application/module/message"
)

//BbsDestroy 广播删除
type BbsDestroy struct {
	base
}

func init() {
	Register("bbsDestroy", func() Tasker {
		return new(BbsDestroy)
	})
}

//NewTask 新任务对象
func (B *BbsDestroy) NewTask(task model.Queue) error {
	B.base.NewTask(task)
	bbsID, err := strconv.Atoi(B.action)
	if err != nil {
		return fmt.Errorf("NewTask strconv.Atoi error,taskID:%d,action:%s", B.taskID, B.action)
	}
	B.bbsID = uint64(bbsID)
	if err := B.getBbsInfo(); err != nil {
		return err
	}
	if err := B.getBoardInfo(); err != nil {
		return err
	}
	return nil
}

//GetPublishScopeUsers 分析发布范围
func (B *BbsDestroy) GetPublishScopeUsers() error {
	B.userIDs = new(model.Msg).GetUserIDs(B.siteID, B.boardID, B.bbsID, -1)
	B.userIDs = append(B.userIDs, B.boardInfo.EditorIDs...)
	userIDs, err := new(model.BbsTaskAudit).GetAuditUserIDs(B.bbsID, -1)
	if err != nil {
		return err
	}
	B.userIDs = append(B.userIDs, userIDs...)
	B.userIDs = function.SliceUnique(B.userIDs).Uint64()
	return B.base.GetPublishScopeUsers()
}

//CreateRelation 不创建关系
func (B *BbsDestroy) CreateRelation() error {
	return nil
}

//SendMsg 发送消息
func (B *BbsDestroy) SendMsg() error {
	type Signal struct {
		httpcurl.SignalMsg
		DiscussID uint64 `json:"discuss_id"`
		BbsID     uint64 `json:"bbs_id"`
	}
	signalData := Signal{}
	signalData.Action = "destroy_msg"
	signalData.BoardID = B.boardID
	signalData.DiscussID = B.boardInfo.DiscussID
	signalData.BbsID = B.bbsID
	data1, err := json.Marshal(signalData)
	if err != nil {
		return err
	}
	msg := message.NewBroadcastMsg(B.bbsInfo.SiteID, message.SIGNAL_MSG)
	msg.PackHead()
	msg.CustomizedMsg(string(data1), "", "", "")
	return msg.Send(B.userIDs)
}
