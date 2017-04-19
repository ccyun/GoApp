package adapter

import (
	"encoding/json"
	"fmt"
	"strconv"

	"time"

	"bbs_server/application/function"
	"bbs_server/application/library/httpcurl"
	"bbs_server/application/model"
	"bbs_server/application/module/feed"
)

//TaskReply 广播任务提醒反馈
type TaskReply struct {
	base
}

func init() {
	Register("taskReply", func() Tasker {
		return new(TaskReply)
	})
}

//NewTask 新任务对象
func (T *TaskReply) NewTask(task model.Queue) error {
	T.base.NewTask(task)
	var action map[string]string
	if err := json.Unmarshal([]byte(T.action), &action); err != nil {
		return fmt.Errorf("NewTask action Unmarshal error,taskID:%d,action:%s", T.taskID, T.action)
	}
	bbsID, err := strconv.Atoi(action["bbs_id"])
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
	return nil
}

//GetPublishScopeUsers 分析发布范围
func (T *TaskReply) GetPublishScopeUsers() error {
	userIDs, err := new(model.BbsTaskReply).GetReplyUserIDs(T.bbsID)
	if err != nil {
		return err
	}
	T.userIDs = function.SliceDiff(T.bbsInfo.PublishScopeUserIDsArr, userIDs).Uint64()
	if len(T.userIDs) > 0 {
		T.PublishScopeuserLoginNames, err = new(httpcurl.UMS).GetUsersLoginName(T.customerCode, T.userIDs, true)
		return err
	}
	return fmt.Errorf("GetPublishScopeUsers error not found UNReply Users")
}

//CreateFeed 创建Feed
func (T *TaskReply) CreateFeed() error {
	feedData := model.Feed{
		SiteID:    T.siteID,
		BoardID:   T.boardID,
		BbsID:     T.bbsID,
		FeedType:  feed.FeedTypeTaskReply,
		CreatedAt: uint64(time.Now().UnixNano() / 1000000),
	}
	data := model.FeedData{
		Title:          T.bbsInfo.Title,
		Description:    T.bbsInfo.Description,
		CreatedAt:      uint64(time.Now().UnixNano() / 1000000),
		UserID:         T.bbsInfo.UserID,
		Type:           T.bbsInfo.Type,
		Category:       T.category,
		CommentEnabled: T.bbsInfo.CommentEnabled,
		EndTime:        T.bbsTaskInfo.EndTime,
		AllowExpired:   T.bbsTaskInfo.AllowExpired,
		Status:         feed.BbsTaskReplyStatus,
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

//CreateRelation 创建接收者关系
func (T *TaskReply) CreateRelation() error {
	feedData := model.Feed{
		ID:       T.feedID,
		BoardID:  T.boardID,
		BbsID:    T.bbsID,
		FeedType: feed.FeedTypeTaskReply,
	}
	return new(model.Feed).SaveHbase(T.userIDs, feedData, T.boardInfo.DiscussID)
}

//SendMsg 发送消息
func (T *TaskReply) SendMsg() error {
	feedData, err := feed.NewTask(feed.FeedTypeTaskReply, feed.Customizer{
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
		CreatedAt:      uint64(time.Now().UnixNano() / 1000000),
	}, feed.CustomizeTasker{
		EndTime:      T.bbsTaskInfo.EndTime,
		AllowExpired: T.bbsTaskInfo.AllowExpired,
		Status:       -1,
	})
	if err != nil {
		return err
	}
	uc := new(httpcurl.UC)
	data := httpcurl.CustomizedSender{
		SiteID:      strconv.FormatUint(T.siteID, 10),
		ToUsers:     T.PublishScopeuserLoginNames,
		ToPartyIds:  T.bbsInfo.PublishScope.GroupIDs,
		WebPushData: "您有一个“i 广播”消息",
	}
	data.Data1 = "{\"action\":null}"
	data3, err := json.Marshal(feedData)
	if err != nil {
		return err
	}
	data.Data3 = string(data3)
	return uc.CustomizedSend(data)
}
