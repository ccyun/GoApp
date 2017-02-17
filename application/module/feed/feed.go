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
	//FeedIcons feed模板图标
	FeedIcons string
)

//Feeder 结构体
type Feeder struct {
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
	CommentEnabled uint8  `json:"comment_enabled"`
	CreatedAt      uint64 `json:"created_at"`
}

//CustomizeTasker 任务定制数据
type CustomizeTasker struct {
	EndTime      uint64 `json:"end_time"`
	AllowExpired uint8  `json:"allow_expired"`
	Status       uint8  `json:"status"`
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
	path, ok := option["Path"]
	if !ok {
		return fmt.Errorf(`config "path" is undefined`)
	}
	Path = path
	//feed 模板图标
	feedIcons, ok := option["feed_icons"]
	if !ok {
		return fmt.Errorf(`config "feed_icons" is undefined`)
	}
	FeedIcons = feedIcons

	return nil
}

//newFeed 创建新的Feed
func newFeed(boardID uint64, feedType string, data Customizer) Feeder {
	feedData := Feeder{
		BoardID:    boardID,
		BbsID:      data.BbsID,
		FeedType:   feedType,
		FeedID:     data.FeedID,
		DetailAuth: 1,
	}
	feedData.DisplayType = "WebView"
	if feedType != "bbs" && feedType != "task" {
		feedData.DisplayType = "RichMedia"
	}
	feedData.DetailURL = fmt.Sprintf("%s%s/bbs/show/bbs.html?id=%d&category=%s&v=2", AppDomain, Path, data.BbsID, data.Category)
	return feedData
}

//NewBbs 处理bbs 图文广播
func NewBbs(boardID uint64, feedType string, data Customizer) (Feeder, error) {

	feedData := newFeed(boardID, feedType, data)
	customizedData, err := json.Marshal(data)
	if err != nil {
		return feedData, err
	}
	feedData.CustomizedData = string(customizedData)
	feedData.Elements = GetBbsView(data)
	return feedData, nil
}

//NewTask 处理task 广播任务
func NewTask(boardID uint64, feedType string, data Customizer, extData CustomizeTasker) (Feeder, error) {
	feedData := newFeed(boardID, feedType, data)
	customized := struct {
		Customizer
		CustomizeTasker
	}{}
	customized = data
	customizedData, err := json.Marshal(data)
	if err != nil {
		return feedData, err
	}
	feedData.CustomizedData = string(customizedData)
	feedData.Elements = GetBbsView(data)
	return feedData, nil
}

//NewFrom 处理BBS 广播表单
func NewFrom(BoardID uint64, feedType string, data Customizer) (Feeder, error) {
	return Feeder{}, nil
}
