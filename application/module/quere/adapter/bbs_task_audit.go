package adapter

import (
	"encoding/json"
	"fmt"

	"bbs_server/application/model"
	"bbs_server/application/module/feed"
	"bbs_server/application/module/message"
)

//TaskAudit 广播任务审核
type TaskAudit struct {
	base
	status int8
}

func init() {
	Register("taskAudit", func() Tasker {
		return new(TaskAudit)
	})
}

//NewTask 新任务对象
func (T *TaskAudit) NewTask(task model.Queue) error {
	T.base.NewTask(task)
	var action map[string]uint64
	if err := json.Unmarshal([]byte(T.action), &action); err != nil {
		return fmt.Errorf("NewTask action Unmarshal error,taskID:%d,action:%s", T.taskID, T.action)
	}
	T.bbsID = action["bbs_id"]
	T.status = int8(action["status"])
	T.userIDs = []uint64{uint64(action["user_id"])}
	if err := T.getBbsInfo(); err != nil {
		return err
	}
	if err := T.getBoardInfo(); err != nil {
		return err
	}
	if err := T.getBbsTaskInfo(); err != nil {
		return err
	}
	T.feedType = feed.FeedTypeTaskAudit
	return nil
}

//CreateFeed 创建Feed
func (T *TaskAudit) CreateFeed() error {
	feedData := model.Feed{
		SiteID:    T.siteID,
		BoardID:   T.boardID,
		BbsID:     T.bbsID,
		FeedType:  feed.FeedTypeTaskAudit,
		CreatedAt: T.nowTime,
	}
	data := model.FeedData{
		Title:          T.bbsInfo.Title,
		Description:    T.bbsInfo.Description,
		CreatedAt:      T.nowTime,
		UserID:         T.bbsInfo.UserID,
		Type:           T.bbsInfo.Type,
		Category:       T.category,
		CommentEnabled: T.bbsInfo.CommentEnabled,
		EndTime:        T.bbsTaskInfo.EndTime,
		AllowExpired:   T.bbsTaskInfo.AllowExpired,
		Status:         T.status,
	}
	dataByte, err := json.Marshal(data)
	if err != nil {
		return err
	}
	feedData.Data = string(dataByte)
	feedID, err := T.o.Insert(&feedData)
	if err == nil {
		T.feedID = uint64(feedID)
	}
	return err
}

//SendMsg 发送消息
func (T *TaskAudit) SendMsg() error {
	feedData, err := feed.NewTask(feed.FeedTypeTaskAudit, feed.Customizer{
		BoardID:        T.boardID,
		BoardName:      T.boardInfo.BoardName,
		Avatar:         T.boardInfo.BoardName,
		DiscussID:      T.boardInfo.DiscussID,
		BbsID:          T.bbsID,
		FeedID:         T.feedID,
		Title:          T.bbsInfo.Title,
		Description:    T.bbsInfo.Description,
		Thumb:          "",
		UserID:         T.bbsInfo.UserID,
		Type:           T.bbsInfo.Type,
		Category:       T.bbsInfo.Category,
		CommentEnabled: T.bbsInfo.CommentEnabled,
		CreatedAt:      T.nowTime,
	}, feed.CustomizeTasker{
		EndTime:      T.bbsTaskInfo.EndTime,
		AllowExpired: T.bbsTaskInfo.AllowExpired,
		Status:       T.status,
	})
	if err != nil {
		return err
	}
	data3, err := json.Marshal(feedData)
	if err != nil {
		return err
	}
	msg := message.NewBroadcastMsg(T.bbsInfo.SiteID, message.FEED_MSG)
	msg.PackHead()
	msg.CustomizedMsg(`{"action":null}`, "", string(data3), "")
	return msg.Send(T.userIDs)
}
