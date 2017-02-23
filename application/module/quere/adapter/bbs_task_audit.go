package adapter

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/ccyun/GoApp/application/library/httpcurl"
	"github.com/ccyun/GoApp/application/model"
	"github.com/ccyun/GoApp/application/module/feed"
)

//TaskAudit 广播任务审核
type TaskAudit struct {
	base
	status int8
}

func init() {
	Register("taskAudit", new(TaskAudit))
}

//NewTask 新任务对象
func (T *TaskAudit) NewTask(task model.Queue) error {
	T.base.NewTask(task)
	var action map[string]string
	if err := json.Unmarshal([]byte(T.action), &action); err != nil {
		return fmt.Errorf("NewTask action Unmarshal error,taskID:%d,action:%s", T.taskID, T.action)
	}
	bbsID, err := strconv.Atoi(action["bbs_id"])
	if err != nil {
		return fmt.Errorf("NewTask strconv.Atoi bbs_id error,taskID:%d,action:%s", T.taskID, T.action)
	}
	T.bbsID = uint64(bbsID)
	status, err := strconv.Atoi(action["status"])
	if err != nil {
		return fmt.Errorf("NewTask strconv.Atoi status error,taskID:%d,action:%s", T.taskID, T.action)
	}
	T.status = int8(status)

	userID, err := strconv.Atoi(action["user_id"])
	if err != nil {
		return fmt.Errorf("NewTask strconv.Atoi user_id error,taskID:%d,action:%s", T.taskID, T.action)
	}
	T.userIDs = []uint64{uint64(userID)}

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

// //GetPublishScopeUsers 分析发布范围
// func (T *TaskAudit) GetPublishScopeUsers() error {
// 	return nil
// }

//CreateFeed 创建Feed
func (T *TaskAudit) CreateFeed() error {
	feedData := model.Feed{
		SiteID:    T.siteID,
		BoardID:   T.boardID,
		BbsID:     T.bbsID,
		FeedType:  "taskAudit",
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
		Status:         T.status,
	}
	dataByte, err := json.Marshal(data)
	if err != nil {
		return err
	}
	feedData.Data = string(dataByte)
	feedID, err := new(model.Feed).CreateFeed(feedData)
	if err == nil {
		T.feedID = feedID
	}
	return err
}

//CreateRelation 创建接收者关系
func (T *TaskAudit) CreateRelation() error {
	feedData := model.Feed{
		ID:       T.feedID,
		BoardID:  T.boardID,
		BbsID:    T.bbsID,
		FeedType: "taskReply",
	}
	return new(model.Feed).SaveHbase(T.userIDs, feedData)
}

//CreateUnread 创建未读计数
// func (T *TaskReply) CreateUnread() error {
// 	return T.base.CreateUnread()
// }

//UpdateStatus 更新状态及接收者用户列表
//无
// func (T *TaskAudit) UpdateStatus() error {
// 	return nil
// }

//SendMsg 发送消息
func (T *TaskAudit) SendMsg() error {
	feedData, err := feed.NewTask("taskReply", feed.Customizer{
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
		Status:       T.status,
	})
	if err != nil {
		return err
	}
	uc := new(httpcurl.UC)
	data := httpcurl.CustomizedSender{
		SiteID:      T.siteID,
		ToUsers:     T.bbsInfo.PublishScope.UserIDs,
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
