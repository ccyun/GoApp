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
	Data      uint64 `orm:"column(data)"`
	MsgID     uint64 `orm:"column(msg_id)"`
	CreatedAt uint64 `orm:"column(created_at)"`
}

//seed rowkey高位随机种子
var seed = [36]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

//TableName 表名
func (F *Feed) TableName() string {
	return "feed"
}

//HbaseTableName hbase表名
func (F *Feed) HbaseTableName() string {
	return "bbs_feed"
}

//ReverseString 反转字符串
func makeRowkey(userID int64) string {
	userIDstr := strconv.FormatInt(userID, 10)
	reverseUserID, _ := strconv.ParseInt(function.ReverseString(userIDstr), 10, 0)
	seedK1 := userID % 36
	seedK2 := reverseUserID % 36
	seedK3 := (seedK1 + seedK2) % 36
	return seed[seedK3] + seed[seedK1] + seed[seedK2] + "_" + userIDstr
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
		rowkey := makeRowkey(int64(u))
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
