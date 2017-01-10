package model

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego/logs"
	"github.com/ccyun/GoApp/application/function"
	"github.com/ccyun/GoApp/application/library/hbase"
)

//Feed 任务表结构
type Feed struct {
	ID        uint64 `orm:"column(id)"`
	SiteID    uint64 `orm:"column(site_id)"`
	BoardID   uint64 `orm:"column(board_id)"`
	BbsID     uint64 `orm:"column(bbs_id)"`
	FeedType  uint64 `orm:"column(feed_type)"`
	Data      uint64 `orm:"column(data)"`
	MsgID     uint64 `orm:"column(msg_id)"`
	CreatedAt uint64 `orm:"column(created_at)"`
}

//HbaseFeed hbase数据结构
type HbaseFeed struct {
	BoardID        uint64 `json:"board_id"`
	BbsID          uint64 `json:"bbs_id"`
	FeedID         uint64 `json:"feed_id"`
	FeedType       string `json:"feed_type"`
	MsgID          uint64 `json:"msg_id"`
	DiscussID      uint64 `json:"discuss_id"`
	UserID         uint64 `json:"user_id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Thumb          string `json:"thumb"`
	Category       string `json:"category"`
	Type           string `json:"type"`
	CommentEnabled uint64 `json:"comment_enabled"`
	CreatedAt      uint64 `json:"created_at"`
}

//MaxFeedNum feed容量
const (
	MaxFeedNum = 200000000000000
	MinFeedNum = 100000000000000
)

//TableName 表名
func (F *Feed) TableName() string {
	return "feed"
}

//HbaseTableName hbase表名
func (F *Feed) HbaseTableName() string {
	return "bbs_feed"
}

//SaveHbase 保存数据到hbase
func (F *Feed) SaveHbase(userIDs []uint64, feedData HbaseFeed) error {
	var (
		Data                       []hbase.DDL
		feedByte                   []byte
		boardID, feedID, fevFeedID string
	)
	feed, err := json.Marshal(feedData)
	if err != nil {
		logs.Error(L("SaveHbase json Marshal error:"), err)
		return err
	}
	feedByte = []byte(string(feed))
	boardID = strconv.FormatUint(feedData.BoardID, 10)
	feedID = strconv.FormatUint((MinFeedNum + feedData.FeedID), 10)
	fevFeedID = strconv.FormatUint((MaxFeedNum - feedData.FeedID), 10)
	for _, u := range userIDs {
		userID := function.ReverseString(strconv.FormatUint(u, 10))
		Data = append(Data, hbase.DDL{
			RowKey:    []byte(fmt.Sprintf("%s:LastFeed:%s", userID, feedID)),
			Family:    []byte("data"),
			Qualifier: []byte("feed"),
			Value:     feedByte,
		}, hbase.DDL{
			RowKey:    []byte(fmt.Sprintf("%s:%s:NewList", userID, boardID)),
			Family:    []byte("data"),
			Qualifier: []byte(fevFeedID),
			Value:     feedByte,
		}, hbase.DDL{
			RowKey:    []byte(fmt.Sprintf("%s:%s:Home:%s", userID, boardID, feedID)),
			Family:    []byte("data"),
			Qualifier: []byte("feed"),
			Value:     feedByte,
		}, hbase.DDL{
			RowKey:    []byte(fmt.Sprintf("%s:%s:%s:%s", userID, boardID, feedData.FeedType, feedID)),
			Family:    []byte("data"),
			Qualifier: []byte("feed"),
			Value:     feedByte,
		})
	}
	return hbase.Puts(F.HbaseTableName(), Data)
}

//GetLastFeed 查询最新feed
func (F *Feed) GetLastFeed(userID uint64) {
	hbase.GetLastOne(F.HbaseTableName(), fmt.Sprintf("%s:LastFeed", function.ReverseString(strconv.FormatUint(userID, 10))), "data", "feed")
}
