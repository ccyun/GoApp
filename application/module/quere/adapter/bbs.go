package adapter

import (
	"encoding/json"
	"strconv"

	"fmt"

	"bbs_server/application/function"
	"bbs_server/application/library/httpcurl"
	"bbs_server/application/model"
	"bbs_server/application/module/feed"
	"bbs_server/application/module/message"
)

//Bbs 图文广播
type Bbs struct {
	base
}

//OACustomizedDataer OA消息定制数据
type OACustomizedDataer struct {
	BoardID        uint64 `json:"board_id"`
	BoardName      string `json:"board_name"`
	Avatar         string `json:"avatar"`
	BbsID          uint64 `json:"bbs_id"`
	FeedID         uint64 `json:"feed_id"`
	Title          string `json:"title"`
	CreatedAt      uint64 `json:"created_at"`
	Category       string `json:"category"`
	Type           string `json:"type"`
	CommentEnabled uint8  `json:"comment_enabled"`
}

func init() {
	Register("bbs", func() Tasker {
		return new(Bbs)
	})
}

//getBbsTaskInfo 读取广播任务信息
func (B *base) getBbsTaskInfo() error {
	var err error
	model := new(model.BbsTask)
	B.bbsTaskInfo, err = model.GetOne(B.bbsID)
	return err
}

//NewTask 新任务对象
func (B *Bbs) NewTask(task model.Queue) error {
	B.base.NewTask(task)

	var action struct {
		BbsID             string `json:"bbs_id"`
		AttachmentsBase64 string `json:"attachments_base64"`
	}
	if err := json.Unmarshal([]byte(B.action), &action); err != nil {
		return fmt.Errorf("NewTask action Unmarshal error,taskID:%d,action:%s", B.taskID, B.action)
	}
	bbsID, err := strconv.Atoi(action.BbsID)
	if err != nil {
		return fmt.Errorf("NewTask strconv.Atoi error,taskID:%d,action:%s", B.taskID, B.action)
	}
	B.bbsID = uint64(bbsID)
	B.attachmentsBase64 = action.AttachmentsBase64
	if err := B.getBbsInfo(); err != nil {
		return err
	}
	if err := B.getBoardInfo(); err != nil {
		return err
	}
	///////判断广播类型
	switch B.category {
	case "task":
		if err := B.getBbsTaskInfo(); err != nil {
			return err
		}
	}
	B.feedType = B.category
	return nil
}

//GetPublishScopeUsers 分析发布范围
func (B *Bbs) GetPublishScopeUsers() error {
	if B.boardInfo.DiscussID != 0 {
		if B.bbsInfo.Type == "preview" {
			B.userIDs = B.bbsInfo.PublishScope.UserIDs
		} else {
			discussInfo, err := new(httpcurl.UCC).GetDiscussInfo(B.bbsInfo.UserID, B.boardInfo.DiscussID)
			if err != nil {
				return err
			}
			B.userIDs = discussInfo.ValidMemberIDs
		}

	} else {
		var (
			userIDs []uint64
			err     error
		)
		if len(B.bbsInfo.PublishScope.GroupIDs) > 0 {
			ums := new(httpcurl.UMS)
			userIDs, err = ums.GetAllUserIDsByOrgIDs(B.customerCode, B.bbsInfo.PublishScope.GroupIDs)
			if err != nil {
				return err
			}
		}
		B.userIDs = function.SliceUnique(append(B.bbsInfo.PublishScope.UserIDs, userIDs...)).Uint64()
	}
	return B.base.GetPublishScopeUsers()
}

//CreateFeed 创建Feed
func (B *Bbs) CreateFeed() error {
	if B.boardInfo.DiscussID > 0 && B.bbsInfo.Type == "preview" {
		return nil
	}
	feedData := model.Feed{
		SiteID:    B.siteID,
		BoardID:   B.boardID,
		BbsID:     B.bbsID,
		FeedType:  B.category,
		CreatedAt: B.bbsInfo.PublishAt,
	}
	data := model.FeedData{
		Title:          B.bbsInfo.Title,
		Description:    B.bbsInfo.Description,
		CreatedAt:      B.bbsInfo.PublishAt,
		UserID:         B.bbsInfo.UserID,
		Type:           B.bbsInfo.Type,
		Category:       B.category,
		CommentEnabled: B.bbsInfo.CommentEnabled,
		IsBrowser:      B.bbsInfo.IsBrowser,
		IsAuth:         B.bbsInfo.IsAuth,
	}
	switch B.category {
	case "bbs":
		data.Thumb = B.bbsInfo.Thumb
		data.Link = B.bbsInfo.Link
	case "task":
		data.EndTime = B.bbsTaskInfo.EndTime
		data.AllowExpired = B.bbsTaskInfo.AllowExpired
		data.Status = feed.BbsTaskStatus
	}
	dataByte, err := json.Marshal(data)
	if err != nil {
		return err
	}
	feedData.Data = string(dataByte)
	feedID, err := B.o.Insert(&feedData)
	if err == nil {
		B.feedID = uint64(feedID)
	}
	return err
}

//UpdateStatus 更新状态及接收者用户列表
//更新BBS状态及接收者总数及列表
func (B *Bbs) UpdateStatus() error {
	data := model.Bbs{
		ID:         B.bbsID,
		Status:     1,
		MsgCount:   uint64(len(B.userIDs)),
		ModifiedAt: B.nowTime,
	}
	if _, err := B.o.Update(&data, "Status", "MsgCount", "ModifiedAt"); err != nil {
		return err
	}
	return B.createQueue()
}

//创建队列
func (B *Bbs) createQueue() error {
	if B.category == "task" {
		data := model.Queue{
			SiteID:       B.siteID,
			CustomerCode: B.customerCode,
			BbsID:        B.bbsID,
			Action:       fmt.Sprintf(`{"bbs_id":%d}`, B.bbsID),
			Status:       0,
			TryCount:     0,
			SetTimer:     B.bbsTaskInfo.ReplyRemindAt,
		}
		queueData := []model.Queue{}
		if B.bbsTaskInfo.ReplyRemindAt > B.nowTime {
			data.TaskType = "taskReply"
			data.SetTimer = B.bbsTaskInfo.ReplyRemindAt
			queueData = append(queueData, data)
		}
		if B.bbsTaskInfo.AuditRemindAt > B.nowTime {
			data.TaskType = "taskAuditRemind"
			data.SetTimer = B.bbsTaskInfo.AuditRemindAt
			queueData = append(queueData, data)
		}
		if B.bbsTaskInfo.AllowExpired == 0 && B.bbsTaskInfo.EndTime > B.nowTime {
			data.TaskType = "taskEnd"
			data.SetTimer = B.bbsTaskInfo.EndTime
			queueData = append(queueData, data)
		}
		if len(queueData) > 0 {
			_, err := B.o.InsertMulti(2, queueData)
			return err
		}
	}
	return nil
}

//SendMsg 发送消息
func (B *Bbs) SendMsg() error {
	if B.boardInfo.DiscussID != 0 {
		return B.discussMsg()
	}
	switch B.category {
	case "bbs":
		return B.oaMsg()
	case "task":
		if err := new(httpcurl.BILL).Accepter(B.siteID, B.userIDs); err != nil { //计费
			return err
		}
		return B.customizedMsg()
	case "form":
		return B.customizedMsg()
	}
	return nil
}

//getFeedCustomizer 定制数据
func (B *Bbs) getFeedCustomizer() feed.Customizer {
	thumb := ""
	if B.category == "bbs" {
		thumb = B.bbsInfo.Thumb
		if B.boardInfo.DiscussID > 0 && B.attachmentsBase64 != "" {
			thumb = B.attachmentsBase64
		}
	}
	return feed.Customizer{
		BoardID:        B.boardID,
		BoardName:      B.boardInfo.BoardName,
		Avatar:         B.boardInfo.Avatar,
		DiscussID:      B.boardInfo.DiscussID,
		BbsID:          B.bbsID,
		FeedID:         B.feedID,
		Title:          B.bbsInfo.Title,
		Description:    B.bbsInfo.Description,
		Thumb:          thumb,
		UserID:         B.bbsInfo.UserID,
		Type:           B.bbsInfo.Type,
		Category:       B.bbsInfo.Category,
		Link:           B.bbsInfo.Link,
		IsBrowser:      B.bbsInfo.IsBrowser,
		IsAuth:         B.bbsInfo.IsAuth,
		CommentEnabled: B.bbsInfo.CommentEnabled,
		CreatedAt:      B.bbsInfo.CreatedAt,
	}
}

//oaMsg OA消息
func (B *Bbs) oaMsg() error {
	feedData, err := feed.NewBbs(B.bbsInfo.Category, B.getFeedCustomizer())
	if err != nil {
		return err
	}
	description := B.bbsInfo.Description
	if B.bbsInfo.Description == "" {
		description = B.boardInfo.BoardName
	}
	msg := message.NewBroadcastMsg(B.bbsInfo.SiteID, message.FEED_MSG)
	msg.PackHead()
	msg.OAMsg(
		[]message.OASendTitleElementser{
			message.OASendTitleElementser{Title: B.bbsInfo.Title},
		},
		feedData.CustomizedData,
		feedData.DetailURL,
		[]message.OASendElementser{
			message.OASendElementser{ImageID: B.bbsInfo.Thumb},
			message.OASendElementser{Content: description},
		},
	)
	return msg.Send(B.userIDs)
}

//customizedMsg 定制消息（任务）
func (B *Bbs) customizedMsg() error {
	feedData, err := feed.NewTask(B.bbsInfo.Category, B.getFeedCustomizer(), feed.CustomizeTasker{
		EndTime:      B.bbsTaskInfo.EndTime,
		AllowExpired: B.bbsTaskInfo.AllowExpired,
		Status:       feed.BbsTaskStatus,
	})
	if err != nil {
		return err
	}
	data3, err := json.Marshal(feedData)
	if err != nil {
		return err
	}
	msg := message.NewBroadcastMsg(B.bbsInfo.SiteID, message.FEED_MSG)
	msg.PackHead()
	msg.CustomizedMsg(`{"action":null}`, "", string(data3), "")
	return msg.Send(B.userIDs)
}

//discussMsg 讨论组消息
func (B *Bbs) discussMsg() error {
	feedData, err := feed.NewBbs(B.bbsInfo.Category, B.getFeedCustomizer())
	if err != nil {
		return err
	}
	ucc := new(httpcurl.UCC)
	postData := httpcurl.UCCMsgSender{
		SiteID:           B.siteID,
		UserID:           B.bbsInfo.UserID,
		ConversationType: 2,
	}
	postData.Message.Content = ""
	postData.Control.NoSendself = 0
	postData.To.ToID = B.boardInfo.DiscussID
	postData.To.ToPrivateIDs = []uint64{}
	if B.bbsInfo.Type == "preview" {
		if B.bbsInfo.UserID != B.bbsInfo.PublishScope.UserIDs[0] {
			postData.Control.NoSendself = 1
		}
		postData.To.ToPrivateIDs = B.bbsInfo.PublishScope.UserIDs
	}
	content := struct {
		Version        uint64 `json:"version"`
		Title          string `json:"title"`
		Content        string `json:"content"`
		DetailAuth     uint8  `json:"detailAuth"`
		DetailURL      string `json:"detailURL"`
		CustomizedData string `json:"customizedData"`
	}{
		Version:        1,
		Title:          B.bbsInfo.Title,
		Content:        feedData.Elements,
		DetailAuth:     1,
		DetailURL:      feedData.DetailURL,
		CustomizedData: feedData.CustomizedData,
	}
	contentByte, err := json.Marshal(content)
	if err != nil {
		return err
	}
	postData.Message.Content = string(contentByte)
	_, err = ucc.MsgSend(postData)
	return err
}
