package model

import "github.com/astaxie/beego/orm"

//Unread 未读计数
type Unread struct {
	base
	ID            uint64 `orm:"column(id)"`
	SiteID        uint64 `orm:"column(site_id)"`
	BoardID       uint64 `orm:"column(board_id)"`
	FeedType      string `orm:"column(feed_type)"`
	UserID        uint64 `orm:"column(user_id)"`
	UnreadCount   uint64 `orm:"column(unread_count)"`
	AckreadFeedID uint64 `orm:"column(ackread_feed_id)"`
}

//TableName 表名
func (U *Unread) TableName() string {
	return "unread"
}

//GetUserIDs 查询未读计数用户列表
func (U *Unread) GetUserIDs(siteID, boardID uint64, feedType string) ([]uint64, error) {
	var (
		data    []*Unread
		userIDs []uint64
	)
	if _, err := o.QueryTable(U).Filter("SiteID", siteID).Filter("BoardID", boardID).Filter("FeedType", feedType).Limit(-1).All(&data, "UserID"); err != nil {
		if err == orm.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	for _, v := range data {
		userIDs = append(userIDs, v.UserID)
	}
	return userIDs, nil
}
