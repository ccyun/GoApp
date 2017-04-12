package model

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

//GetUserIDs
func (U *Unread) GetUserIDs() []uint64 {

	return nil
}