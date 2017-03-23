package model

//BbsTaskReply 任务表结构
type BbsTaskReply struct {
	base
	ID        uint64 `orm:"column(id)"`
	SiteID    uint64 `orm:"column(site_id)"`
	BoardID   uint64 `orm:"column(board_id)"`
	BbsID     uint64 `orm:"column(bbs_id)"`
	SerialNum uint64 `orm:"column(serial_num)"`
	UserID    uint64 `orm:"column(user_id)"`
	CreatedAt uint64 `orm:"column(CreatedAt)"`
	Status    int8   `orm:"column(status)"`
	CloseAt   uint64 `orm:"column(close_at)"`
}

//TableName 表名
func (B *BbsTaskReply) TableName() string {
	return "bbs_task_reply"
}
