package model

//Todo 任务表结构
type Todo struct {
	base
	ID       uint64 `orm:"column(id)"`
	SiteID   uint64 `orm:"column(site_id)"`
	Type     string `orm:"column(type)"`
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

//GetUNReplyUserIDs 读取未反馈用户列表
func (T *Todo) GetUNReplyUserIDs(BbsID uint64) ([]uint64, error) {
	var (
		TodoList []Todo
		data     []uint64
	)
	_, err := o.QueryTable(T).Filter("BbsID", BbsID).Filter("Type", "unreply").Limit(-1).All(&TodoList, "UserID")
	if err != nil {
		return nil, err
	}
	for _, v := range TodoList {
		data = append(data, v.UserID)
	}
	return data, nil
}
