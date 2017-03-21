package adapter

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/ccyun/GoApp/application/library/httpcurl"
	"github.com/ccyun/GoApp/application/model"
	"github.com/ccyun/GoApp/application/module/feed"
)

//TaskClose 广播任务提醒反馈
type TaskClose struct {
	base
}

func init() {
	Register("TaskClose", new(TaskClose))
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
	return nil
}

//GetPublishScopeUsers 分析发布范围
func (T *TaskClose) GetPublishScopeUsers() error {
	var err error
	T.userIDs = T.bbsInfo.PublishScopeUserIDsArr
	T.userIDs = append(T.userIDs, T.boardInfo.EditorIDs...)
	T.userLoginNames, err = new(httpcurl.UMS).GetUsersLoginName(T.customerCode, T.userIDs, true)
	return err
}

//CreateFeed 创建Feed
func (T *TaskClose) CreateFeed() error {
	feedData := model.Feed{
		SiteID:    T.siteID,
		BoardID:   T.boardID,
		BbsID:     T.bbsID,
		FeedType:  "taskClose",
		CreatedAt: T.bbsInfo.CreatedAt,
	}
	data := model.FeedData{
		Title:          T.bbsInfo.Title,
		Description:    T.bbsInfo.Description,
		CreatedAt:      T.bbsInfo.CreatedAt,
		UserID:         T.bbsInfo.UsesID,
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

//CreateRelation 创建接收者关系
func (T *TaskClose) CreateRelation() error {
	feedData := model.Feed{
		ID:       T.feedID,
		BoardID:  T.boardID,
		BbsID:    T.bbsID,
		FeedType: "taskClose",
	}
	return new(model.Feed).SaveHbase(T.userIDs, feedData, T.boardInfo.DiscussID)
}

//SendMsg 发送消息
func (T *TaskClose) SendMsg() error {
	feedData, err := feed.NewTask("taskClose", feed.Customizer{
		BoardID:        T.boardID,
		BoardName:      T.boardInfo.BoardName,
		Avatar:         T.boardInfo.BoardName,
		DiscussID:      T.boardInfo.DiscussID,
		BbsID:          T.bbsID,
		FeedID:         T.feedID,
		Title:          T.bbsInfo.Title,
		Description:    T.bbsInfo.Description,
		Thumb:          "",
		UserID:         T.bbsInfo.UsesID,
		Type:           T.bbsInfo.Type,
		Category:       T.bbsInfo.Category,
		CommentEnabled: T.bbsInfo.CommentEnabled,
		CreatedAt:      T.bbsInfo.CreatedAt,
	}, feed.CustomizeTasker{
		EndTime:      T.bbsTaskInfo.EndTime,
		AllowExpired: T.bbsTaskInfo.AllowExpired,
		Status:       feed.BbsTaskCloseStatus,
	})
	if err != nil {
		return err
	}
	uc := new(httpcurl.UC)
	data := httpcurl.CustomizedSender{
		SiteID:      T.siteID,
		ToUsers:     T.userLoginNames,
		ToPartyIds:  T.bbsInfo.PublishScope.GroupIDs,
		WebPushData: "您有一个“i 广播”消息",
	}
	data3, err := json.Marshal(feedData)
	if err != nil {
		return err
	}
	data.Data3 = string(data3)
	return uc.CustomizedSend(data)
}
