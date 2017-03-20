package model

//Todo 任务表结构
type Todo struct {
	base
	ID       uint64 `orm:"column(id)"`
	SiteID   uint64 `orm:"column(site_id)"`
	BoardID  uint64 `orm:"column(board_id)"`
	BbsID    uint64 `orm:"column(bbs_id)"`
	FeedID   uint64 `orm:"column(feed_id)"`
	FeedType string `orm:"column(feed_type)"`
	UserID   uint64 `orm:"column(user_id)"`
}

//TableName 表名
func (T *Todo) TableName() string {
	return "todo"
}
