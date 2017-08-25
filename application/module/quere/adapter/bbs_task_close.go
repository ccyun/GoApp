package adapter

import (
	"encoding/json"
	"fmt"
	"strconv"

	"bbs_server/application/model"
	"bbs_server/application/module/feed"
	"bbs_server/application/module/message"
)

//TaskClose 广播任务提醒反馈
type TaskClose struct {
	base
}

func init() {
	Register("taskClose", func() Tasker {
		return new(TaskClose)
	})
}

//NewTask 新任务对象
func (T *TaskClose) NewTask(task model.Queue) error {
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
func (T *TaskClose) GetPublishScopeUsers() error {
	T.userIDs = new(model.Msg).GetUserIDs(T.siteID, T.boardID, T.bbsID, -1)
	return T.base.GetPublishScopeUsers()
}

//CreateFeed 创建Feed
func (T *TaskClose) CreateFeed() error {
	feedData := model.Feed{
		SiteID:    T.siteID,
		BoardID:   T.boardID,
		BbsID:     T.bbsID,
		FeedType:  feed.FeedTypeTaskClose,
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
		Status:         feed.BbsTaskCloseStatus,
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
func (T *TaskClose) SendMsg() error {
	feedData, err := feed.NewTask(feed.FeedTypeTaskClose, feed.Customizer{
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
		Status:       feed.BbsTaskCloseStatus,
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
