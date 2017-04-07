package adapter

import (
	"encoding/json"
	"strconv"
	"time"

	"fmt"

	"bbs_server/application/library/httpcurl"
	"bbs_server/application/model"
	"bbs_server/application/module/feed"
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
		BbsID             string   `json:"bbs_id"`
		AttachmentsBase64 string   `json:"attachments_base64"`
		DiscussMemberList []uint64 `json:"discuss_member_list"`
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
	if len(action.DiscussMemberList) > 0 {
		B.userIDs = action.DiscussMemberList
	}

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

	return nil
}

//GetPublishScopeUsers 分析发布范围
func (B *Bbs) GetPublishScopeUsers() error {
	if B.boardInfo.DiscussID != 0 {
		return nil
	}
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
	B.PublishScope = make(map[string][]uint64)
	B.PublishScope["group_ids"] = B.bbsInfo.PublishScope.GroupIDs
	B.PublishScope["user_ids"] = B.bbsInfo.PublishScope.UserIDs
	B.userIDs = append(B.bbsInfo.PublishScope.UserIDs, userIDs...)
	if len(B.bbsInfo.PublishScope.UserIDs) > 0 {
		B.PublishScopeuserLoginNames, err = new(httpcurl.UMS).GetUsersLoginName(B.customerCode, B.bbsInfo.PublishScope.UserIDs, true)
	}
	return err
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
		UserID:         B.bbsInfo.UsesID,
		Type:           B.bbsInfo.Type,
		Category:       B.category,
		CommentEnabled: B.bbsInfo.CommentEnabled,
	}
	switch B.category {
	case "bbs":
		data.Thumb = B.bbsInfo.Attachments[0]["url"]
	case "task":
		data.EndTime = B.bbsTaskInfo.EndTime
		data.AllowExpired = B.bbsTaskInfo.AllowExpired
		data.Status = -1
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

//CreateRelation 创建接收者关系
func (B *Bbs) CreateRelation() error {
	if B.boardInfo.DiscussID > 0 && B.bbsInfo.Type == "preview" {
		return nil
	}
	feedData := model.Feed{
		ID:       B.feedID,
		BoardID:  B.boardID,
		BbsID:    B.bbsID,
		FeedType: B.category,
	}
	return new(model.Feed).SaveHbase(B.userIDs, feedData, B.boardInfo.DiscussID)

}

//CreateTodo 创建未处理数
func (B *Bbs) CreateTodo() error {
	userIDs := []uint64{}
	if B.boardInfo.DiscussID > 0 {
		if B.bbsInfo.Type == "preview" {
			return nil
		}
		for _, u := range B.userIDs {
			if u != B.bbsInfo.UsesID {
				userIDs = append(userIDs, u)
			}
		}
	} else {
		userIDs = B.userIDs
	}
	if len(userIDs) == 0 {
		return nil
	}
	var data []model.Todo
	for _, userID := range userIDs {
		data = append(data, model.Todo{
			SiteID:   B.siteID,
			Type:     "unread",
			BoardID:  B.boardID,
			BbsID:    B.bbsID,
			FeedID:   B.feedID,
			FeedType: B.category,
			UserID:   userID,
		})
		if B.category == "task" && B.bbsInfo.Type == "default" {
			data = append(data, model.Todo{
				SiteID:   B.siteID,
				Type:     "unreply",
				BoardID:  B.boardID,
				BbsID:    B.bbsID,
				FeedID:   B.feedID,
				FeedType: B.category,
				UserID:   userID,
			})
		}
	}
	_, err := B.o.InsertMulti(100000, data)
	return err
}

//UpdateStatus 更新状态及接收者用户列表
//更新BBS状态及接收者总数及列表
func (B *Bbs) UpdateStatus() error {
	data := model.Bbs{
		ID:         B.bbsID,
		Status:     1,
		MsgCount:   uint64(len(B.userIDs)),
		ModifiedAt: uint64(time.Now().UnixNano() / 1e6),
	}
	u, err := json.Marshal(B.userIDs)
	if err != nil {
		return nil
	}
	data.PublishScopeUserIDs = string(u)
	_, err = B.o.Update(&data, "Status", "MsgCount", "PublishScopeUserIDs", "ModifiedAt")
	return err
}

//SendMsg 发送消息
func (B *Bbs) SendMsg() error {
	if B.boardInfo.DiscussID != 0 {
		return B.discussMsg()
	}
	switch B.category {
	case "bbs":
		return B.oaMsg()
	case "task", "form":
		return B.customizedMsg()
	}
	return nil
}

//getFeedCustomizer 定制数据
func (B *Bbs) getFeedCustomizer() feed.Customizer {
	thumb := ""
	if B.category == "bbs" {
		thumb = B.bbsInfo.Attachments[0]["url"]
		if B.boardInfo.DiscussID > 0 && B.attachmentsBase64 != "" {
			thumb = B.attachmentsBase64
		}
	}
	return feed.Customizer{
		BoardID:        B.boardID,
		BoardName:      B.boardInfo.BoardName,
		Avatar:         B.boardInfo.BoardName,
		DiscussID:      B.boardInfo.DiscussID,
		BbsID:          B.bbsID,
		FeedID:         B.feedID,
		Title:          B.bbsInfo.Title,
		Description:    B.bbsInfo.Description,
		Thumb:          thumb,
		UserID:         B.bbsInfo.UsesID,
		Type:           B.bbsInfo.Type,
		Category:       B.bbsInfo.Category,
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
	uc := new(httpcurl.UC)

	data := httpcurl.OASender{
		SiteID: B.siteID,
		Title:  B.bbsInfo.Title,
		TitleElements: []httpcurl.OASendTitleElementser{
			httpcurl.OASendTitleElementser{Title: B.bbsInfo.Title},
		},
		DetailURL: feedData.DetailURL,
		Elements: []httpcurl.OASendElementser{
			httpcurl.OASendElementser{ImageID: B.bbsInfo.Attachments[0]["url"]},
			httpcurl.OASendElementser{Content: description},
		},
		ToUsers:    B.PublishScopeuserLoginNames,
		ToPartyIds: B.bbsInfo.PublishScope.GroupIDs,
	}
	data.CustomizedData = feedData.CustomizedData
	return uc.OASend(data)
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
	uc := new(httpcurl.UC)
	data := httpcurl.CustomizedSender{
		SiteID:      strconv.FormatUint(B.siteID, 10),
		ToUsers:     B.PublishScopeuserLoginNames,
		ToPartyIds:  B.bbsInfo.PublishScope.GroupIDs,
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

//discussMsg 讨论组消息
func (B *Bbs) discussMsg() error {
	feedData, err := feed.NewBbs(B.bbsInfo.Category, B.getFeedCustomizer())
	if err != nil {
		return err
	}
	ucc := new(httpcurl.UCC)
	postData := httpcurl.UCCMsgSender{
		SiteID:           B.siteID,
		UserID:           B.bbsInfo.UsesID,
		ConversationType: 2,
	}
	postData.Message.Content = ""
	postData.Control.NoSendself = 0
	postData.To.ToID = B.boardInfo.DiscussID
	postData.To.ToPrivateIDs = []uint64{}
	if B.bbsInfo.Type == "preview" {
		if B.bbsInfo.UsesID == B.bbsInfo.PublishScope.UserIDs[0] {
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
		return nil
	}
	postData.Message.Content = string(contentByte)
	_, err = ucc.MsgSend(postData)
	return err
}
