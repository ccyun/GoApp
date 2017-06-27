package feed

import (
	"encoding/json"
	"fmt"
)

var (
	//AppDomain url
	AppDomain string
	//Path 相对路径
	Path string
)

//Version feed版本
const Version uint64 = 1

//DetailAuth url请求需带上验证信息
const DetailAuth uint8 = 1

const (
	//BbsTaskStatus 任务发布
	BbsTaskStatus int8 = -1
	//BbsTaskCloseStatus 任务关闭
	BbsTaskCloseStatus int8 = 4
	//BbsTaskReplyStatus //提醒反馈
	BbsTaskReplyStatus int8 = -1
	//BbsTaskAuditPassed 审核通过
	BbsTaskAuditPassed int8 = 1
	//BbsTaskAuditNotPassed 审核不通过
	BbsTaskAuditNotPassed int8 = 2
	//BbsTaskTimeout 审核不通过
	BbsTaskTimeout int8 = 3
)
const (
	//FeedTypeBbs feed type bbs 图文广播
	FeedTypeBbs = "bbs"
	//FeedTypeForm feed type form 表单
	FeedTypeForm = "form"
	//FeedTypeTask feed type task 广播任务
	FeedTypeTask = "task"
	//FeedTypeTaskReply feed type taskReply 广播任务提醒反馈
	FeedTypeTaskReply = "taskReply"
	//FeedTypeTaskAudit feed type taskAudit 广播任务审核
	FeedTypeTaskAudit = "taskAudit"
	//FeedTypeTaskClose feed type taskClose 广播任务关闭任务
	FeedTypeTaskClose = "taskClose"
)

//Feeder 结构体
type Feeder struct {
	Version        uint64 `json:"version"`
	BoardID        uint64 `json:"board_id"`
	BbsID          uint64 `json:"bbs_id"`
	FeedType       string `json:"feed_type"`
	FeedID         uint64 `json:"feed_id"`
	DetailAuth     uint8  `json:"detailAuth"`
	DetailURL      string `json:"detailURL"`
	DisplayType    string `json:"displayType"`
	CustomizedData string `json:"customizedData"`
	Elements       string `json:"elements"`
}

//Customizer 定制数据
type Customizer struct {
	BoardID        uint64 `json:"board_id"`
	BoardName      string `json:"board_name"`
	Avatar         string `json:"avatar"`
	DiscussID      uint64 `json:"discuss_id"`
	BbsID          uint64 `json:"bbs_id"`
	FeedID         uint64 `json:"feed_id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Thumb          string `json:"thumb"`
	UserID         uint64 `json:"user_id"`
	Type           string `json:"type"`
	Category       string `json:"category"`
	Link           string `json:"link"`
	IsBrowser      uint8  `json:"is_browser"`
	IsAuth         uint8  `json:"is_auth"`
	CommentEnabled uint8  `json:"comment_enabled"`
	CreatedAt      uint64 `json:"created_at"`
}

//CustomizeTasker 任务定制数据
type CustomizeTasker struct {
	EndTime      uint64 `json:"end_time"`
	AllowExpired uint8  `json:"allow_expired"`
	Status       int8   `json:"status"`
}

//Bbs 任务定制数据
type Bbs struct {
	Customizer
}

//Task 任务定制数据
type Task struct {
	Customizer
	CustomizeTasker
}

//Form 任务定制数据
type Form struct {
	Customizer
}

//Handle 定制消息处理
func (C *Customizer) Handle(data Customizer) {
	C.BoardID = data.BoardID
	C.BoardName = data.BoardName
	C.Avatar = data.Avatar
	C.DiscussID = data.DiscussID
	C.BbsID = data.BbsID
	C.FeedID = data.FeedID
	C.Title = data.Title
	C.Description = data.Description
	C.Thumb = data.Thumb
	C.UserID = data.UserID
	C.Type = data.Type
	C.Category = data.Category
	C.CommentEnabled = data.CommentEnabled
	C.CreatedAt = data.CreatedAt
}

//Init 初始化配置
func Init(option map[string]string) error {
	//App 域名
	appDomain, ok := option["app_domain"]
	if !ok {
		return fmt.Errorf(`config "app_domain" is undefined`)
	}
	AppDomain = appDomain
	//App 路径
	path, ok := option["app_path"]
	if !ok {
		return fmt.Errorf(`config "path" is undefined`)
	}
	Path = path
	return nil
}

//newFeed 创建新的Feed
func newFeed(feedType string, data Customizer) Feeder {
	feedData := Feeder{
		Version:    Version,
		BoardID:    data.BoardID,
		BbsID:      data.BbsID,
		FeedType:   feedType,
		FeedID:     data.FeedID,
		DetailAuth: DetailAuth,
	}
	feedData.DisplayType = "WebView"
	if feedType != "bbs" && feedType != "task" {
		feedData.DisplayType = "RichMedia"
	}
	if data.Link != "" {
		feedData.DetailURL = data.Link
	} else {
		feedData.DetailURL = fmt.Sprintf("%s%s/bbs/show/bbs.html?id=%d&category=%s&v=2", AppDomain, Path, data.BbsID, data.Category)
	}
	return feedData
}

//NewBbs 处理bbs 图文广播
func NewBbs(feedType string, data Customizer) (Feeder, error) {
	feedData := newFeed(feedType, data)
	customized := new(Bbs)
	customized.Handle(data)
	customizedData, err := json.Marshal(customized)
	if err != nil {
		return feedData, err
	}
	feedData.CustomizedData = string(customizedData)
	feedData.Elements = GetBbsView(customized)
	return feedData, nil
}

//NewTask 处理task 广播任务
func NewTask(feedType string, data Customizer, extData CustomizeTasker) (Feeder, error) {
	feedData := newFeed(feedType, data)
	customized := new(Task)
	customized.Handle(data)
	customized.AllowExpired = extData.AllowExpired
	customized.EndTime = extData.EndTime
	customized.Status = extData.Status
	customizedData, err := json.Marshal(customized)
	if err != nil {
		return feedData, err
	}
	feedData.CustomizedData = string(customizedData)
	switch feedType {
	case "task":
		feedData.Elements = GetTaskView(customized)
	case "taskReply":
		feedData.Elements = GetTaskReplyView(customized)
	case "taskAudit":
		feedData.Elements = GetTaskAuditView(customized)
	case "taskClose":
		feedData.Elements = GetTaskCloseView(customized)
	}
	return feedData, nil
}

//NewFrom 处理BBS 广播表单
func NewFrom(feedType string, data Customizer) (Feeder, error) {
	feedData := newFeed(feedType, data)
	customized := new(Form)
	customized.Handle(data)
	customizedData, err := json.Marshal(customized)
	if err != nil {
		return feedData, err
	}
	feedData.CustomizedData = string(customizedData)
	feedData.Elements = GetFromView(customized)
	return feedData, nil
}
