package model

import (
	"bbs_server/application/library/httpcurl"
	"fmt"

	"strings"

	"github.com/astaxie/beego/orm"
)

//Msg 消息表
type Msg struct {
	base
	ID         uint64 `orm:"column(id)"`
	SiteID     uint64 `orm:"column(site_id)"`
	BoardID    uint64 `orm:"column(board_id)"`
	DiscussID  uint64 `orm:"column(discuss_id)"`
	BbsID      uint64 `orm:"column(bbs_id)"`
	FeedType   string `orm:"column(feed_type)"`
	FeedID     uint64 `orm:"column(feed_id)"`
	UserID     uint64 `orm:"column(user_id)"`
	UserOrgID  uint64 `orm:"column(user_org_id)"`
	TaskStatus uint8  `orm:"column(task_status)"`
	IsRead     uint8  `orm:"column(is_read)"`
	CreatedAt  uint64 `orm:"column(created_at)"`
}

//TableName 表名
func (M *Msg) TableName() string {
	return "msg"
}

//TrueTableName 真实表名
func (M *Msg) TrueTableName() string {
	return DBPrefix + M.TableName()
}

//Create 创建消息
func (M *Msg) Create(msgData Msg, userIDs []httpcurl.UMSUser, defaultReadStatus uint8, ackReadUserID uint64) error {
	db := orm.NewOrm()
	db.Using("msg")
	sql := "insert into `" + M.TrueTableName() + "`(`site_id`,`board_id`,`discuss_id`,`bbs_id`,`feed_type`,`feed_id`,`user_id`,`user_org_id`,`task_status`,`is_read`,`created_at`) values"
	values := []string{}
	startIndex := 0
	userCount := len(userIDs)
	for true {
		if startIndex >= userCount {
			break
		}
		endIndex := startIndex + 10000
		if endIndex > userCount {
			endIndex = userCount
		}
		for _, u := range userIDs[startIndex:endIndex] {
			isRead := defaultReadStatus
			if isRead == 0 && u.UserID == ackReadUserID {
				isRead = 1
			}
			values = append(values, fmt.Sprintf("(%d,%d,%d,%d,'%s',%d,%d,%d,0,%d,%d)", msgData.SiteID, msgData.BoardID, msgData.DiscussID, msgData.BbsID, msgData.FeedType, msgData.FeedID, u.UserID, u.OrganizationID, isRead, msgData.CreatedAt))
		}
		if _, err := db.Raw(sql + strings.Join(values, ",")).Exec(); err != nil {
			return err
		}
		values = []string{}
		startIndex += 10000
	}
	return nil
}

//GetUserIDs 读取用户id
func (M *Msg) GetUserIDs(siteID, boardID, bbsID uint64, isRead int) []uint64 {
	db := orm.NewOrm()
	db.Using("msg")
	var (
		msgData []*Msg
		data    []uint64
	)
	sql := fmt.Sprintf("select `user_id` from `%s` where `site_id`=%d and `board_id`=%d and `bbs_id`=%d", M.TrueTableName(), siteID, boardID, bbsID)
	if isRead == 0 || isRead == 1 {
		sql += fmt.Sprintf("and `is_read`=%d", isRead)
	} else if isRead != -1 {
		return nil
	}
	db.Raw(sql).QueryRows(&msgData)
	for _, item := range msgData {
		data = append(data, item.UserID)
	}
	return data
}
