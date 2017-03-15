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

//Add 增加待办数据
func (T *Todo) Add(siteID, boardID, bbsID, feedID uint64, feedType string, userIDs []uint64) error {
	if len(userIDs) == 0 {
		return nil
	}
	var data []Todo
	for _, userID := range userIDs {

		data = append(data, Todo{
			SiteID:   siteID,
			BoardID:  boardID,
			BbsID:    bbsID,
			FeedID:   feedID,
			FeedType: feedType,
			UserID:   userID,
		})

	}
	_, err := o.InsertMulti(100000, data)
	return err
}
