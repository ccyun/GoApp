package model

import (
	"strconv"

	"github.com/ccyun/GoApp/application/function"
	"github.com/ccyun/GoApp/application/library/hbase"
)

//Feed 任务表结构
type Feed struct {
	ID        uint64 `orm:"column(id)"`
	SiteID    uint64 `orm:"column(site_id)"`
	BoardID   uint64 `orm:"column(board_id)"`
	BbsID     uint64 `orm:"column(bbs_id)"`
	FeedType  string `orm:"column(feed_type)"`
	Data      string `orm:"column(data)"`
	MsgID     uint64 `orm:"column(msg_id)"`
	CreatedAt uint64 `orm:"column(created_at)"`
}

//FeedBbs 图文广播feed
type FeedBbs struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	CreatedAt      string `json:"created_at"`
	UserID         string `json:"user_id"`
	Type           string `json:"type"`
	Category       string `json:"category"`
	CommentEnabled uint8  `json:"comment_enabled"`
}

//FeedForm 表单feed
type FeedForm struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	CreatedAt      string `json:"created_at"`
	UserID         string `json:"user_id"`
	Type           string `json:"type"`
	Category       string `json:"category"`
	CommentEnabled uint8  `json:"comment_enabled"`
}

//FeedTask 广播任务feed
type FeedTask struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	CreatedAt      string `json:"created_at"`
	UserID         string `json:"user_id"`
	Type           string `json:"type"`
	Category       string `json:"category"`
	CommentEnabled uint8  `json:"comment_enabled"`
	EndTime        uint64 `json:"end_time"`
	AllowExpired   uint8  `json:"allow_expired"`
}

//FeedTaskReply 广播任务提醒
type FeedTaskReply struct {
	Title          string `json:"title"`
	CreatedAt      string `json:"created_at"`
	Status         uint8  `json:"status"`
	Category       string `json:"category"`
	CommentEnabled uint8  `json:"comment_enabled"`
	EndTime        uint64 `json:"end_time"`
	AllowExpired   uint8  `json:"allow_expired"`
}

//FeedTaskAudit 广播任务审核
type FeedTaskAudit struct {
	Title          string `json:"title"`
	CreatedAt      string `json:"created_at"`
	Status         uint8  `json:"status"`
	Category       string `json:"category"`
	CommentEnabled uint8  `json:"comment_enabled"`
	EndTime        uint64 `json:"end_time"`
	AllowExpired   uint8  `json:"allow_expired"`
}

//FeedTaskClose 广播任务关闭
type FeedTaskClose struct {
	Title          string `json:"title"`
	CreatedAt      string `json:"created_at"`
	Category       string `json:"category"`
	Status         uint8  `json:"status"`
	CommentEnabled uint8  `json:"comment_enabled"`
	EndTime        uint64 `json:"end_time"`
	AllowExpired   uint8  `json:"allow_expired"`
}

//TableName 表名
func (F *Feed) TableName() string {
	return "feed"
}

//HbaseTableName hbase表名
func (F *Feed) HbaseTableName() string {
	return "bbs_feed"
}

//SaveHbase 保存数据到hbase
func (F *Feed) SaveHbase(userIDs []uint64, feedData Feed) error {
	client, err := hbase.OpenClient()
	defer hbase.CloseClient(client)
	if err != nil {
		return err
	}
	var (
		boardID, bbsID, family []byte
		data                   []*hbase.TPut
		timeStamp              int64
	)
	boardID = []byte(strconv.FormatUint(feedData.BoardID, 10))
	bbsID = []byte(strconv.FormatUint(feedData.BbsID, 10))
	family = []byte("cf")
	timeStamp = int64(feedData.ID)
	for _, u := range userIDs {
		rowkey := function.MakeRowkey(int64(u))
		data = append(data, &hbase.TPut{
			Row: []byte(rowkey + "_home"),
			ColumnValues: []*hbase.TColumnValue{
				&hbase.TColumnValue{
					Family:    family,
					Qualifier: boardID,
					Value:     bbsID,
					Timestamp: &timeStamp,
				},
			},
		})
		if feedData.FeedType == "bbs" || feedData.FeedType == "task" || feedData.FeedType == "form" {
			data = append(data, &hbase.TPut{
				Row: []byte(rowkey + "_list"),
				ColumnValues: []*hbase.TColumnValue{
					&hbase.TColumnValue{
						Family:    family,
						Qualifier: boardID,
						Value:     bbsID,
						Timestamp: &timeStamp,
					},
				},
			})
		}
	}
	return client.PutMultiple([]byte(F.HbaseTableName()), data)
}
