package model

import (
	"strconv"

	"github.com/ccyun/GoApp/application/function"
	"github.com/ccyun/GoApp/application/library/hbase"
)

//Feed 任务表结构
type Feed struct {
	base
	ID        uint64 `orm:"column(id)"`
	SiteID    uint64 `orm:"column(site_id)"`
	BoardID   uint64 `orm:"column(board_id)"`
	BbsID     uint64 `orm:"column(bbs_id)"`
	FeedType  string `orm:"column(feed_type)"`
	Data      string `orm:"column(data)"`
	MsgID     string `orm:"column(msg_id)"`
	CreatedAt uint64 `orm:"column(created_at)"`
}

//FeedData feeddata 结构
type FeedData struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	CreatedAt      uint64 `json:"created_at"`
	UserID         uint64 `json:"user_id"`
	Thumb          string `json:"thumb"`
	Type           string `json:"type"`
	Category       string `json:"category"`
	CommentEnabled uint8  `json:"comment_enabled"`
	EndTime        uint64 `json:"end_time"`
	AllowExpired   uint8  `json:"allow_expired"`
	Status         int8   `json:"status"`
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
//'taskReply','taskAudit','taskClose'
func (F *Feed) SaveHbase(userIDs []uint64, feedData Feed, discussID uint64) error {
	client, err := hbase.OpenClient()
	defer hbase.CloseClient(client)
	if err != nil {
		return err
	}
	var data []*hbase.TPut
	valueData := F.makeHbaseValue(feedData)
	for _, u := range userIDs {
		rowkey := function.MakeRowkey(int64(u))
		if discussID > 0 {
			rowkey += "_discuss"
		}
		data = append(data, &hbase.TPut{Row: []byte(rowkey + "_home"), ColumnValues: valueData})
		data = append(data, &hbase.TPut{Row: []byte(rowkey + "_list"), ColumnValues: valueData})
		if feedData.FeedType == "bbs" || feedData.FeedType == "task" || feedData.FeedType == "form" {
			data = append(data, &hbase.TPut{Row: []byte(rowkey + "_" + feedData.FeedType), ColumnValues: valueData})
		}
	}
	// }
	return client.PutMultiple([]byte(F.HbaseTableName()), data)
}

//makeHbaseValue 构造hbase value
func (F *Feed) makeHbaseValue(feedData Feed) []*hbase.TColumnValue {
	var (
		boardID, bbsID, family []byte
		timeStamp              int64
	)
	boardID = []byte(strconv.FormatUint(feedData.BoardID, 10))
	bbsID = []byte(strconv.FormatUint(feedData.BbsID, 10))
	family = []byte("cf")
	timeStamp = int64(feedData.ID)
	return []*hbase.TColumnValue{
		&hbase.TColumnValue{
			Family:    family,
			Qualifier: boardID,
			Value:     bbsID,
			Timestamp: &timeStamp,
		},
	}
}
